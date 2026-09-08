# AGENTS.md

MiawPOS API uses the `docs/` submodule as its canonical project and engineering rules package.

Before planning, editing, implementing, reviewing, or proposing commands, read in this order:

1. `docs/README.md`
2. `docs/AGENTS.md`
3. `docs/0000_index.md`
4. `docs/0001_north_star.md`
5. `docs/0002_decision_policy.md`
6. `docs/0003_document_authority_register.md`
7. `docs/transition/0000_master_execution_workflow.md`
8. `docs/transition/0002_foundation_redesign_roadmap.md`
9. the exact active blueprint named by the roadmap or owner

For Go implementation work, then read the relevant local rules under `docs/engineering/`.

Mandatory behavior:

- Do not invent facts, repository state, test results, or progress.
- Start implementation from an active blueprint.
- Keep one active scope and one complete proof chain.
- Use `fd` for file discovery and `rg` for text search.
- Treat current Go structure and historical implementation detail as evidence pending the foundation runtime audit, not as automatic future architecture.
- Do not implement domain CRUD without an accepted domain contract and authorization/capability decisions appropriate to that domain.
- Do not make trusted financial, authorization, ROOT-isolation, transaction, audit, or security invariants removable capabilities.
- Do not claim SQLite parity until the relevant conformance proof exists.
- Treat `docs/` changes as changes in the documentation submodule; prove parent and submodule status separately.
