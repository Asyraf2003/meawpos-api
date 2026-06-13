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

# Supplier Implementation Slice 1

## Status

Accepted locally for implementation.

## Active Scope

Implement Supplier domain, ports, and usecase contracts only.

## FACT

- Supplier domain contract is accepted locally with proof.
- Supplier is the next dependency before Faktur.
- Accepted duplicate policy: active supplier name must be unique by normalized name; inactive supplier names do not block active supplier name reuse.
- Accepted permission model: `supplier.read` and `supplier.manage`.
- Product ID remains the downstream Product reference key.
- Supplier must not own ProductCatalog data.
- Supplier must not mutate stock.
- Supplier must not create Faktur.

## SCOPE-IN

- `internal/modules/supplier/domain`
- `internal/modules/supplier/ports`
- `internal/modules/supplier/usecase`
- Domain unit tests.
- Usecase unit tests with in-memory fake repository.

## SCOPE-OUT

- PostgreSQL persistence.
- HTTP transport.
- Presenter/output DTOs.
- Route registration.
- Capability seed migration.
- Supplier API docs.
- Faktur.
- Inventory mutation.
- Stock movement.
- Audit/outbox.
- Localization.
- Extended filters.

## PROOF REQUIRED

Run:

```bash
go test ./internal/modules/supplier/...
bash scripts/audit_hexagonal.sh
make verify
```

## NEXT

After proof, update ledger/handoff.

Next valid active step after this slice:

Supplier PostgreSQL persistence blueprint.
