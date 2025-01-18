-- name: CreateUser :one
INSERT INTO users (
    id, first_name, last_name, email, age, phone, password_hash, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByPhone :one
SELECT * FROM users WHERE phone = $1;

-- name: UpdateUser :one
UPDATE users
SET
    first_name = $1,
    last_name = $2,
    email = $3,
    age = $4,
    phone = $5,
    updated_at = NOW()
WHERE id = $6
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ResetUsers :exec
DELETE FROM users;