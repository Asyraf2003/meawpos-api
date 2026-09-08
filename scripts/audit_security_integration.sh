#!/usr/bin/env bash
# Copyright (C) 2026 Asyraf Mubarak
# This file is part of gopos-api and licensed under GNU AGPLv3.

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

if [[ -z "${DATABASE_URL:-}" && -f .env ]]; then
	set -a
	# shellcheck disable=SC1091
	source .env
	set +a
fi
if [[ -z "${DATABASE_URL:-}" ]]; then
	echo "[FAIL] DATABASE_URL is required for mandatory security integration proof"
	exit 1
fi

export GOCACHE="${GOCACHE:-/tmp/go-build-cache}"
proof='RootAuthority|RootCreation|SessionStatusChecker|SessionRevoker|Catalog_|CashSale_|WalkingSkeletonHTTP'

echo "== security audit: PostgreSQL-backed behavior and API abuse =="
echo "tests: $proof"
echo

go test -tags integration -p=1 ./internal/platform/postgres/... ./internal/app/bootstrap \
	-run "$proof" -count=1

echo
echo "[PASS] PostgreSQL-backed security integration audit passed"
