# Decision Policy

## Status

This document is the conflict protocol for all `docs/` rules.

## Rule Hierarchy

1. User's explicit active scope.
2. P0 safety and architecture rules.
3. Public API contract protection.
4. Domain correctness.
5. Database integrity.
6. Test/proof requirements.
7. Local style preference.

Real evidence overrides guesses.

## P0 Rules

- Do not invent facts.
- Do not bypass hexagonal boundaries.
- Do not put business rules in Echo handlers, middleware, SQL builders, or UI contracts.
- Do not expose API behavior outside capability control.
- Do not perform a mutation without transaction, authorization, capability, validation, audit decision, and test proof.
- Do not change public API contract shape without explicit migration notes and contract tests.
- Do not delete data unless the domain lifecycle explicitly allows delete.

## P1 Rules

- Keep files small and single-purpose.
- Prefer explicit ports and DTOs over hidden map-based contracts.
- Use PostgreSQL constraints for database-owned invariants.
- Use focused tests for every use case and adapter.
- Keep docs organized by purpose: standards, blueprint, ADR, handoff, archive.

## Mandatory Decision Sequence

Before choosing an implementation path:

1. Identify proven facts.
2. Identify active step goal.
3. Identify scope in and scope out.
4. Identify affected public contracts.
5. Identify affected domain invariants.
6. Identify affected DB tables and constraints.
7. Identify capability-control impact.
8. Identify test/proof needed.
9. Mark GAP if the data is insufficient.

## GAP Rule

If data is insufficient:

- state what is unknown;
- explain why it blocks the decision;
- ask for the smallest missing proof;
- do not fill the gap with assumptions.

## Decision And Data Request Rule

When missing data blocks work, ask for the smallest specific source batch instead of a broad dump.

For Laravel migration work, name the exact folders, files, routes, migrations, seeders, tests, or command output needed.

When an ADR or owner decision is needed, ask a concise decision question with 2-3 viable options. Each option must include:

- when to choose it;
- tradeoff or risk;
- recommended option first when there is a clear recommendation.

Do not continue implementation across an unresolved ADR-level decision.
