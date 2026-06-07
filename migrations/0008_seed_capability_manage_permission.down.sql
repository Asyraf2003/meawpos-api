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
