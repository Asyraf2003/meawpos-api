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

# Supplier PostgreSQL Persistence Slice Blueprint

## Status

Draft for owner acceptance.

## Date

2026-06-14

## Active Scope

Plan Supplier PostgreSQL persistence after Supplier implementation slice 1 domain, ports, and usecase contracts were locally closed with proof.

## Source Slices

```text
docs/blueprints/0037_supplier_domain_contract.md
docs/blueprints/0038_supplier_implementation_slice_1.md
```

## Closeout Handoff

```text
docs/handoffs/2026-06-14-supplier-implementation-slice-1-closeout.md
```

## FACT

- Supplier implementation slice 1 is locally closed with proof.
- Closed Go package scope:
  - `internal/modules/supplier/domain`
  - `internal/modules/supplier/ports`
  - `internal/modules/supplier/usecase`
- Supplier domain contract accepts Supplier as independent master data for future Faktur flows.
- Supplier must not own ProductCatalog data.
- Supplier must not mutate stock.
- Supplier must not create Faktur.
- Supplier lifecycle uses active/inactive status.
- Physical delete is forbidden for normal use.
- Accepted duplicate policy:
  - active Supplier name must be unique by normalized name;
  - inactive Supplier names do not block active Supplier name reuse;
  - reactivating inactive Supplier rejects if another active Supplier already owns the same normalized name.
- Supplier PostgreSQL table proposal: `suppliers`.

Accepted fields:

```text
id
name
name_normalized
phone nullable
email nullable
address nullable
notes nullable
is_active
created_at
updated_at
```

## GAP

- No Supplier PostgreSQL migration exists yet.
- No Supplier PostgreSQL repository adapter exists yet.
- No Supplier PostgreSQL integration tests exist yet.
- Laravel Supplier source parity is not proven.
- Exact Laravel Supplier normalization behavior is not confirmed.
- No Supplier HTTP routes exist yet.
- No Supplier capability seed exists yet.
- No Supplier runtime/capability proof exists yet.
- No Faktur implementation exists yet.
- No inventory or stock movement implementation exists yet.
- No audit/outbox implementation exists yet.

## DECISION

This slice should implement Supplier PostgreSQL persistence only.

Do not implement Echo HTTP transport.

Do not register routes.

Do not add capability seed migrations.

Do not start Faktur.

Do not start inventory mutation.

Do not start stock movement.

Do not start audit/outbox.

Do not start localization.

Do not start extended filters.

Do not start architecture folder cleanup.

Auth/System ADR 0012 output contract centralization remains deferred by owner decision and must not block Supplier/Faktur progress.

## SCOPE-IN

Expected new files:

```text
migrations/0014_create_suppliers_table.up.sql
migrations/0014_create_suppliers_table.down.sql
internal/platform/postgres/supplier_repository.go
internal/platform/postgres/supplier_repository_row.go
internal/platform/postgres/supplier_repository_write.go
internal/platform/postgres/supplier_repository_query.go
internal/platform/postgres/supplier_repository_integration_helpers_test.go
internal/platform/postgres/supplier_repository_create_integration_test.go
internal/platform/postgres/supplier_repository_update_integration_test.go
internal/platform/postgres/supplier_repository_lifecycle_integration_test.go
internal/platform/postgres/supplier_repository_query_integration_test.go
```

If file-size audit requires smaller files, split files without changing package boundaries.

## SCOPE-OUT

```text
internal/modules/supplier/transport/http
internal/presentation/http/id/supplier
cmd/server route wiring
capability seed migration
Faktur
inventory stock mutation
stock movement
audit/outbox
localization
extended filters
UI
```

## POSTGRESQL SCHEMA

Table:

```text
suppliers
```

Columns:

```sql
id text primary key,
name text not null,
name_normalized text not null,
phone text null,
email text null,
address text null,
notes text null,
is_active boolean not null default true,
created_at timestamptz not null default now(),
updated_at timestamptz not null default now()
```

## POSTGRESQL INDEXES

Active name uniqueness:

```sql
create unique index suppliers_active_name_normalized_unique
on suppliers (name_normalized)
where is_active = true;
```

Read/search indexes:

```sql
create index suppliers_active_name_idx
on suppliers (is_active, name_normalized, id);

create index suppliers_name_normalized_idx
on suppliers (name_normalized);

create index suppliers_updated_at_idx
on suppliers (updated_at);
```

Index rules:

- Use partial unique index for active duplicate policy.
- Do not make inactive Supplier names block active name reuse.
- Do not add speculative indexes for future Faktur joins in this slice.
- Do not add ProductCatalog or inventory joins in this slice.

## REPOSITORY BEHAVIOR

SupplierRepository adapter must implement:

```text
Create
Update
FindByID
FindByNormalizedName
FindActiveByNormalizedName
List
Lookup
SetActive
```

Rules:

- Repository accepts and returns Supplier domain objects.
- Repository must not import Echo or HTTP.
- Repository must not own capability behavior.
- Repository must not mutate ProductCatalog.
- Repository must not mutate stock.
- Repository must not create Faktur.
- Repository must preserve domain-generated normalized values.
- Repository must translate missing rows to bool=false where the port expects bool found.
- Repository must support deterministic list and lookup ordering.

## QUERY BEHAVIOR

List should support existing Supplier list filter:

```text
Query
Status
Page
PerPage
```

