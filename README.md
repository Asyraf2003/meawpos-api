# MiawPOS API

MiawPOS API is an open-source, AGPL-licensed, self-host-first, API-only point-of-sale and business foundation written in Go with Echo.

The project is designed to remain lightweight and locally usable while preserving strict financial correctness, security, authorization, transaction integrity, and audit truth. Business capabilities may be composed where their real dependencies allow it; the trusted truth boundary is not an ordinary feature toggle.

GlassPOS is design DNA and historical reference lineage. It is not the active product name and its SaaS, PAYG, Laravel, or fixed-domain assumptions do not define future MiawPOS architecture.

## Current Architecture Direction

```text
PostgreSQL = primary/default server persistence
SQLite     = first-class portability/local/offline-standalone target
MySQL      = not an active target
```

The first walking skeleton is intentionally narrow:

```text
minimal catalog
-> sale
-> cash payment
-> financial trust boundary
-> exact financial calculation
-> transactional persistence
-> audit/history truth
-> API response/readback
```

SQLite support will be added through proven persistence boundaries. No parity claim exists yet.

## Start Here

Read [docs/README.md](docs/README.md), then follow its canonical read order. The documentation lives in a Git submodule and has a separate change/proof surface from this parent repository.

R4's PostgreSQL walking skeleton is implemented and proven. The exact next scope is R5's reproducible security and release gate; R6 SQLite conformance and broader capability expansion remain later.

## Current Runtime

The default business composition is `catalog.core,catalog.pricing,sales,payment.cash`. Set `BUSINESS_COMPONENTS` to an explicit comma-separated set, or `none`, at startup. Missing dependencies and unknown component names fail startup; trusted ROOT, authorization, money, transaction, audit, and idempotency responsibilities are not business toggles.

The new foundation API is ROOT-scoped under `/api/roots/{root_id}/...`. Existing ProductCatalog, ServiceCatalog, Supplier, and operation-capability behavior remains a separate compatibility surface.

## License

MiawPOS API is licensed under GNU Affero General Public License v3.0 only. See [LICENSE](LICENSE). Third-party notices and attribution remain governed by their applicable terms.
