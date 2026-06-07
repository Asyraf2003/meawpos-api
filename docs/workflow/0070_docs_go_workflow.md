# Docs Go Workflow

## Purpose

Keep documentation clean and prevent handoff, blueprint, ADR, archive, and standards from overwriting each other.

## Document Classes

- `docs/`: canonical AI and engineering rules.
- `docs/blueprints/`: active implementation plans.
- `docs/adr/`: accepted decisions that change architecture or contract.
- `docs/handoffs/`: session-to-session continuation notes.
- `docs/evidence/`: command output summaries and proof references.
- `docs/archive/`: old material that is no longer active.

If the future Go repo uses different paths, it must preserve these document roles.

## Blueprint Rules

Blueprints must be scoped to one domain, capability, migration, or architecture decision.

Blueprints must not be used as handoffs.

Blueprints should describe plan, scope, acceptance gates, and step order.

Ongoing session logs, command transcripts, and progress history belong in handoffs or progress ledgers.

A closed blueprint may keep a compact closeout summary or proof reference, but detailed proof history should stay in the linked handoff or ledger.

## Handoff Rules

Handoffs must state:

- last active scope;
- files changed;
- proof collected;
- tests run;
- open gaps;
- next valid active step.

Handoffs must not silently change decisions.

## Archive Rules

Archive is read-only historical context unless the user explicitly revives it.

Archive must not override active rules, ADRs, blueprints, source code, or command output.

## Naming Rules

Use numbered files for ordered standards and blueprints:

```text
0001_index.md
0002_decision_policy.md
0010_domain_name_blueprint.md
```

Do not reuse numbers for different meanings.

## Cleanliness Rules

- One document owns one purpose.
- Every folder under `docs/` should have a `README.md` that explains the folder role, its active files, and when to use it.
- No duplicate active blueprint for the same scope.
- No "latest final v2 fixed" naming.
- No completion status without proof reference.
- No decision hidden only in chat.
