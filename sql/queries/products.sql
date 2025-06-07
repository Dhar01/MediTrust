-- name: CreateProduct :one
INSERT INTO products (
    id, name, manufacturer, description, price, cost, stock, type, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProductsByType :many
SELECT * FROM products WHERE type = $1;

-- name: GetProductByID :one
SELECT * FROM products WHERE id = $1;

-- name: GetProductByName :one
SELECT * FROM products WHERE name = $1;


-- name: UpdateProduct :one
UPDATE products
SET
    name = $1,
    manufacturer = $2,
    description = $3,
    price = $4,
    cost = $5,
    stock = $6,
    type = $7,
    updated_at = NOW()
WHERE id = $8
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;