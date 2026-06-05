# AGENTS.md

This repository uses `docsgo/` as the canonical AI and engineering rules package.

Before planning, editing, implementing, reviewing, or proposing commands, read:

1. `docsgo/README.md`
2. `docsgo/AGENTS.md`
3. `docsgo/0001_index.md`
4. `docsgo/0002_decision_policy.md`
5. `docsgo/0003_session_start_protocol.md`

Then read the active blueprint in `docs/blueprints/`.

Mandatory behavior:
- Do not invent facts, repo state, test results, or progress.
- Start from a blueprint before implementation.
- Use one active step per response.
- Show proof before claiming progress.
- Follow local Go package boundaries and existing repo structure.
- Use `fd` for file discovery and `rg` for text search.
- Do not implement domain CRUD without a domain contract and capability keys.
