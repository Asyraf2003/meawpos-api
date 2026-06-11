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

ALTER TABLE api_capabilities
    DROP CONSTRAINT api_capabilities_method_check;

ALTER TABLE api_capabilities
    ADD CONSTRAINT api_capabilities_method_check
        CHECK (method IN ('GET', 'POST', 'PUT', 'PATCH', 'DELETE', '*'));

INSERT INTO permissions (key, name)
VALUES ('capability.manage', 'Manage API capabilities')
ON CONFLICT (key) DO UPDATE SET
    name = EXCLUDED.name;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key = 'capability.manage'
WHERE r.key = 'admin'
ON CONFLICT (role_id, permission_id) DO NOTHING;

INSERT INTO api_capabilities (
    key,
    domain,
    operation,
    method,
    path,
    default_enabled,
    enabled,
    required_permission,
    risk_level,
    audit_required,
    idempotency_required,
    owner_package,
    test_proof,
    disabled_reason
)
VALUES (
    'capability.manage',
    'capability',
    'manage',
    '*',
    '/api/admin/capabilities',
    true,
    true,
    'capability.manage',
    'high',
    true,
    false,
    'internal/modules/capability/transport/http',
    'internal/modules/capability/transport/http/capability_handler_test.go plus SQL proof placeholders',
    NULL
)
ON CONFLICT (key) DO UPDATE SET
    domain = EXCLUDED.domain,
    operation = EXCLUDED.operation,
    method = EXCLUDED.method,
    path = EXCLUDED.path,
    default_enabled = EXCLUDED.default_enabled,
    enabled = EXCLUDED.enabled,
    required_permission = EXCLUDED.required_permission,
    risk_level = EXCLUDED.risk_level,
    audit_required = EXCLUDED.audit_required,
    idempotency_required = EXCLUDED.idempotency_required,
    owner_package = EXCLUDED.owner_package,
    test_proof = EXCLUDED.test_proof,
    disabled_reason = EXCLUDED.disabled_reason,
    updated_at = now();
