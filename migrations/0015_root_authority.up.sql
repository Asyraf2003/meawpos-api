-- Copyright (C) 2026 Asyraf Mubarak
--
-- This file is part of gopos-api.
--
-- gopos-api is free software: you can redistribute it and/or modify
-- it under the terms of the GNU Affero General Public License as published by
-- the Free Software Foundation, version 3 only.

CREATE TABLE roots (
    id uuid PRIMARY KEY,
    name text NOT NULL CHECK (btrim(name) <> ''),
    primary_owner_account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

CREATE TABLE root_memberships (
    id uuid PRIMARY KEY,
    root_id uuid NOT NULL REFERENCES roots(id) ON DELETE CASCADE,
    account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT,
    created_at timestamptz NOT NULL,
    CONSTRAINT root_memberships_root_account_unique UNIQUE (root_id, account_id),
    CONSTRAINT root_memberships_root_id_id_unique UNIQUE (root_id, id)
);

CREATE TABLE root_roles (
    id uuid PRIMARY KEY,
    root_id uuid NOT NULL REFERENCES roots(id) ON DELETE CASCADE,
    key text NOT NULL CHECK (btrim(key) <> ''),
    name text NOT NULL CHECK (btrim(name) <> ''),
    created_at timestamptz NOT NULL,
    CONSTRAINT root_roles_root_key_unique UNIQUE (root_id, key),
    CONSTRAINT root_roles_root_id_id_unique UNIQUE (root_id, id)
);

CREATE TABLE root_role_permissions (
    root_id uuid NOT NULL,
    role_id uuid NOT NULL,
    permission_key text NOT NULL CHECK (btrim(permission_key) <> ''),
    created_at timestamptz NOT NULL,
    PRIMARY KEY (root_id, role_id, permission_key),
    FOREIGN KEY (root_id, role_id) REFERENCES root_roles(root_id, id) ON DELETE CASCADE
);

CREATE TABLE root_membership_roles (
    root_id uuid NOT NULL,
    membership_id uuid NOT NULL,
    role_id uuid NOT NULL,
    created_at timestamptz NOT NULL,
    PRIMARY KEY (root_id, membership_id, role_id),
    FOREIGN KEY (root_id, membership_id) REFERENCES root_memberships(root_id, id) ON DELETE CASCADE,
    FOREIGN KEY (root_id, role_id) REFERENCES root_roles(root_id, id) ON DELETE CASCADE
);

CREATE INDEX root_memberships_account_id_idx ON root_memberships(account_id);
CREATE INDEX root_membership_roles_role_id_idx ON root_membership_roles(role_id);
