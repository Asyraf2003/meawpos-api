# Handoff: Capability Admin HTTP Surface

## Date

2026-06-07

## Active Scope

Add admin capability HTTP surface for `docs/blueprints/0010_capability_control_foundation.md` step 6.

## Files Changed

```text
migrations/0008_seed_capability_manage_permission.up.sql
migrations/0008_seed_capability_manage_permission.down.sql
internal/presentation/http/id/capability/capability.go
internal/modules/capability/transport/http/capability_handler.go
internal/modules/capability/transport/http/capability_handler_read.go
internal/modules/capability/transport/http/capability_handler_write.go
internal/modules/capability/transport/http/capability_handler_response.go
internal/modules/capability/transport/http/capability_handler_test.go
internal/modules/capability/transport/http/capability_handler_read_test.go
internal/modules/capability/transport/http/capability_handler_show_error_test.go
internal/modules/capability/transport/http/capability_handler_write_test.go
internal/modules/capability/transport/http/capability_handler_write_error_test.go
internal/modules/capability/transport/http/capability_handler_test_assert_test.go
internal/modules/capability/transport/http/capability_handler_test_context_test.go
internal/modules/capability/transport/http/capability_handler_test_decode_test.go
internal/modules/capability/transport/http/capability_handler_test_fake_test.go
internal/modules/capability/transport/http/capability_handler_test_fake_show_test.go
internal/modules/capability/transport/http/capability_handler_test_fake_write_test.go
internal/app/bootstrap/app.go
docs/handoffs/2026-06-07-capability-admin-http-surface.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
```

## Implementation Facts

- Migration `0008` seeds permission `capability.manage`.
- Migration `0008` assigns `capability.manage` only to role `admin`.
- Migration `0008` seeds `api_capabilities.key = 'capability.manage'` for `/api/admin/capabilities` with method `*`, required permission `capability.manage`, high risk, audit required, and owner package `internal/modules/capability/transport/http`.
- Migration `0008` extends `api_capabilities_method_check` to allow `*` because the admin capability control surface uses one aggregate capability key for list/show/enable/disable routes.
- Capability DTO mapping exists in `internal/presentation/http/id/capability/capability.go`.
- Admin capability handler exists in split files under `internal/modules/capability/transport/http/`.
- Registered handler routes are `GET /capabilities`, `GET /capabilities/:key`, `POST /capabilities/:key/enable`, and `POST /capabilities/:key/disable` on the provided Echo group.
- Handler tests use fake use cases and do not require PostgreSQL.
- Bootstrap wires the capability PostgreSQL repository, list/show/enable/disable use cases, check use case, and admin capability handler.
- Bootstrap protects `/api/admin/capabilities...` with `RequireAuth`, `RequirePermission("capability.manage")`, and `RequireCapability("capability.manage", checkCapabilityUsecase)`.
- Existing account-role route behavior remains on its separate `/api/admin` group with permission `account.role.assign`.
- Handler and handler test files were split so non-allowlisted files remain below the 100-line file-size audit limit.

## Proof Collected

User-provided proof before file-size fix:

- `make dev` applied migration `0008` successfully.
- SQL proof showed permission `capability.manage` exists.
- SQL proof showed admin role has `capability.manage`.
- SQL proof showed `api_capabilities` has `capability.manage`.
- Targeted tests passed for `go test ./internal/modules/capability/...`, `go test ./internal/modules/capability/transport/http/...`, and `go test ./internal/app/bootstrap/...`.
- Initial `make verify` failed only at file-size audit for oversized capability handler files.

Local proof after file-size fix:

```text
wc -l internal/modules/capability/transport/http/*.go
all files <= 100 lines

env GOCACHE=/tmp/go-build-cache go test ./internal/modules/capability/transport/http/...
ok  	pos-go/internal/modules/capability/transport/http	0.007s

env GOCACHE=/tmp/go-build-cache go test ./internal/modules/capability/...
ok  	pos-go/internal/modules/capability/domain	(cached)
?   	pos-go/internal/modules/capability/ports	[no test files]
ok  	pos-go/internal/modules/capability/transport/http	(cached)
ok  	pos-go/internal/modules/capability/usecase	(cached)

env GOCACHE=/tmp/go-build-cache go test ./internal/app/bootstrap/...
ok  	pos-go/internal/app/bootstrap	0.204s

make verify
[PASS] file size audit passed
[PASS] hexagonal import audit passed
[PASS] gosec audit passed
[PASS] aggregate audit passed
```

## Remaining Gaps

- Route-to-capability audit script remains out of scope and not implemented.
- Route-level disabled protected endpoint proof remains open unless covered by later proof.
- POS CRUD remains blocked until capability-control foundation proof is complete.

## Next Valid Step

Move to route-to-capability audit script work only after user accepts this step's proof.
