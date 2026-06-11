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
VALUES
    (
        'profile.self.show',
        'profile',
        'show_self',
        'GET',
        '/api/me',
        true,
        true,
        'profile.self.read',
        'low',
        false,
        false,
        'internal/modules/system/transport/http',
        'migrations/0007_seed_existing_protected_capabilities.up.sql plus repository/integration seed verification',
        NULL
    ),
    (
        'authz.profile.self.show',
        'authz',
        'show_self',
        'GET',
        '/api/authz/me',
        true,
        true,
        'profile.self.read',
        'low',
        false,
        false,
        'internal/modules/system/transport/http',
        'migrations/0007_seed_existing_protected_capabilities.up.sql plus repository/integration seed verification',
        NULL
    ),
    (
        'auth.session.logout',
        'auth',
        'logout',
        'POST',
        '/api/auth/logout',
        true,
        true,
        'auth.session.logout',
        'medium',
        true,
        false,
        'internal/modules/auth/transport/http',
        'migrations/0007_seed_existing_protected_capabilities.up.sql plus repository/integration seed verification',
        NULL
    ),
    (
        'account.role.assign',
        'account',
        'assign_role',
        'POST',
        '/api/admin/accounts/:account_id/roles',
        true,
        true,
        'account.role.assign',
        'high',
        true,
        false,
        'internal/modules/auth/transport/http',
        'migrations/0007_seed_existing_protected_capabilities.up.sql plus repository/integration seed verification',
        NULL
    ),
    (
        'account.role.remove',
        'account',
        'remove_role',
        'DELETE',
        '/api/admin/accounts/:account_id/roles/:role_key',
        true,
        true,
        'account.role.assign',
        'high',
        true,
        false,
        'internal/modules/auth/transport/http',
        'migrations/0007_seed_existing_protected_capabilities.up.sql plus repository/integration seed verification',
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
