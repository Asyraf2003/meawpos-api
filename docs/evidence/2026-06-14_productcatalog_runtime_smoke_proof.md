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

# ProductCatalog Runtime Smoke Proof

Date recorded: 2026-06-14

Blueprint:

```text
docs/blueprints/0033_productcatalog_runtime_smoke_proof_slice.md
```

## FACT

The ProductCatalog catalog API has now been proven through the local runtime path:

```text
local PostgreSQL -> migrations -> Echo server -> manual auth token -> protected ProductCatalog HTTP routes -> DB-backed responses
```

The proof is local runtime proof only. It does not close Product inventory/stock API, ProductCatalog audit/outbox persistence, shared success envelope centralization, ADR `0012`, runtime localization, or extended Laravel filters.

Bearer tokens, refresh tokens, raw auth payloads, database credentials, and secrets are intentionally omitted from this evidence.

## PROOF SOURCE

Runtime proof was collected during the interrupted 2026-06-13 to 2026-06-14 Terminal Codex smoke session and resumed from the owner-provided proof summary.

The interrupted run produced sanitized terminal output showing the checks below. The owner then manually stopped the still-running server process on port `18081` and verified the port was closed.

## COMMANDS AND RESULTS

Migration proof:

```text
make db-migrate

[SKIP] already applied: 0001_auth_minimum.up.sql
[SKIP] already applied: 0002_authorization_minimum.up.sql
[SKIP] already applied: 0003_authorization_seed_minimum.up.sql
[SKIP] already applied: 0004_authorization_role_permissions_seed.up.sql
[SKIP] already applied: 0005_authorization_assign_base_role_to_existing_accounts.up.sql
[SKIP] already applied: 0006_capability_control.up.sql
[SKIP] already applied: 0007_seed_existing_protected_capabilities.up.sql
[SKIP] already applied: 0008_seed_capability_manage_permission.up.sql
[SKIP] already applied: 0009_create_service_catalog_items.up.sql
[SKIP] already applied: 0010_seed_service_catalog_permissions_capabilities.up.sql
[SKIP] already applied: 0011_create_product_catalog_tables.up.sql
[SKIP] already applied: 0012_add_product_version_timeline_order_index.up.sql
[SKIP] already applied: 0013_seed_product_catalog_permissions_capabilities.up.sql
[PASS] db migrate completed
```

Migration status proof:

```text
make db-status

[APPLIED] 0001_auth_minimum.up.sql
[APPLIED] 0002_authorization_minimum.up.sql
[APPLIED] 0003_authorization_seed_minimum.up.sql
[APPLIED] 0004_authorization_role_permissions_seed.up.sql
[APPLIED] 0005_authorization_assign_base_role_to_existing_accounts.up.sql
[APPLIED] 0006_capability_control.up.sql
[APPLIED] 0007_seed_existing_protected_capabilities.up.sql
[APPLIED] 0008_seed_capability_manage_permission.up.sql
[APPLIED] 0009_create_service_catalog_items.up.sql
[APPLIED] 0010_seed_service_catalog_permissions_capabilities.up.sql
[APPLIED] 0011_create_product_catalog_tables.up.sql
[APPLIED] 0012_add_product_version_timeline_order_index.up.sql
[APPLIED] 0013_seed_product_catalog_permissions_capabilities.up.sql
```

Server start proof:

```text
GOCACHE=/tmp/gopos-go-build AUTH_DEBUG_ENABLED=true HTTP_PORT=18081 make run
```

The first sandboxed `curl` could not reach the unsandboxed server:

```text
curl: (7) Failed to connect to 127.0.0.1 port 18081
health_status=000
```

The same health check outside the sandbox reached the running Echo server:

```text
health_status=200
health_body_keys=database,status
health_status_field=ok
```

Sanitized ProductCatalog runtime smoke output:

```text
manual_login status=200 access_token_present=True refresh_token_present=True
unauth_products status=401 success=False error_code=authentication_required has_meta=True
capability_product_catalog_list_initial status=200 enabled=True
list_products status=200 success=True items=0 has_meta=True
lookup_products status=200 success=True items=0 has_meta=True
create_product status=201 success=True id_present=True kode_barang_present=True has_meta=True
show_product status=200 success=True id_match=True kode_barang_present=True status_field=active has_meta=True
db_product_row table=products count=1
disable_capability key=product_catalog.list status=200 enabled=False
list_products_capability_disabled status=403 success=False error_code=capability_disabled has_meta=True
enable_capability key=product_catalog.list status=200 enabled=True
list_products_after_reenable status=200 success=True items=1
[PASS] product catalog runtime smoke proof passed
```

Shutdown proof:

```text
server process on port 18081 was manually stopped after the interrupted run
port 18081 verified closed after stop
```

## DECISION

ProductCatalog runtime smoke proof is locally proven.

ProductCatalog catalog API remains closed locally for pure API scope.

Product/inventory area remains partial because stock API, audit/outbox, extended filters, localization policy, and shared output-envelope work are not closed.

## GAP

Current remaining Product/API-only gaps:

- Product inventory/stock API is not closed.
- ProductCatalog audit/outbox persistence is not implemented.
- Shared success envelope centralization is not implemented.
- ADR `0012` output contract centralization remains partial across all API surfaces.
- Runtime language/output policy is not implemented beyond the current Indonesian public DTO mapping.
- Extended Laravel ProductCatalog filters remain unexposed.
- Router/server/bootstrap cleanup remains a future blueprint candidate, not a blocker for the next slice.

## NEXT

Execution channel: Terminal Codex.

Next valid slice:

```text
Shared success envelope and ADR 0012 output contract centralization.
```
