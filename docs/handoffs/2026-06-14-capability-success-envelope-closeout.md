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

# Capability Success Envelope Closeout

## Status

Closed with local proof.

## FACT

- Capability success responses now use `httpresponse.Success(...)`.
- Capability local success envelope leftovers were removed.
- Capability success test decode support now includes `meta`.
- Capability success tests prove canonical empty `meta`.
- ProductCatalog, ServiceCatalog, and Capability now use the shared success envelope helper.
- Auth/System output contract centralization is deferred by owner decision.

## PROOF

Owner local terminal proof passed:

```text
go test ./internal/transport/http/response/...
go test ./internal/modules/capability/transport/http/...
go test ./internal/modules/productcatalog/transport/http/... ./internal/modules/servicecatalog/transport/http/... ./internal/modules/system/transport/http/... ./internal/modules/auth/transport/http/...
bash scripts/audit_hexagonal.sh
bash scripts/audit_route_capabilities.sh
make verify

[PASS] focused response and Capability transport tests
[PASS] ProductCatalog, ServiceCatalog, System, and Auth transport tests
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit
[PASS] aggregate audit passed
```

## OWNER DECISION

Auth/System output contract centralization is deferred.

Reason:

```text
Auth currently has enough foundation for the next Product/Supplier/Faktur path when tests, PostgreSQL data, and middleware gates pass.
Auth token response and System health/me response shapes are higher-risk public contracts.
Full ADR 0012 closeout must not block Product API readiness, Supplier, or Faktur planning.
```

## GAP

ADR `0012` remains partial.

Remaining output-contract gaps:

```text
Auth success responses still return raw success DTOs or 204 No Content.
System /api/me and /api/health still return raw presenter/map responses.
Full response/error envelope coverage is not proven for every API surface.
```

## NEXT

Next valid active step:

```text
Product API readiness for Supplier and Faktur path.
```

Scope guard:

```text
Do not start inventory mutation, audit/outbox, localization, extended filters, Supplier, or Faktur implementation in the readiness slice.
```