Status behavior:

```text
active -> is_active = true
inactive -> is_active = false
all -> no is_active filter
```

Defaults remain owned by usecase:

```text
page = 1
per_page = 10
status = active
```

Lookup should support existing Supplier lookup filter:

```text
Query
Limit
ActiveOnly
```

Lookup rules:

- Default active-only unless ActiveOnly is false.
- Bounded max limit must remain usecase/adapter safe.
- Deterministic order by name_normalized and id.

## PERFORMANCE AND FLEXIBILITY STANDARD

Performance rules:

- Create -> insert path must rely on primary key and active-name partial unique index.
- Update -> primary key lookup/write path must use suppliers primary key.
- FindByID/show -> primary key lookup must use suppliers primary key.
- FindActiveByNormalizedName -> must use suppliers_active_name_normalized_unique-compatible path.
- List active/default -> must be bounded by Page and PerPage.
- Lookup -> must be bounded by Limit and must not scan unbounded rows.

Query-plan proof rule:

Integration proof should include EXPLAIN or EXPLAIN ANALYZE notes for show/find-by-id, active list first page, lookup bounded search, and active-name duplicate guard when local database is available.

No fake SLA rule:

Do not claim sub-second or millisecond performance without local database proof.

Flexibility rules:

- Keep Supplier list and lookup adapter translation centralized.
- Do not spread SQL filter construction across usecase or HTTP layers.
- Keep future sort/filter additions localized to supplier_repository_query.go.
- Do not add Faktur joins in this slice.
- Do not over-index speculative filters until a query exists and proof shows the need.

## TRANSACTION POLICY

This slice may use the existing platform/postgres transaction context pattern if needed.

Supplier create/update/activate/deactivate must be compatible with later audit/outbox orchestration, but audit/outbox writes are not implemented in this slice.

If transaction reuse is blocked by current platform shape, document the gap in the handoff and keep behavior isolated.

## TEST MATRIX

Migration proof:

- Up migration creates suppliers.
- Down migration drops suppliers.
- suppliers.id is primary key.
- suppliers.name rejects null.
- suppliers.name_normalized rejects null.
- suppliers.is_active defaults true.
- active name_normalized is unique among active suppliers.
- inactive supplier name does not block active name reuse.

Repository integration tests:

- Create stores supplier fields.
- Create stores normalized name from domain.
- Create rejects duplicate active normalized name.
- Create allows inactive name reuse.
- Update changes supplier fields.
- Update rejects duplicate active normalized name when current supplier is active.
- FindByID returns supplier.
- FindByID returns found=false for missing supplier.
- FindByNormalizedName returns matching supplier.
- FindActiveByNormalizedName returns only active supplier.
- SetActive deactivates supplier.
- SetActive activates supplier.
- SetActive rejects duplicate activation through repository/usecase proof path.
- List filters active, inactive, and all.
- List supports pagination.
- Lookup excludes inactive by default.
- Lookup can include inactive when ActiveOnly is false.
- Lookup respects limit.

## PROOF REQUIRED

Focused Supplier module proof:

```bash
go test ./internal/modules/supplier/...
```

Focused PostgreSQL proof:

```bash
go test ./internal/platform/postgres/... -run Supplier
```

Hexagonal proof:

```bash
bash scripts/audit_hexagonal.sh
```

Aggregate proof:

```bash
make verify
```

Migration and query-plan proof if local DB is available:

```bash
make db-up
make db-status
```

Expected migration status must include:

```text
0014_create_suppliers_table.up.sql applied
```

Additional local DB proof should capture EXPLAIN or EXPLAIN ANALYZE for:

- suppliers primary-key show/find-by-id
- suppliers active list first page
- suppliers lookup bounded search
- suppliers active name duplicate guard

## API CONTRACT IMPACT

No HTTP API contract change in this slice.

Supplier runtime API remains out of scope.

## CAPABILITY IMPACT

No capability key or seed change in this slice.

Capability seed belongs to a later Supplier runtime/capability slice.

## ARCHITECTURE IMPACT

This slice adds Supplier persistence under:

```text
internal/platform/postgres
migrations
```

It must not add imports from domain/usecase into transport or HTTP layers.

Hexagonal boundaries must remain enforced by make verify.

## RISKS

- The active-name uniqueness rule must be implemented as a PostgreSQL partial unique index, not a global unique index.
- A global unique index on name_normalized would incorrectly block inactive name reuse.
- Search with broad LIKE/ILIKE may need a better strategy later, but only after proof.
- Faktur joins are out of scope and must not leak into Supplier persistence.
- Laravel Supplier parity is still unproven and must be documented until source evidence is available.

## STEP ORDER

1. Accept or revise this blueprint.
2. Add migration only.
3. Add repository adapter skeletons.
4. Add repository integration helpers.
5. Implement Create and FindByID behavior.
6. Implement Update behavior.
7. Implement FindByNormalizedName and FindActiveByNormalizedName behavior.
8. Implement SetActive behavior.
9. Implement List and Lookup behavior.
10. Run focused Supplier module proof.
11. Run focused PostgreSQL proof.
12. Run aggregate make verify.
13. Update ledger and handoff.
14. Close this slice only after local proof and connector validation.

## ACCEPTANCE RULE

This blueprint is not accepted until the owner explicitly accepts it.

Implementation must not start before acceptance.
