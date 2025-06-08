-- name: CreateMedicine :one
INSERT INTO medicines (
    id, dosage, expiry_date
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFullMedicineByID :one
SELECT
    p.id, p.name, p.manufacturer, p.description, p.price,
    p.cost, p.stock, p.type, p.created_at, p.updated_at,
    m.dosage, m.expiry_date, p.created_at, p.updated_at
FROM medicines m
JOIN products p on m.id = p.id
WHERE p.id = $1;

-- name: GetMedicineInfoByID :one
SELECT * FROM medicines WHERE id = $1;

-- name: UpdateMedicine :one
UPDATE medicines
SET
    dosage = $1,
    expiry_date = $2
WHERE id = $3
RETURNING *;