-- +goose Up
CREATE TABLE IF NOT EXISTS admin_roles (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    role TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS permissions (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS admin_permissions (
    admin_id UUID REFERENCES admin_roles(user_id) ON DELETE CASCADE,
    permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY(admin_id, permission_id)
);

-- +goose Down
DROP TABLE IF EXISTS admin_roles;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS admin_permissions;