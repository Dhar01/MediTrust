// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: medicines.sql

package database

import (
	"context"
)

const createMedicine = `-- name: CreateMedicine :exec
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
`

type CreateMedicineParams struct {
	Name         string
	Dosage       string
	Manufacturer string
	Price        int32
	Stock        int32
}

func (q *Queries) CreateMedicine(ctx context.Context, arg CreateMedicineParams) error {
	_, err := q.db.ExecContext(ctx, createMedicine,
		arg.Name,
		arg.Dosage,
		arg.Manufacturer,
		arg.Price,
		arg.Stock,
	)
	return err
}

const getMedicines = `-- name: GetMedicines :many
SELECT id, name, dosage, manufacturer, price, stock, created_at, updated_at FROM medicines
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
