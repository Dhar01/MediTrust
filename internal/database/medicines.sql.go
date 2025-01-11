// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: medicines.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createMedicine = `-- name: CreateMedicine :one
INSERT INTO medicines (
    id, name, description, dosage, manufacturer, price, stock, created_at, updated_at
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
RETURNING id, name, dosage, description, manufacturer, price, stock, created_at, updated_at
`

type CreateMedicineParams struct {
	Name         string
	Description  string
	Dosage       string
	Manufacturer string
	Price        int32
	Stock        int32
}

func (q *Queries) CreateMedicine(ctx context.Context, arg CreateMedicineParams) (Medicine, error) {
	row := q.db.QueryRowContext(ctx, createMedicine,
		arg.Name,
		arg.Description,
		arg.Dosage,
		arg.Manufacturer,
		arg.Price,
		arg.Stock,
	)
	var i Medicine
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Dosage,
		&i.Description,
		&i.Manufacturer,
		&i.Price,
		&i.Stock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteMedicine = `-- name: DeleteMedicine :exec
DELETE FROM medicines
WHERE id = $1
`

func (q *Queries) DeleteMedicine(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteMedicine, id)
	return err
}

const getMedicine = `-- name: GetMedicine :one
SELECT id, name, dosage, description, manufacturer, price, stock, created_at, updated_at FROM medicines
WHERE id = $1
`

func (q *Queries) GetMedicine(ctx context.Context, id uuid.UUID) (Medicine, error) {
	row := q.db.QueryRowContext(ctx, getMedicine, id)
	var i Medicine
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Dosage,
		&i.Description,
		&i.Manufacturer,
		&i.Price,
		&i.Stock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getMedicines = `-- name: GetMedicines :many
SELECT id, name, dosage, description, manufacturer, price, stock, created_at, updated_at FROM medicines
`

func (q *Queries) GetMedicines(ctx context.Context) ([]Medicine, error) {
	rows, err := q.db.QueryContext(ctx, getMedicines)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Medicine
	for rows.Next() {
		var i Medicine
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Dosage,
			&i.Description,
			&i.Manufacturer,
			&i.Price,
			&i.Stock,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const reset = `-- name: Reset :exec
DELETE FROM medicines
`

func (q *Queries) Reset(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, reset)
	return err
}

const updateMedicine = `-- name: UpdateMedicine :one
UPDATE medicines
SET
    name = $1,
    dosage = $2,
    description = $3,
    manufacturer = $4,
    price = $5,
    stock = $6,
    updated_at = NOW()
WHERE id = $7
RETURNING id, name, dosage, description, manufacturer, price, stock, created_at, updated_at
`

type UpdateMedicineParams struct {
	Name         string
	Dosage       string
	Description  string
	Manufacturer string
	Price        int32
	Stock        int32
	ID           uuid.UUID
}

func (q *Queries) UpdateMedicine(ctx context.Context, arg UpdateMedicineParams) (Medicine, error) {
	row := q.db.QueryRowContext(ctx, updateMedicine,
		arg.Name,
		arg.Dosage,
		arg.Description,
		arg.Manufacturer,
		arg.Price,
		arg.Stock,
		arg.ID,
	)
	var i Medicine
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Dosage,
		&i.Description,
		&i.Manufacturer,
		&i.Price,
		&i.Stock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
