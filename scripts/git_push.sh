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
COUNTER_FILE="$ROOT_DIR/.git/.auto_commit_counter"

# Fungsi internal untuk eksekusi push per direktori
push_repo() {
    local target_dir="$1"
    local repo_name="$2"
    
    cd "$target_dir"
    
    local branch
    branch="$(git rev-parse --abbrev-ref HEAD 2>/dev/null || true)"
    if [[ -z "$branch" || "$branch" == "HEAD" ]]; then
        echo "[FAIL] [$repo_name] unable to detect active git branch"
        exit 1
    fi

    # Ambil atau inisialisasi counter
    local last_number
    if [[ -f "$COUNTER_FILE" ]]; then
        last_number="$(cat "$COUNTER_FILE")"
    else
        last_number="0"
    fi

    if ! [[ "$last_number" =~ ^[0-9]+$ ]]; then
        last_number="0"
    fi

    local next_number=$((last_number + 1))
    local msg="commit $next_number"

    # Cek staging area
    git add .

    if git diff --cached --quiet; then
        echo "[INFO] [$repo_name] no staged changes to commit. Skipping."
        return 0
    fi

    echo "== git push [$repo_name] =="
    echo "branch: $branch"
    echo "message: $msg"
    echo
    git status --short
    echo

    git commit -m "$msg"
    git push origin "$branch"
    
    echo "$next_number" > "$COUNTER_FILE"
    echo "[PASS] [$repo_name] push completed"
    echo
}

# 1. Jalankan untuk submodule docs terlebih dahulu jika foldernya ada
if [[ -d "$ROOT_DIR/docs/.git" || -f "$ROOT_DIR/docs/.git" ]]; then
    push_repo "$ROOT_DIR/docs" "submodule: docs"
fi

# 2. Jalankan untuk root repo pos-go
push_repo "$ROOT_DIR" "root: pos-go"

echo "[PASS] All operations completed successfully."
