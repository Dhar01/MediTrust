-- +goose Up
CREATE TABLE IF NOT EXISTS cart_item(
    id SERIAL PRIMARY KEY,
    cart_id UUID NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
    medicine_id UUID NOT NULL REFERENCES medicines(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    price INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    UNIQUE(cart_id, medicine_id)
);

-- +goose Down
DROP TABLE IF EXISTS cart_item;