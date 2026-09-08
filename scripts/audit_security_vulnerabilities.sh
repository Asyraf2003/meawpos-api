#!/usr/bin/env bash
# Copyright (C) 2026 Asyraf Mubarak
# This file is part of gopos-api and licensed under GNU AGPLv3.

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

# shellcheck disable=SC1091
source scripts/config/security_tools.env
export GOTOOLCHAIN=local

echo "== security audit: known vulnerabilities =="
echo "tool: ${GOVULNCHECK_MODULE}@${GOVULNCHECK_VERSION}"
echo

go run "${GOVULNCHECK_MODULE}@${GOVULNCHECK_VERSION}" ./...

echo
echo "[PASS] known-vulnerability audit passed"
