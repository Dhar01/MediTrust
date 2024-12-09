-- +goose Up
CREATE TABLE IF NOT EXISTS medicine (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    dosage TEXT NOT NULL,
    manufacturer TEXT NOT NULL,
    price INTEGER NOT NULL,
    stock INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS medicine;