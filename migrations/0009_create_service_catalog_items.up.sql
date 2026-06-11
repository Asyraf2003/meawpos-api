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

CREATE TABLE service_catalog_items (
    id text PRIMARY KEY,
    name text NOT NULL,
    normalized_name text NOT NULL,
    default_price_rupiah bigint NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT service_catalog_items_default_price_positive_check
        CHECK (default_price_rupiah > 0)
);

CREATE UNIQUE INDEX service_catalog_items_normalized_name_unique
    ON service_catalog_items (normalized_name);

CREATE INDEX service_catalog_items_active_name_idx
    ON service_catalog_items (is_active, normalized_name);
