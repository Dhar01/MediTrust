-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    age INTEGER NOT NULL,
    phone TEXT NOT NULL,
    isVerified BOOLEAN NOT NULL DEFAULT FALSE,
    verification_code TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE UNIQUE INDEX idx_users_email ON users(email);

CREATE UNIQUE INDEX idx_users_phone ON users(phone);

-- +goose Down
DROP TABLE IF EXISTS users;