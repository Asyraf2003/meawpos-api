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

DELETE FROM role_permissions
WHERE permission_id IN (
    SELECT id
    FROM permissions
    WHERE key = 'capability.manage'
)
AND role_id IN (
    SELECT id
    FROM roles
    WHERE key = 'admin'
);

DELETE FROM api_capabilities
WHERE key = 'capability.manage';

DELETE FROM permissions
WHERE key = 'capability.manage';

ALTER TABLE api_capabilities
    DROP CONSTRAINT api_capabilities_method_check;

ALTER TABLE api_capabilities
    ADD CONSTRAINT api_capabilities_method_check
        CHECK (method IN ('GET', 'POST', 'PUT', 'PATCH', 'DELETE'));
