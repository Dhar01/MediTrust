-- name: CreateMedicine :exec
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
);

-- name: GetMedicines :many
SELECT * FROM medicines;
