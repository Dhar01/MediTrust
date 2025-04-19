-- name: CreateMedicine :one
INSERT INTO medicines (
    id, name, dosage, description, manufacturer, price, stock, created_at, updated_at
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

-- name: GetMedicines :many
SELECT * FROM medicines;

-- name: GetMedicine :one
SELECT * FROM medicines
WHERE id = $1;

-- name: UpdateMedicine :one
UPDATE medicines
SET
    name = $1,
    dosage = $2,
    description = $3,
    manufacturer = $4,
    price = $5,
    stock = $6,
    updated_at = NOW() + interval '1 second'
WHERE id = $7
RETURNING *;

-- name: DeleteMedicine :exec
DELETE FROM medicines
WHERE id = $1;

-- name: ResetMedicines :exec
DELETE FROM medicines;