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

The next implementation scope is not a source refactor. It is a read-first classification of the current Go runtime before a redesign blueprint is accepted.

## Current Runtime

The repository currently contains Go/Echo/PostgreSQL implementation history for authentication, authorization, capability control, ProductCatalog, ServiceCatalog, and Supplier-related work. Those facts are preserved for the later runtime audit; they are not promises that the same boundaries or domains remain mandatory.

## License

MiawPOS API is licensed under GNU Affero General Public License v3.0 only. See [LICENSE](LICENSE). Third-party notices and attribution remain governed by their applicable terms.
