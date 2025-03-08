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
WHERE cart_item.cart_id = $1;