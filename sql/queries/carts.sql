-- name: GetCart :many
SELECT
    cart.id AS cart_id,
    cart.created_at,
    ci.id,
    ci.medicine_id,
    m.name AS medicine_name,
    ci.quantity,
    m.price AS price
FROM cart
LEFT JOIN cart_item ci ON cart.id = ci.cart_id
LEFT JOIN medicines m ON ci.medicine_id = m.id
WHERE cart.user_id = $1;

-- name: DeleteCart :exec
DELETE FROM cart
WHERE cart.user_id = $1;

-- name: RemoveCartItem :exec
DELETE FROM cart_item
WHERE cart_item.cart_id = $1
AND medicine_id = $2;

-- name: AddItemToCart :one
INSERT INTO cart_item (
    medicine_id, cart_id, quantity, price
) VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING cart_id;

-- name: UpdateCartItem :exec
UPDATE cart_item
SET
    quantity = $1
WHERE
    medicine_id = $2
    AND cart_id = $3;

-- name: GetCartByUserID :one
SELECT id FROM cart WHERE user_id = $1;

-- name: CreateCart :one
INSERT INTO cart(user_id) VALUES($1) RETURNING id;