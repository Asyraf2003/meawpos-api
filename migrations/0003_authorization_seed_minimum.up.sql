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

INSERT INTO roles (key, name)
VALUES
    ('base', 'Base'),
    ('cashier', 'Cashier'),
    ('admin', 'Admin')
ON CONFLICT (key) DO NOTHING;

INSERT INTO permissions (key, name)
VALUES
    ('auth.session.refresh', 'Refresh auth session'),
    ('auth.session.logout', 'Logout auth session'),
    ('profile.self.read', 'Read own profile'),
    ('sale.order.create', 'Create sale order'),
    ('sale.order.read', 'Read sale order'),
    ('payment.create', 'Create payment'),
    ('inventory.manage', 'Manage inventory'),
    ('report.read', 'Read reports'),
    ('account.role.assign', 'Assign account roles')
ON CONFLICT (key) DO NOTHING;
