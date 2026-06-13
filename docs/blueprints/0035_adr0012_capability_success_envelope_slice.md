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

# ADR 0012 Capability Success Envelope Slice

## Status

Accepted locally for implementation.

## Active Scope

Normalize Capability success responses to the shared HTTP success envelope.

This slice advances ADR `0012` output contract centralization without changing Auth/System response shapes and without claiming full ADR closeout.

## FACT

- ADR `0012` requires centralized API output contracts and staged refactors.
- The canonical success envelope is `success`, `data`, and `meta`.
- Shared success envelope primitives already exist in `internal/transport/http/response`.
- ProductCatalog and ServiceCatalog already use `httpresponse.Success(...)`.
- Capability currently returns `success` and `data` without `meta`.
- Capability already uses presentation DTOs from `internal/presentation/http/id/capability`.

## GAP

- Capability still has a local success response envelope.
- Capability success responses do not include `meta`.
- Capability success tests do not prove canonical `meta` output.
- Auth and System response centralization remains incomplete and is intentionally out of this slice.

## DECISION

Implement the smallest safe ADR `0012` follow-up:

- Replace Capability local success envelope usage with `httpresponse.Success(...)`.
- Remove the local Capability `responseEnvelope` type.
- Update Capability success test decode support to require `meta` as an empty object.
- Preserve Capability presenter output.
- Preserve Capability error behavior.
- Preserve Capability route/capability middleware behavior.
- Do not change Auth success DTOs.
- Do not change Auth `204 No Content` responses.
- Do not change System `/api/me` or `/api/health` responses in this slice.
- Do not claim ADR `0012` fully proven.

## SCOPE-IN

- `internal/modules/capability/transport/http/capability_handler_response.go`
- `internal/modules/capability/transport/http/capability_handler_read.go`
- `internal/modules/capability/transport/http/capability_handler_write.go`
- `internal/modules/capability/transport/http/capability_handler_test_decode_test.go`
- Focused Capability transport tests.
- Existing response package tests.
- Existing route capability audit.
- Existing aggregate verify gate.

## SCOPE-OUT

- Auth response shape changes.
- System response shape changes.
- Error envelope behavior changes.
- Capability middleware changes.
- Capability domain/usecase/ports changes.
- Route/capability manifest changes.
- Migrations.
- Inventory/stock/audit/outbox/localization/extended filters.

## RESPONSE CONTRACT

Capability success responses must use:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

Capability errors continue through the existing shared error handler.

## PATCH PLAN

- Import `httpresponse "pos-go/internal/transport/http/response"` in Capability read/write handlers.
- Replace local `responseEnvelope{Success: true, Data: ...}` with `httpresponse.Success(...)`.
- Remove the local `responseEnvelope` type from `capability_handler_response.go`.
- Update Capability test envelope decoder to include and assert empty `meta`.
- Run focused tests and audits.
- Keep ADR `0012` partial until Auth/System output contract coverage is handled by a later accepted slice.

## PROOF REQUIRED

Run:

```bash
go test ./internal/transport/http/response/...
go test ./internal/modules/capability/transport/http/...
go test ./internal/modules/productcatalog/transport/http/... ./internal/modules/servicecatalog/transport/http/... ./internal/modules/system/transport/http/... ./internal/modules/auth/transport/http/...
bash scripts/audit_hexagonal.sh
bash scripts/audit_route_capabilities.sh
make verify
```

## NEXT

After proof, update:

```text
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/evidence/0004_adr_implementation_proof_index.md
docs/handoffs/
```

Next remaining ADR `0012` scope after this slice:

```text
Auth/System output contract coverage analysis and blueprint.
```
