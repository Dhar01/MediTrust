-- name: CreateMedicine :one
INSERT INTO medicines (id, name, dosage, manufacturer, price, stock, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
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
    manufacturer = $3,
    price = $4,
    stock = $5,
    updated_at = NOW()
WHERE id = $6
RETURNING *;

-- name: DeleteMedicine :exec
DELETE FROM medicines
WHERE id = $1;