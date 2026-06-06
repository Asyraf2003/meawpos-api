# Domain

This folder defines domain contract rules.

## Contents

- `0030_domain_contracts.md`: required create, update, delete, show, list, authorization, capability, audit, transaction, and test declarations for database-backed domains.

## Use This Folder When

- creating a new business domain;
- adding CRUD behavior;
- forbidding or constraining a domain operation;
- mapping domain lifecycle rules before API or database implementation.

Do not implement domain CRUD without an explicit domain contract and capability keys.
