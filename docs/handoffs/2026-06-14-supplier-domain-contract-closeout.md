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

# Supplier Domain Contract Closeout

## Status

Accepted with local proof.

## FACT

- Supplier domain contract blueprint exists at `docs/blueprints/0037_supplier_domain_contract.md`.
- Supplier is accepted as the next domain dependency before Faktur.
- Supplier is master data for future external party references.
- Future Faktur may reference Supplier through `supplier_id`.
- Supplier must not own ProductCatalog data.
- Supplier must not mutate stock.
- Supplier must not create Faktur.
- Product ID remains the downstream Product reference key.
- Auth/System output contract centralization remains deferred by owner decision.

## OWNER DECISIONS

Accepted duplicate policy:

```text
Active Supplier name must be unique by normalized name.
Inactive Supplier names do not block active Supplier name reuse.
```

Accepted permission model:

```text
supplier.read
supplier.manage
```

Accepted implementation order:

```text
Supplier before Faktur.
```

## PROOF

Owner local terminal proof passed:

```text
rg -n 'Supplier Domain Contract Blueprint|DUPLICATE POLICY|POSTGRESQL SCHEMA PROPOSAL|API CONTRACT PROPOSAL|AUTHORIZATION POLICY|RELATION TO FAKTUR|ACCEPTANCE GATE|Supplier implementation slice 1' docs/blueprints/0037_supplier_domain_contract.md
make verify

[PASS] blueprint marker check
[PASS] go test ./...
[PASS] go vet audit
[PASS] format audit
[PASS] AI rules audit
[PASS] license header audit
[PASS] file size audit
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit
[PASS] aggregate audit passed
```

## GAP

- Supplier implementation is not started.
- Supplier PostgreSQL migration is not started.
- Supplier HTTP routes are not started.
- Supplier capability seed migration is not started.
- Faktur domain contract is not accepted yet.
- Faktur implementation is not started.
- Inventory/stock mutation remains incomplete.
- Audit/outbox persistence remains incomplete.
- Runtime localization remains incomplete.
- Extended filters remain incomplete.

## NEXT

Next valid active step:

```text
Supplier implementation slice 1: domain, ports, and usecase contracts only.
```

Scope guard:

```text
Do not implement PostgreSQL persistence, HTTP routes, capability seed migration, Faktur, inventory mutation, stock movement, audit/outbox, localization, extended filters, or architecture folder rename work in Supplier implementation slice 1.
```
