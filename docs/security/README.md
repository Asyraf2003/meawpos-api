# Security

This folder defines baseline security rules.

## Contents

- `0080_security_baseline.md`: authentication, authorization, capability, input, output, secret, audit, and abuse-control rules.

## Use This Folder When

- adding protected APIs;
- changing authn, authz, roles, permissions, or capabilities;
- handling secrets, tokens, input validation, or audit decisions;
- deciding whether a mutation is safe to expose.

Security rules are P0 constraints when they intersect with API exposure or state mutation.
