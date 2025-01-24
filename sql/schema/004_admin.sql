-- +goose Up
CREATE TABLE IF NOT EXISTS admins(
    admin_id UUID PRIMARY KEY REFERENCES users(id),
    is_super_admin BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS admins;
