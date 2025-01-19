-- +goose Up
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS permissions;
