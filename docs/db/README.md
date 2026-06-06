# Database

This folder defines PostgreSQL and persistence rules.

## Contents

- `0040_postgresql_policy.md`: schema, migration, transaction, repository, and data-integrity policy.

## Use This Folder When

- adding or changing migrations;
- designing repositories or query adapters;
- deciding transaction boundaries;
- defining database-owned invariants and constraints.

Database-backed mutations must be covered by domain contracts, capability rules, transactions, and test proof.
