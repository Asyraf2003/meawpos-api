-- Copyright (C) 2026 Asyraf Mubarak
-- This file is part of gopos-api and licensed under GNU AGPLv3.

CREATE TABLE financial_idempotency (
    root_id uuid NOT NULL REFERENCES roots(id) ON DELETE RESTRICT,
    operation text NOT NULL CHECK (btrim(operation) <> ''),
    idempotency_key text NOT NULL CHECK (btrim(idempotency_key) <> ''),
    request_fingerprint text NOT NULL CHECK (btrim(request_fingerprint) <> ''),
    resource_type text NULL,
    resource_id uuid NULL,
    created_at timestamptz NOT NULL,
    PRIMARY KEY (root_id, operation, idempotency_key),
    CONSTRAINT financial_idempotency_binding_check CHECK (
        (resource_type IS NULL AND resource_id IS NULL)
        OR (resource_type IS NOT NULL AND resource_id IS NOT NULL)
    )
);

CREATE TABLE audit_events (
    id uuid PRIMARY KEY,
    root_id uuid NOT NULL REFERENCES roots(id) ON DELETE RESTRICT,
    actor_account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT,
    session_id uuid NOT NULL,
    request_id text NOT NULL CHECK (btrim(request_id) <> ''),
    authority_used text NOT NULL CHECK (btrim(authority_used) <> ''),
    operation text NOT NULL CHECK (btrim(operation) <> ''),
    resource_type text NOT NULL CHECK (btrim(resource_type) <> ''),
    resource_id uuid NOT NULL,
    reason text NULL,
    occurred_at timestamptz NOT NULL
);

CREATE INDEX audit_events_resource_idx
    ON audit_events(root_id, resource_type, resource_id, occurred_at, id);
