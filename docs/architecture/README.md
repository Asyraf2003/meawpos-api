# Architecture

This folder defines the structural rules for the Go API.

## Contents

- `0020_hexagonal_go_api.md`: required hexagonal dependency direction and mutation/read flow.
- `0021_package_boundaries.md`: package ownership, file-size policy, and forbidden responsibility mixing.
- `0022_api_capability_control.md`: capability metadata and runtime capability-control requirements.
- `0023_public_contracts.md`: public contract protection rules.
- `0024_current_repo_layout.md`: current repository layout, protected paths, and layout-change rules.

## Use This Folder When

- creating or moving packages;
- adding modules, adapters, handlers, repositories, or use cases;
- changing route exposure or capability-control behavior;
- deciding whether a change needs an ADR.

Architecture rules override convenience. If implementation pressure conflicts with these documents, mark GAP and update the blueprint or ADR first.
