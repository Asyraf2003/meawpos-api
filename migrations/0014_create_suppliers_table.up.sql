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

create table suppliers (
  id text primary key,
  name text not null,
  name_normalized text not null,
  phone text null,
  email text null,
  address text null,
  notes text null,
  is_active boolean not null default true,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create unique index suppliers_active_name_normalized_unique
on suppliers (name_normalized)
where is_active = true;

create index suppliers_active_name_idx
on suppliers (is_active, name_normalized, id);

create index suppliers_name_normalized_idx
on suppliers (name_normalized);

create index suppliers_updated_at_idx
on suppliers (updated_at);
