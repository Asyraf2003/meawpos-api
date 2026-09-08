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

R4's PostgreSQL walking skeleton and R5's reproducible security/release gate are implemented and proven. R6 SQLite conformance is the exact next scope but has not started; broader capability expansion remains later.

## Current Runtime

The default business composition is `catalog.core,catalog.pricing,sales,payment.cash`. Set `BUSINESS_COMPONENTS` to an explicit comma-separated set, or `none`, at startup. Missing dependencies and unknown component names fail startup; trusted ROOT, authorization, money, transaction, audit, and idempotency responsibilities are not business toggles.

The new foundation API is ROOT-scoped under `/api/roots/{root_id}/...`. Existing ProductCatalog, ServiceCatalog, Supplier, and operation-capability behavior remains a separate compatibility surface.

## Release Gate

After configuring and migrating the local PostgreSQL target, run:

```text
make release-gate
```

This single gate runs the existing repository checks plus pinned gosec,
govulncheck, Gitleaks worktree/history scans, and the PostgreSQL-backed
authentication, authorization, ROOT-isolation, operation-authority, and API
abuse suite. An uncached scanner run and vulnerability check require network
access. `DATABASE_URL` may be exported or loaded from `.env`; missing or
unmigrated PostgreSQL is a gate failure rather than a skipped security proof.

## License

MiawPOS API is licensed under GNU Affero General Public License v3.0 only. See [LICENSE](LICENSE). Third-party notices and attribution remain governed by their applicable terms.
