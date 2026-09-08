-- Copyright (C) 2026 Asyraf Mubarak
-- This file is part of gopos-api and licensed under GNU AGPLv3.

CREATE TABLE sales (
    id uuid PRIMARY KEY,
    root_id uuid NOT NULL REFERENCES roots(id) ON DELETE RESTRICT,
    status text NOT NULL CHECK (status IN ('POSTED', 'REVERSED')),
    total_rupiah bigint NOT NULL CHECK (total_rupiah > 0),
    actor_account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT,
    session_id uuid NOT NULL,
    posted_at timestamptz NOT NULL,
    reversed_at timestamptz NULL,
    CONSTRAINT sales_root_id_id_unique UNIQUE (root_id, id),
    CONSTRAINT sales_reversal_time_check CHECK (
        (status = 'POSTED' AND reversed_at IS NULL)
        OR (status = 'REVERSED' AND reversed_at IS NOT NULL)
    )
);

CREATE TABLE sale_lines (
    id uuid PRIMARY KEY,
    root_id uuid NOT NULL,
    sale_id uuid NOT NULL,
    catalog_item_id uuid NOT NULL,
    item_name_snapshot text NOT NULL CHECK (btrim(item_name_snapshot) <> ''),
    unit_price_rupiah bigint NOT NULL CHECK (unit_price_rupiah > 0),
    quantity bigint NOT NULL CHECK (quantity > 0),
    line_total_rupiah bigint NOT NULL CHECK (line_total_rupiah > 0),
    created_at timestamptz NOT NULL,
    FOREIGN KEY (root_id, sale_id) REFERENCES sales(root_id, id) ON DELETE RESTRICT,
    FOREIGN KEY (root_id, catalog_item_id) REFERENCES catalog_items(root_id, id) ON DELETE RESTRICT
);

CREATE TABLE sale_reversals (
    id uuid PRIMARY KEY,
    root_id uuid NOT NULL,
    sale_id uuid NOT NULL,
    actor_account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT,
    session_id uuid NOT NULL,
    reason text NOT NULL CHECK (btrim(reason) <> ''),
    reversed_at timestamptz NOT NULL,
    CONSTRAINT sale_reversals_sale_unique UNIQUE (root_id, sale_id),
    CONSTRAINT sale_reversals_root_id_id_unique UNIQUE (root_id, id),
    FOREIGN KEY (root_id, sale_id) REFERENCES sales(root_id, id) ON DELETE RESTRICT
);

CREATE INDEX sales_root_posted_idx ON sales(root_id, posted_at, id);
CREATE INDEX sale_lines_sale_idx ON sale_lines(root_id, sale_id, id);
