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

CREATE TABLE api_capabilities (
    key text PRIMARY KEY,
    domain text NOT NULL,
    operation text NOT NULL,
    method text NOT NULL,
    path text NOT NULL,
    default_enabled boolean NOT NULL,
    enabled boolean NOT NULL,
    required_permission text NOT NULL,
    risk_level text NOT NULL,
    audit_required boolean NOT NULL,
    idempotency_required boolean NOT NULL,
    owner_package text NOT NULL,
    test_proof text NOT NULL,
    disabled_reason text NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT api_capabilities_risk_level_check
        CHECK (risk_level IN ('low', 'medium', 'high')),
    CONSTRAINT api_capabilities_method_check
        CHECK (method IN ('GET', 'POST', 'PUT', 'PATCH', 'DELETE'))
);

CREATE INDEX api_capabilities_domain_operation_idx
    ON api_capabilities (domain, operation);

CREATE INDEX api_capabilities_required_permission_idx
    ON api_capabilities (required_permission);
