-- +goose Up
CREATE TYPE product_type AS ENUM (
    'medicine',
    'medical_instrument'
);

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    manufacturer TEXT NOT NULL,
    description TEXT NOT NULL,
    price INTEGER NOT NULL,
    cost INTEGER NOT NULL,
    stock INTEGER NOT NULL,
    type product_type NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS products;