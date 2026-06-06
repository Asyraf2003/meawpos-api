#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

if [[ -z "${DATABASE_URL:-}" && -f ".env" ]]; then
  set -a
  # shellcheck disable=SC1091
  source .env
  set +a
fi

if [[ -z "${DATABASE_URL:-}" ]]; then
  echo "[FAIL] DATABASE_URL is required"
  echo "Hint: copy .env.example to .env and adjust DATABASE_URL."
  exit 1
fi

if ! command -v psql >/dev/null 2>&1; then
  echo "[FAIL] psql is required"
  exit 1
fi

python3 - "$DATABASE_URL" <<'PY' > /tmp/pos_go_db_env.sh
import sys
from urllib.parse import urlparse, unquote

url = sys.argv[1]
parsed = urlparse(url)

if parsed.scheme not in ("postgres", "postgresql"):
    raise SystemExit("[FAIL] DATABASE_URL must use postgres:// or postgresql://")

user = unquote(parsed.username or "")
password = unquote(parsed.password or "")
host = parsed.hostname or "127.0.0.1"
port = parsed.port or 5432
dbname = (parsed.path or "").lstrip("/")

if not user or not password or not dbname:
    raise SystemExit("[FAIL] DATABASE_URL must include user, password, and database name")

def sh_quote(value: str) -> str:
    return "'" + value.replace("'", "'\"'\"'") + "'"

print(f"DB_USER={sh_quote(user)}")
print(f"DB_PASSWORD={sh_quote(password)}")
print(f"DB_HOST={sh_quote(host)}")
print(f"DB_PORT={sh_quote(str(port))}")
print(f"DB_NAME={sh_quote(dbname)}")
PY

# shellcheck disable=SC1091
source /tmp/pos_go_db_env.sh
rm -f /tmp/pos_go_db_env.sh

POSTGRES_ADMIN_USER="${POSTGRES_ADMIN_USER:-postgres}"
POSTGRES_ADMIN_DB="${POSTGRES_ADMIN_DB:-postgres}"

echo "== db dev setup =="
echo "admin: ${POSTGRES_ADMIN_USER}@${DB_HOST}:${DB_PORT}/${POSTGRES_ADMIN_DB}"
echo "app database: ${DB_NAME}"
echo "app user: ${DB_USER}"
echo

psql -h "$DB_HOST" -p "$DB_PORT" -U "$POSTGRES_ADMIN_USER" -d "$POSTGRES_ADMIN_DB" -v ON_ERROR_STOP=1 \
  -v app_user="$DB_USER" \
  -v app_password="$DB_PASSWORD" \
  -v app_db="$DB_NAME" <<'SQL'
SELECT format('CREATE ROLE %I LOGIN PASSWORD %L', :'app_user', :'app_password')
WHERE NOT EXISTS (
  SELECT 1 FROM pg_roles WHERE rolname = :'app_user'
)\gexec

SELECT format('ALTER ROLE %I WITH LOGIN PASSWORD %L', :'app_user', :'app_password')\gexec

SELECT format('CREATE DATABASE %I OWNER %I', :'app_db', :'app_user')
WHERE NOT EXISTS (
  SELECT 1 FROM pg_database WHERE datname = :'app_db'
)\gexec

GRANT ALL PRIVILEGES ON DATABASE :"app_db" TO :"app_user";
SQL

echo
echo "== app connection check =="
psql "$DATABASE_URL" -v ON_ERROR_STOP=1 -c "select current_user, current_database();"

echo
echo "[PASS] db dev setup completed"
