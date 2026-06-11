#!/usr/bin/env bash
# Copyright (C) 2026 Asyraf Mubarak
#
# This file is part of gopos-api.
#
# gopos-api is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, version 3 only.
#
# gopos-api is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

if [[ -z "${DATABASE_URL:-}" && -f ".env" ]]; then
  set -a
  source .env
  set +a
fi

if [[ -z "${DATABASE_URL:-}" ]]; then
  echo "[FAIL] DATABASE_URL is required"
  exit 1
fi

echo "== db migrate =="
echo "database: DATABASE_URL is set"
echo

psql "$DATABASE_URL" -v ON_ERROR_STOP=1 <<'SQL'
CREATE TABLE IF NOT EXISTS schema_migrations (
    name text PRIMARY KEY,
    applied_at timestamptz NOT NULL DEFAULT now()
);
SQL

for file in $(find migrations -maxdepth 1 -type f -name '*.up.sql' | sort); do
  name="$(basename "$file")"

  applied="$(psql "$DATABASE_URL" -t -A -v ON_ERROR_STOP=1 -c "SELECT 1 FROM schema_migrations WHERE name = '$name' LIMIT 1;")"

  if [[ "$applied" == "1" ]]; then
    echo "[SKIP] already applied: $name"
    continue
  fi

  echo "[APPLY] $name"

  psql "$DATABASE_URL" -v ON_ERROR_STOP=1 <<SQL
BEGIN;
\i $file
INSERT INTO schema_migrations (name) VALUES ('$name');
COMMIT;
SQL
done

echo
echo "[PASS] db migrate completed"
