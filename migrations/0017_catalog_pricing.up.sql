-- Copyright (C) 2026 Asyraf Mubarak
-- This file is part of gopos-api and licensed under GNU AGPLv3.

CREATE TABLE catalog_item_prices (
    root_id uuid NOT NULL,
    catalog_item_id uuid NOT NULL,
    currency text NOT NULL DEFAULT 'IDR' CHECK (currency = 'IDR'),
    amount_rupiah bigint NOT NULL CHECK (amount_rupiah > 0),
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    PRIMARY KEY (root_id, catalog_item_id),
    FOREIGN KEY (root_id, catalog_item_id)
        REFERENCES catalog_items(root_id, id) ON DELETE CASCADE
);
