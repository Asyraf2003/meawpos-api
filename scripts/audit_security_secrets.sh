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
parent_ignore_file="$ROOT_DIR/.gitleaksignore"
docs_ignore_file="$ROOT_DIR/docs/.gitleaksignore"

echo "== security audit: secrets =="
echo "tool: $tool"
echo

if [[ ! -e "$ROOT_DIR/docs/.git" ]]; then
	echo "[FAIL] initialized docs submodule is required for secret scanning"
	exit 1
fi
for repository in "$ROOT_DIR" "$ROOT_DIR/docs"; do
	if [[ "$(git -C "$repository" rev-parse --is-shallow-repository)" != false ]]; then
		echo "[FAIL] complete Git history is required: $repository"
		exit 1
	fi
done
for ignore_file in "$parent_ignore_file" "$docs_ignore_file"; do
	if [[ ! -f "$ignore_file" ]]; then
		echo "[FAIL] repository-local Gitleaks ignore file is required: $ignore_file"
		exit 1
	fi
done

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

go run "$tool" dir --no-banner --no-color --redact --gitleaks-ignore-path "$parent_ignore_file" "$ROOT_DIR"
go run "$tool" git --no-banner --no-color --redact --gitleaks-ignore-path "$parent_ignore_file" "$ROOT_DIR"

go run "$tool" dir --no-banner --no-color --redact --gitleaks-ignore-path "$docs_ignore_file" "$ROOT_DIR/docs"
go run "$tool" git --no-banner --no-color --redact --gitleaks-ignore-path "$docs_ignore_file" "$ROOT_DIR/docs"

echo
echo "[PASS] secret audit passed"
