-- name: CreateAdmin :one
INSERT INTO admin_roles (
    user_id, role, created_at, updated_at
) VALUES (
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetAdminByUserID :one
SELECT * FROM admin_roles
WHERE user_id = $1;

-- name: GetAdminWithPermission :many
SELECT
    admin_roles.*, permissions.name
FROM admin_roles
JOIN admin_permissions ON admin_roles.user_id = admin_permissions.admin_id
JOIN permissions ON permissions.permission_id = admin_permissions.permission_id
WHERE admin_roles.user_id = $1;