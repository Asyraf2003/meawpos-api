# Testing

This folder defines test and quality-gate rules.

## Contents

- `0060_test_and_quality_gates.md`: unit, integration, API, contract, architecture, migration, and script proof expectations.

## Use This Folder When

- planning proof for a blueprint step;
- adding or changing use cases, repositories, handlers, migrations, or contracts;
- deciding whether `go test`, API tests, DB tests, or architecture checks are required.

Tests are delivery gates. Do not claim completion without visible proof appropriate to the changed surface.
