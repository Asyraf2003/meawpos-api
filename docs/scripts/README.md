# Scripts

This folder defines command and automation contracts.

## Contents

- `0090_makefile_and_scripts.md`: stable Makefile targets, script behavior, seed profiles, and architecture audit expectations.

## Use This Folder When

- adding or changing a Makefile target;
- writing an audit, test, migration, or seed script;
- deciding what command output counts as proof;
- maintaining repeatable checks for terminal, CI, and AI handoffs.

Scripts must be deterministic, print clear PASS/FAIL output, and exit non-zero on failure.
