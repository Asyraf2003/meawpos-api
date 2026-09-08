# AGENTS.md

MiawPOS API uses the `docs/` submodule as its canonical project and engineering rules package.

## Bootstrap

Before planning, editing, implementing, reviewing, or proposing commands, read this file, then follow the canonical read order owned by `docs/AGENTS.md` and any applicable nested `AGENTS.md` files.

The current cross-system chain begins at:

1. `docs/README.md`
2. `docs/AGENTS.md`
3. the canonical files and exact active artifact those instructions name

For Go implementation work, continue with the relevant local rules under `docs/engineering/` only after the cross-system chain has been read.

Prompts do not need to restate the full canonical read list when the agent can access this repository. The agent is responsible for following `AGENTS.md`, discovering the active scope, and inspecting directly relevant files itself.

## Mandatory Behavior

- Do not invent facts, repository state, test results, or progress.
- Start implementation from an active blueprint.
- Keep one active scope and one complete proof chain.
- One active scope does not mean one tiny action: batch independent repository reads, searches, checks, edits, and proofs that belong to the same authorized goal.
- Sequence work only when a later action genuinely depends on an earlier result.
- Execute the largest safe bounded slice that completes the active goal without widening scope.
- Do not stop merely to report an intermediate fact that requires no decision. Report at decision, blocker, proof-failure, or scope-completion boundaries.
- When repository tools can inspect a fact directly, use them instead of asking the owner/Codex/local terminal to run equivalent discovery commands. Request local proof only for inaccessible or local-only state.
- Use `fd` for local file discovery and `rg` for local text search when local CLI is the execution surface.
- Treat current Go structure and historical implementation detail according to the accepted R1 classification and canonical foundation direction, not as automatic future architecture.
- Do not implement domain CRUD without an accepted domain contract and authorization/capability decisions appropriate to that domain.
- Do not make trusted financial, authorization, ROOT-isolation, transaction, audit, or security invariants removable capabilities.
- Do not claim SQLite parity until the relevant conformance proof exists.
- Treat `docs/` changes as changes in the documentation submodule; prove parent and submodule state separately.

## Current Execution Rule

The canonical throughput and stop rules live in `docs/transition/0000_master_execution_workflow.md`.

In short:

```text
batch facts -> classify gaps -> decide where required -> execute bounded goal -> aggregate proof -> report
```

not:

```text
inspect one fact -> report -> inspect next fact -> report
```
