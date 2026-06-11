-- Copyright (C) 2026 Asyraf Mubarak
--
-- This file is part of gopos-api.
--
-- gopos-api is free software: you can redistribute it and/or modify
-- it under the terms of the GNU Affero General Public License as published by
-- the Free Software Foundation, version 3 only.
--
-- gopos-api is distributed in the hope that it will be useful,
-- but WITHOUT ANY WARRANTY; without even the implied warranty of
-- MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
-- GNU Affero General Public License for more details.
--
-- You should have received a copy of the GNU Affero General Public License
-- along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

CREATE TABLE products (
    id text PRIMARY KEY,
    kode_barang text NULL,
    nama_barang text NOT NULL,
    nama_barang_normalized text NOT NULL,
    merek text NOT NULL,
    merek_normalized text NOT NULL,
    ukuran integer NULL,
    harga_jual bigint NOT NULL,
    reorder_point_qty integer NULL,
    critical_threshold_qty integer NULL,
    deleted_at timestamptz NULL,
    deleted_by_actor_id text NULL,
    delete_reason text NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT products_harga_jual_positive_check
        CHECK (harga_jual > 0),
    CONSTRAINT products_reorder_point_non_negative_check
        CHECK (reorder_point_qty IS NULL OR reorder_point_qty >= 0),
    CONSTRAINT products_critical_threshold_non_negative_check
        CHECK (critical_threshold_qty IS NULL OR critical_threshold_qty >= 0),
    CONSTRAINT products_threshold_pair_check
        CHECK (
            (reorder_point_qty IS NULL AND critical_threshold_qty IS NULL)
            OR
            (reorder_point_qty IS NOT NULL AND critical_threshold_qty IS NOT NULL)
        ),
    CONSTRAINT products_threshold_order_check
        CHECK (
            critical_threshold_qty IS NULL
            OR reorder_point_qty IS NULL
            OR critical_threshold_qty <= reorder_point_qty
        )
);

CREATE TABLE product_versions (
    id text PRIMARY KEY,
    product_id text NOT NULL,
    revision_no integer NOT NULL,
    event_name text NOT NULL,
    changed_by_actor_id text NULL,
    change_reason text NULL,
    changed_at timestamptz NOT NULL,
    snapshot_json jsonb NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT fk_product_versions_product
        FOREIGN KEY (product_id)
        REFERENCES products (id)
        ON DELETE RESTRICT
);

CREATE INDEX products_merek_idx
    ON products (merek);

CREATE INDEX products_ukuran_idx
    ON products (ukuran);

CREATE INDEX products_harga_jual_idx
    ON products (harga_jual);

CREATE INDEX products_duplicate_lookup_idx
    ON products (nama_barang, merek, ukuran);

CREATE INDEX products_deleted_at_idx
    ON products (deleted_at);

CREATE INDEX products_nama_barang_normalized_idx
    ON products (nama_barang_normalized);

CREATE INDEX products_merek_normalized_idx
    ON products (merek_normalized);

CREATE UNIQUE INDEX products_kode_barang_unique
    ON products (kode_barang)
    WHERE deleted_at IS NULL AND kode_barang IS NOT NULL;

CREATE INDEX products_active_list_idx
    ON products (nama_barang_normalized, merek_normalized, ukuran, id)
    WHERE deleted_at IS NULL;

CREATE INDEX products_active_identity_lookup_idx
    ON products (nama_barang_normalized, merek_normalized, ukuran, kode_barang, id)
    WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX product_versions_product_revision_unique
    ON product_versions (product_id, revision_no);

CREATE INDEX product_versions_product_changed_at_idx
    ON product_versions (product_id, changed_at);

CREATE INDEX product_versions_event_name_idx
    ON product_versions (event_name);
