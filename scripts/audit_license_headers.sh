#!/usr/bin/env bash
# Copyright (C) 2026 Asyraf Mubarak

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$REPO_ROOT"

echo "== repository license presentation audit =="

[[ -f LICENSE ]] || { echo "[FAIL] missing root LICENSE"; exit 1; }
rg -Fq "GNU AFFERO GENERAL PUBLIC LICENSE" LICENSE || { echo "[FAIL] LICENSE missing AGPL title"; exit 1; }
rg -Fq "Version 3, 19 November 2007" LICENSE || { echo "[FAIL] LICENSE missing AGPL version"; exit 1; }
rg -Fq "END OF TERMS AND CONDITIONS" LICENSE || { echo "[FAIL] LICENSE is incomplete"; exit 1; }

if [[ -f NOTICE ]]; then
  rg -Fq "MiawPOS API" NOTICE || { echo "[FAIL] NOTICE has stale project identity"; exit 1; }
  rg -Fq "GNU Affero General Public License v3.0 only" NOTICE || { echo "[FAIL] NOTICE missing license summary"; exit 1; }
fi

echo "[PASS] root license presentation is valid"
echo "[INFO] repeated first-party per-file license headers are not required"
