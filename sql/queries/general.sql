-- name: ResetUsers :exec
DELETE FROM users;

-- name: ResetCart :exec
DELETE FROM cart;

-- name: ResetMedicines :exec
DELETE FROM medicines;

-- name: ResetAddress :exec
DELETE FROM user_address;