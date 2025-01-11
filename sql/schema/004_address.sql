-- +goose Up
CREATE TABLE IF NOT EXISTS user_address(
    user_id UUID PRIMARY KEY REFERENCES users(id),
    country TEXT NOT NULL,
    city TEXT NOT NULL,
    street_address TEXT NOT NULL,
    postal_code INTEGER,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS user_address;