#!/usr/bin/env bash
# Copyright (C) 2026 Asyraf Mubarak
# This file is part of gopos-api and licensed under GNU AGPLv3.

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

# shellcheck disable=SC1091
source scripts/config/security_tools.env
export GOTOOLCHAIN=local
tool="${GITLEAKS_MODULE}@${GITLEAKS_VERSION}"
ignore_file="$ROOT_DIR/.gitleaksignore"

echo "== security audit: secrets =="
echo "tool: $tool"
echo

sentinel_dir="$(mktemp -d /tmp/pos-go-gitleaks-sentinel.XXXXXX)"
trap 'rm -rf "$sentinel_dir"' EXIT
printf 'credential = "%s%s"\n' 'ghp_' 'Q7mZ2pL9xR4vN8cT5kW3sH6jF1bD0yUaE2qG' > "$sentinel_dir/sentinel.txt"
set +e
sentinel_output="$(go run "$tool" dir --no-banner --no-color --redact "$sentinel_dir" 2>&1)"
sentinel_status=$?
set -e
if [[ $sentinel_status -eq 0 || "$sentinel_output" != *"leaks found"* ]]; then
	echo "[FAIL] Gitleaks sentinel was not detected"
	printf '%s\n' "$sentinel_output"
	exit 1
fi
echo "[PASS] Gitleaks detection sentinel"

go run "$tool" dir --no-banner --no-color --redact --gitleaks-ignore-path "$ignore_file" "$ROOT_DIR"
go run "$tool" git --no-banner --no-color --redact --gitleaks-ignore-path "$ignore_file" "$ROOT_DIR"

if git -C docs rev-parse --is-inside-work-tree >/dev/null 2>&1; then
	go run "$tool" dir --no-banner --no-color --redact "$ROOT_DIR/docs"
	go run "$tool" git --no-banner --no-color --redact "$ROOT_DIR/docs"
fi

echo
echo "[PASS] secret audit passed"
