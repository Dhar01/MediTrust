-- name: CreateUserAddress :one
INSERT INTO user_address (
    user_id, country, city, street_address, postal_code, created_at, updated_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    NOW(),
    NOW()
)
RETURNING *;

-- name: CheckAddressExist :one
SELECT EXISTS (SELECT 1 FROM user_address WHERE user_id = $1);

-- name: GetAddress :one
SELECT * FROM user_address
WHERE user_id = $1;

-- name: GetUserWithAddress :one
SELECT users.*, user_address.*
FROM users
LEFT JOIN user_address ON users.id = user_address.id
WHERE users.id = $1;

-- name: UpdateAddress :one
UPDATE user_address
SET
    country = $1,
    city = $2,
    street_address = $3,
    postal_code = $4,
    updated_at = NOW()
WHERE user_id = $5
RETURNING *;
