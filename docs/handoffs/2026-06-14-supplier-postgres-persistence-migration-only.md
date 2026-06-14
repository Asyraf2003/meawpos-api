<!--
Copyright (C) 2026 Asyraf Mubarak

This file is part of gopos-api.

gopos-api is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, version 3 only.

gopos-api is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with gopos-api. If not, see <https://www.gnu.org/licenses/>.
-->

# Supplier PostgreSQL Persistence Handoff

## Active Scope

Supplier PostgreSQL persistence slice, migration-only plus repository adapter skeleton checkpoint.

## Status

In progress.

Migration-only step is locally implemented and applied.

Repository adapter skeletons are locally implemented with compile-time port assertion.

Repository behavior implementation has not started; skeleton methods return explicit placeholder errors.

## Files Changed

```text
docs/blueprints/0039_supplier_postgres_persistence_slice.md
migrations/0014_create_suppliers_table.up.sql
migrations/0014_create_suppliers_table.down.sql
internal/platform/postgres/supplier_repository.go
internal/platform/postgres/supplier_repository_row.go
internal/platform/postgres/supplier_repository_write.go
internal/platform/postgres/supplier_repository_query.go
docs/handoffs/2026-06-14-supplier-postgres-persistence-migration-only.md
docs/handoffs/README.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
```

## FACT

- Supplier implementation slice 1 was restored and closed locally with proof after fixing duplicate/split helper issues and aligning ActivateSupplier with FindActiveByNormalizedName.
- Supplier PostgreSQL persistence blueprint `0039_supplier_postgres_persistence_slice.md` is accepted by owner decision.
- The accepted step order says implementation must start with migration-only work.
- Migration slot `0014` was confirmed empty before creating Supplier migration files.
- Migration files created:
  - `migrations/0014_create_suppliers_table.up.sql`
  - `migrations/0014_create_suppliers_table.down.sql`
- Migration creates `suppliers`.
- Supplier table includes:
  - `id`
  - `name`
  - `name_normalized`
  - `phone`
  - `email`
  - `address`
  - `notes`
  - `is_active`
  - `created_at`
  - `updated_at`
- Supplier indexes include:
  - `suppliers_active_name_normalized_unique`
  - `suppliers_active_name_idx`
  - `suppliers_name_normalized_idx`
  - `suppliers_updated_at_idx`
- Repository adapter skeleton files created:
  - `internal/platform/postgres/supplier_repository.go`
  - `internal/platform/postgres/supplier_repository_row.go`
  - `internal/platform/postgres/supplier_repository_write.go`
  - `internal/platform/postgres/supplier_repository_query.go`
- `SupplierRepository` stores the existing PostgreSQL pool dependency pattern.
- `NewSupplierRepository` is the constructor.
- Compile-time assertion verifies the adapter satisfies `ports.SupplierRepository`.
- Query helpers follow the existing transaction-aware `TxFromContext` pattern.
- Row mapping structure targets the `suppliers` table columns from migration `0014`.
- Create, Update, SetActive, FindByID, FindByNormalizedName, FindActiveByNormalizedName, List, and Lookup are present as explicit placeholder behavior.

The active normalized-name uniqueness rule uses a PostgreSQL partial unique index:

```sql
create unique index suppliers_active_name_normalized_unique
on suppliers (name_normalized)
where is_active = true;
```

This preserves the accepted duplicate policy:

- active Supplier name must be unique by normalized name;
- inactive Supplier names do not block active Supplier name reuse;
- reactivating inactive Supplier rejects if another active Supplier already owns the same normalized name.

## Proof Collected

Supplier module proof:

```bash
go test ./internal/modules/supplier/...
```

Visible result:

```text
ok   pos-go/internal/modules/supplier/domain
?    pos-go/internal/modules/supplier/ports [no test files]
ok   pos-go/internal/modules/supplier/usecase
```

Hexagonal proof:

```bash
bash scripts/audit_hexagonal.sh
```

Visible result:

```text
[PASS] hexagonal import audit passed
```

Aggregate proof:

```bash
make verify
```

Visible result:

```text
[PASS] aggregate audit passed
```

Supplier PostgreSQL skeleton compile proof:

```bash
go test ./internal/platform/postgres/... -run Supplier
```

Visible result:

```text
?    pos-go/internal/platform/postgres [no test files]
```

Migration proof:

```bash
bash scripts/db_migrate.sh
```

Visible result:

```text
[APPLY] 0014_create_suppliers_table.up.sql
BEGIN
CREATE TABLE
CREATE INDEX
CREATE INDEX
CREATE INDEX
CREATE INDEX
INSERT 0 1
COMMIT

[PASS] db migrate completed
```

## Tests Run

```bash
go test ./internal/modules/supplier/...
go test ./internal/platform/postgres/... -run Supplier
bash scripts/audit_hexagonal.sh
make verify
bash scripts/db_migrate.sh
```

## Open Gaps

- Supplier PostgreSQL repository behavior is not implemented beyond compile-safe skeletons.
- Supplier repository integration tests are not implemented.
- Supplier query-plan proof is not collected.
- Supplier HTTP runtime is not implemented.
- Supplier capability seed is not implemented.
- Supplier route capability manifest rows are not implemented.
- Faktur is not implemented.
- Inventory/stock movement is not implemented.
- Audit/outbox persistence is not implemented.
- Localization is not implemented.
- Extended filters are not implemented.
- Laravel Supplier MySQL/source parity remains unproven.
- Remote connector validation for final local Supplier migration changes remains pending unless the files become visible through connector.

## Next Valid Active Step

Supplier PostgreSQL repository Create and FindByID behavior.

Keep the next step limited to Create and FindByID behavior plus the focused repository proof needed for that behavior.

## Scope Guard

Do not start Supplier HTTP transport.

Do not add Supplier capability seed.

Do not start Faktur.

Do not start inventory mutation.

Do not start stock movement.

Do not start audit/outbox.

Do not start localization.

Do not start extended filters.

Do not start architecture folder cleanup.

Auth/System ADR 0012 output contract centralization remains deferred by owner decision and must not block Supplier/Faktur progress.

## Estimated Scope Progress Percentage

Supplier PostgreSQL persistence slice: 30%.

Reason:

- blueprint accepted;
- migration files created;
- migration applied locally;
- repository adapter skeleton files created;
- compile-time port assertion exists;
- focused Supplier and PostgreSQL compile proof passed;
- integration tests not started;
- repository behavior still pending;
- query-plan proof not collected.

## Estimated Context-Window Status

Current context is sufficient to start Supplier PostgreSQL repository Create and FindByID behavior in the next session.

Recommended next session target:

```text
Terminal Codex
```

Recommended template source:

```text
docs/templates/0121_codex_session_prompts.md
```
