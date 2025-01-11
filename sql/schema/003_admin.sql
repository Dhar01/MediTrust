-- +goose Up
CREATE TABLE IF NOT EXISTS admin_roles (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    role TEXT NOT NULL,
    can_manage_users BOOLEAN NOT NULL DEFAULT false,
    can_manage_order BOOLEAN NOT NULL DEFAULT false,
    can_manage_store BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS admin_roles;