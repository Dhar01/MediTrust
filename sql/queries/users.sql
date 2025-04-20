-- name: CreateUser :one
INSERT INTO users (
    id, first_name, last_name, email, age, phone, role, password_hash, created_at, updated_at
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

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: GetRole :one
SELECT role FROM users WHERE id = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByPhone :one
SELECT * FROM users WHERE phone = $1;

-- name: GetPass :one
SELECT password_hash, id FROM users WHERE email = $1;

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

-- name: SetVerified :exec
UPDATE users
SET
    verified = TRUE,
    updated_at = NOW()
WHERE id = $1;

-- name: GetVerified :one
SELECT verified FROM users
WHERE id = $1;

-- name: ResetPassword :exec
UPDATE users
SET
    password_hash = $1,
    updated_at = NOW()
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
