-- +goose Up
CREATE TABLE IF NOT EXISTS medicines (
    id UUID PRIMARY KEY REFERENCES products(id) ON DELETE CASCADE,
    dosage TEXT NOT NULL,
    expiry_date TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS medicines;