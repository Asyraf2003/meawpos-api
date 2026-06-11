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

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key IN (
    'auth.session.refresh',
    'auth.session.logout',
    'profile.self.read'
)
WHERE r.key = 'base'
ON CONFLICT (role_id, permission_id) DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key IN (
    'auth.session.refresh',
    'auth.session.logout',
    'profile.self.read',
    'sale.order.create',
    'sale.order.read',
    'payment.create'
)
WHERE r.key = 'cashier'
ON CONFLICT (role_id, permission_id) DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
JOIN permissions p ON p.key IN (
    'auth.session.refresh',
    'auth.session.logout',
    'profile.self.read',
    'sale.order.create',
    'sale.order.read',
    'payment.create',
    'inventory.manage',
    'report.read',
    'account.role.assign'
)
WHERE r.key = 'admin'
ON CONFLICT (role_id, permission_id) DO NOTHING;
