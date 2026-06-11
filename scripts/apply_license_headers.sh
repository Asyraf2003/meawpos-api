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

ROOT = Path(".")
COPYRIGHT = "Copyright (C) 2026 Asyraf Mubarak"

HEADER_LINES = [
    "Copyright (C) 2026 Asyraf Mubarak",
    "",
    "This file is part of gopos-api.",
    "",
    "gopos-api is free software: you can redistribute it and/or modify",
    "it under the terms of the GNU Affero General Public License as published by",
    "the Free Software Foundation, version 3 only.",
    "",
    "gopos-api is distributed in the hope that it will be useful,",
    "but WITHOUT ANY WARRANTY; without even the implied warranty of",
    "MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the",
    "GNU Affero General Public License for more details.",
    "",
    "You should have received a copy of the GNU Affero General Public License",
    "along with gopos-api. If not, see <https://www.gnu.org/licenses/>.",
]

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

def comment_prefix(path: Path) -> str | None:
    if path.suffix == ".go":
        return "//"
    if path.suffix == ".sql":
        return "--"
    if path.suffix in {".sh", ".py", ".yml", ".yaml", ".toml", ".mk"}:
        return "#"
    if path.name == "Makefile":
        return "#"
    if path.name.startswith("Dockerfile"):
        return "#"
    return None

def build_header(prefix: str) -> str:
    rendered = []
    for line in HEADER_LINES:
        if line:
            rendered.append(f"{prefix} {line}")
        else:
            rendered.append(prefix)
    return "\n".join(rendered) + "\n\n"

def insert_header(text: str, header: str) -> str:
    if text.startswith("#!"):
        newline_index = text.find("\n")
        if newline_index == -1:
            return text + "\n" + header
        return text[: newline_index + 1] + header + text[newline_index + 1 :]
    return header + text

changed = []
skipped_binary = []

for path in sorted(ROOT.rglob("*")):
    if not path.is_file():
        continue
    if is_excluded(path):
        continue

    prefix = comment_prefix(path)
    if prefix is None:
        continue

    try:
        text = path.read_text()
    except UnicodeDecodeError:
        skipped_binary.append(str(path))
        continue

    first_lines = "\n".join(text.splitlines()[:30])
    if COPYRIGHT in first_lines:
        continue

    path.write_text(insert_header(text, build_header(prefix)))
    changed.append(str(path))

print("== apply license headers ==")
if changed:
    for item in changed:
        print(f"[UPDATED] {item}")
else:
    print("[OK] no files needed header updates")

if skipped_binary:
    print()
    print("== skipped non-text files ==")
    for item in skipped_binary:
        print(f"[SKIP] {item}")
PY
