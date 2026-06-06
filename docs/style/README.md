# Style

This folder defines Go code style and forbidden implementation patterns.

## Contents

- `0100_go_style.md`: naming, dependency, context, error, DTO, and forbidden-pattern rules.

## Use This Folder When

- naming use cases, repositories, commands, queries, DTOs, or handlers;
- deciding where dependencies should be passed;
- reviewing Go code for framework leakage or hidden globals;
- checking whether a helper or package is becoming too broad.

Style rules do not override architecture or security, but they keep implementation predictable.
