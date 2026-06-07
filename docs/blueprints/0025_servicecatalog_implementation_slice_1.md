# ServiceCatalog Implementation Slice 1 Blueprint

## Status

Accepted implementation slice plan.

## Date

2026-06-08

## Active Scope

Plan ServiceCatalog implementation slice 1 after accepted domain contract.

## Accepted Contract

```text
docs/blueprints/0024_servicecatalog_domain_contract.md
```

## FACT

- ServiceCatalog domain contract is accepted.
- Capability-control foundation is closed.
- ServiceCatalog implementation has not started.
- Business Phase 1 implementation remains 0%.
- Protected endpoints still require authn, authz, capability check, validation, then usecase execution.
- This slice is planning only.

## DECISION

Slice 1 should implement only ServiceCatalog domain and usecase contracts with tests.

Do not implement HTTP transport yet.

Do not implement PostgreSQL adapter yet.

Do not add migrations yet.

Do not add capability seeds yet.

## SCOPE-IN

- `internal/modules/servicecatalog/domain`
- `internal/modules/servicecatalog/ports`
- `internal/modules/servicecatalog/usecase`
- Domain normalization behavior
- Domain validation
- Usecase command/result contracts
- In-memory fake repositories for tests
- Unit tests for create, update, activate, deactivate, show, list, lookup behavior

## SCOPE-OUT

- Echo handlers
- Route registration
- PostgreSQL migrations
- PostgreSQL repositories
- Capability seed migrations
- Route-to-capability audit manifest changes
- UI
- ProductCatalog
- Inventory
- Audit sink implementation

## TARGET PACKAGE PLAN

```text
internal/modules/servicecatalog/
  domain/
    service_catalog_item.go
    normalizer.go
    errors.go
    validation.go
  ports/
    service_catalog_repository.go
  usecase/
    create_item.go
    update_item.go
    activate_item.go
    deactivate_item.go
    show_item.go
    list_items.go
    lookup_items.go
```

## DOMAIN TYPES

- `ServiceCatalogItem`
- `ServiceCatalogItemStatus`
- `ServiceCatalogItemID`
- `NormalizedName`
- `MoneyRupiah`

Minimum fields:

- `id`
- `name`
- `normalized_name`
- `default_price_rupiah`
- `is_active`
- `created_at`
- `updated_at`

## DOMAIN RULES

- `name` is required after trim.
- Repeated internal whitespace is compacted.
- Normalized name is lowercase.
- Default price must be greater than zero.
- New item is active by default.
- Deactivate keeps item stored.
- Activate restores active status.
- Physical delete is forbidden.

## PORTS

Repository port should support:

- `Create`
- `Update`
- `FindByID`
- `FindByNormalizedName`
- `List`
- `Lookup`
- `SetActive`

Transactor port is deferred unless usecase tests need explicit transaction boundary in this slice.

## USECASES

- `CreateServiceCatalogItem`
- `UpdateServiceCatalogItem`
- `ActivateServiceCatalogItem`
- `DeactivateServiceCatalogItem`
- `ShowServiceCatalogItem`
- `ListServiceCatalogItems`
- `LookupServiceCatalogItems`

## TEST MATRIX

Domain tests:

- Normalize trims whitespace.
- Normalize compacts repeated whitespace.
- Normalize lowercases.
- Reject blank name.
- Reject zero default price.
- Reject negative default price.
- Create active item by default.

Usecase tests:

- Create stores item.
- Create rejects duplicate normalized name.
- Update changes name and price.
- Update rejects duplicate normalized name.
- Activate marks inactive item active.
- Deactivate marks active item inactive.
- Show missing item returns not found.
- List filters active/inactive/all.
- Lookup excludes inactive by default.
- Lookup enforces max limit.

## PROOF REQUIRED

- `go test ./internal/modules/servicecatalog/...`
- `make verify`

## NEXT ACTIVE STEP

ServiceCatalog implementation slice 1 is accepted.

Next valid active step:

```text
Implement ServiceCatalog slice 1: domain, ports, usecase contracts, and unit tests only.
```

Do not implement HTTP transport, PostgreSQL adapter, migrations, route registration, or capability seeds in this slice.

## ACCEPTANCE

ServiceCatalog implementation slice 1 accepted on 2026-06-08.

Accepted implementation scope:

```text
internal/modules/servicecatalog/domain
internal/modules/servicecatalog/ports
internal/modules/servicecatalog/usecase
unit tests for domain and usecase behavior
```

Forbidden in this slice:

```text
HTTP transport
PostgreSQL adapter
PostgreSQL migrations
route registration
capability seed migrations
ProductCatalog
Inventory
```

Proof required after implementation:

```text
go test ./internal/modules/servicecatalog/...
make verify
```
