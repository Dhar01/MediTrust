-- +goose Up
CREATE TABLE IF NOT EXISTS cart_items(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cart_id UUID NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
    medicine_id UUID NOT NULL REFERENCES medicines(id),
    quantity INT CHECK (quantity > 0),
    price INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    UNIQUE(cart_id, medicine_id)
);

-- +goose Down
DROP TABLE IF EXISTS cart_items;