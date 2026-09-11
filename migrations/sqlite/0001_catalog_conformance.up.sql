-- Copyright (C) 2026 Asyraf Mubarak
-- This file is part of gopos-api and licensed under GNU AGPLv3.

CREATE TABLE roots (
    id TEXT PRIMARY KEY
        CHECK (
            length(id) = 36
            AND id = lower(id)
            AND substr(id, 9, 1) = '-'
            AND substr(id, 14, 1) = '-'
            AND substr(id, 19, 1) = '-'
            AND substr(id, 24, 1) = '-'
            AND substr(id, 1, 8) NOT GLOB '*[^0-9a-f]*'
            AND substr(id, 10, 4) NOT GLOB '*[^0-9a-f]*'
            AND substr(id, 15, 4) NOT GLOB '*[^0-9a-f]*'
            AND substr(id, 20, 4) NOT GLOB '*[^0-9a-f]*'
            AND substr(id, 25, 12) NOT GLOB '*[^0-9a-f]*'
        ),
    name TEXT NOT NULL CHECK (trim(name) <> ''),
    created_at INTEGER NOT NULL CHECK (typeof(created_at) = 'integer'),
    updated_at INTEGER NOT NULL CHECK (typeof(updated_at) = 'integer')
) STRICT;

CREATE TABLE catalog_items (
    id TEXT PRIMARY KEY
        CHECK (
            length(id) = 36
            AND id = lower(id)
            AND substr(id, 9, 1) = '-'
            AND substr(id, 14, 1) = '-'
            AND substr(id, 19, 1) = '-'
            AND substr(id, 24, 1) = '-'
            AND substr(id, 1, 8) NOT GLOB '*[^0-9a-f]*'
            AND substr(id, 10, 4) NOT GLOB '*[^0-9a-f]*'
            AND substr(id, 15, 4) NOT GLOB '*[^0-9a-f]*'
            AND substr(id, 20, 4) NOT GLOB '*[^0-9a-f]*'
            AND substr(id, 25, 12) NOT GLOB '*[^0-9a-f]*'
        ),
    root_id TEXT NOT NULL REFERENCES roots(id) ON DELETE RESTRICT,
    name TEXT NOT NULL CHECK (trim(name) <> ''),
    created_at INTEGER NOT NULL CHECK (typeof(created_at) = 'integer'),
    updated_at INTEGER NOT NULL CHECK (typeof(updated_at) = 'integer'),
    CONSTRAINT catalog_items_root_id_id_unique UNIQUE (root_id, id)
) STRICT;

CREATE INDEX catalog_items_root_id_idx
    ON catalog_items(root_id, created_at, id);

CREATE TABLE catalog_item_prices (
    root_id TEXT NOT NULL,
    catalog_item_id TEXT NOT NULL,
    currency TEXT NOT NULL DEFAULT 'IDR' CHECK (currency = 'IDR'),
    amount_rupiah INTEGER NOT NULL
        CHECK (typeof(amount_rupiah) = 'integer' AND amount_rupiah > 0),
    created_at INTEGER NOT NULL CHECK (typeof(created_at) = 'integer'),
    updated_at INTEGER NOT NULL CHECK (typeof(updated_at) = 'integer'),
    PRIMARY KEY (root_id, catalog_item_id),
    FOREIGN KEY (root_id, catalog_item_id)
        REFERENCES catalog_items(root_id, id) ON DELETE CASCADE
) STRICT;
