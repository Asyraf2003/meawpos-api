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

GOSEC_BIN="${GOSEC_BIN:-/home/asyraf/go/bin/gosec}"

if [[ ! -x "$GOSEC_BIN" ]]; then
  echo "[FAIL] gosec binary not found or not executable: $GOSEC_BIN"
  exit 1
fi

echo "== security audit: gosec =="
echo "binary: $GOSEC_BIN"
echo

"$GOSEC_BIN" ./...

echo
echo "[PASS] gosec audit passed"
