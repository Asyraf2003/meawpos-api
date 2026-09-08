-- Copyright (C) 2026 Asyraf Mubarak
-- This file is part of gopos-api and licensed under GNU AGPLv3.

CREATE TABLE catalog_items (
    id uuid PRIMARY KEY,
    root_id uuid NOT NULL REFERENCES roots(id) ON DELETE RESTRICT,
    name text NOT NULL CHECK (btrim(name) <> ''),
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT catalog_items_root_id_id_unique UNIQUE (root_id, id)
);

CREATE INDEX catalog_items_root_id_idx ON catalog_items(root_id, created_at, id);
