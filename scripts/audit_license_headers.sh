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

python3 - <<'PY'
from pathlib import Path
import sys

ROOT = Path(".")
COPYRIGHT = "Copyright (C) 2026 Asyraf Mubarak"

EXCLUDED_DIRS = {
    ".git",
    ".bin",
    ".cache",
    "vendor",
    "node_modules",
    "tmp",
    "dist",
    "build",
    "coverage",
}

EXCLUDED_NAMES = {
    "LICENSE",
    "NOTICE",
    "go.mod",
    "go.sum",
}

def is_excluded(path: Path) -> bool:
    parts = set(path.parts)
    if parts & EXCLUDED_DIRS:
        return True
    if path.name in EXCLUDED_NAMES:
        return True
    if path.name.endswith(".min.js"):
        return True
    return False

def is_target(path: Path) -> bool:
    if path.suffix in {".go", ".sh", ".py", ".sql", ".yml", ".yaml", ".toml", ".mk"}:
        return True
    if path.name == "Makefile":
        return True
    if path.name.startswith("Dockerfile"):
        return True
    return False

missing = []
unreadable = []

for path in sorted(ROOT.rglob("*")):
    if not path.is_file():
        continue
    if is_excluded(path):
        continue
    if not is_target(path):
        continue

    try:
        first_lines = "\n".join(path.read_text().splitlines()[:30])
    except UnicodeDecodeError:
        unreadable.append(str(path))
        continue

    if COPYRIGHT not in first_lines:
        missing.append(str(path))

print("== license header audit ==")

fail = False

if not Path("LICENSE").is_file():
    print("[FAIL] missing LICENSE")
    fail = True
else:
    license_text = Path("LICENSE").read_text(errors="replace")
    required_license_markers = [
        "GNU AFFERO GENERAL PUBLIC LICENSE",
        "Version 3, 19 November 2007",
        "END OF TERMS AND CONDITIONS",
    ]
    for marker in required_license_markers:
        if marker not in license_text:
            print(f"[FAIL] LICENSE missing marker: {marker}")
            fail = True

if not Path("NOTICE").is_file():
    print("[FAIL] missing NOTICE")
    fail = True
else:
    notice_text = Path("NOTICE").read_text(errors="replace")
    required_notice_markers = [
        "gopos-api",
        "Copyright (C) 2026 Asyraf Mubarak",
        "GNU Affero General Public License v3.0 only",
        "AGPL-3.0",
    ]
    for marker in required_notice_markers:
        if marker not in notice_text:
            print(f"[FAIL] NOTICE missing marker: {marker}")
            fail = True

for path in missing:
    print(f"[FAIL] missing license header: {path}")
    fail = True

for path in unreadable:
    print(f"[WARN] skipped non-text target: {path}")

if fail:
    print()
    print("[FAIL] license header audit failed")
    sys.exit(1)

print("[PASS] license header audit passed")
PY
