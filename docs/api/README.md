# API

This folder defines public HTTP API contract rules.

## Contents

- `0050_echo_http_contract.md`: Echo transport rules, request/response ownership, envelope expectations, route behavior, and forbidden HTTP-layer shortcuts.

## Use This Folder When

- adding or changing an HTTP endpoint;
- defining response envelopes or public DTO behavior;
- mapping application errors to public API errors;
- checking whether transport code is leaking domain or persistence logic.

API contracts are public contracts. Changes here usually need blueprint coverage, focused API tests, and sometimes an ADR.
