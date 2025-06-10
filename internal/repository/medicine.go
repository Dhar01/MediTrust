package repository

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// MedicineRepo defines the interaction methods with the database
type MedicineRepo interface {
	Create(ctx context.Context, med model.Medicine) (*model.Medicine, error)
	Update(ctx context.Context, med model.Medicine) (*model.Medicine, error)
	FetchByID(ctx context.Context, medID uuid.UUID) (*model.Medicine, error)
}

var _ MedicineRepo = (*medPostgresRepo)(nil)

// medPostgresRepo defines the db connection for medicine on postgres
type medPostgresRepo struct {
	db *database.Queries
}

// NewMedPostgresRepo will return a new Postgres connection for medicine domain
func NewMedPostgresRepo(db *database.Queries) *medPostgresRepo {
	return &medPostgresRepo{
		db: db,
	}
}

// Create creates a new medicine info on the database
func (r *medPostgresRepo) Create(ctx context.Context, med model.Medicine) (*model.Medicine, error) {
	newMedicine, err := r.db.CreateMedicine(ctx, database.CreateMedicineParams{
		ID:     med.ID,
		Dosage: med.Dosage,
		ExpiryDate: pgtype.Timestamp{
			Time: med.ExpiryDate,
		},
	})

	if err != nil {
		return nil, setErrorMsg(err)
	}

	return toMedicine(newMedicine), nil
}

// Update updates a medicine info on the database
func (r *medPostgresRepo) Update(ctx context.Context, med model.Medicine) (*model.Medicine, error) {
	updatedMed, err := r.db.UpdateMedicine(ctx, database.UpdateMedicineParams{
		ID:     med.ID,
		Dosage: med.Dosage,
		ExpiryDate: pgtype.Timestamp{
			Time: med.ExpiryDate,
		},
	})

	if err != nil {
		return nil, setErrorMsg(err)
	}

	return toMedicine(updatedMed), nil
}

// FetchByID returns a medicine info by its ID from the database
func (r *medPostgresRepo) FetchByID(ctx context.Context, medID uuid.UUID) (*model.Medicine, error) {
	medicine, err := r.db.GetMedicineInfoByID(ctx, medID)
	if err != nil {
		return nil, err
	}

	return toMedicine(medicine), nil
}

// helper: toMedicine converts database.Medicine type to model.Medicine type
func toMedicine(medicine database.Medicine) *model.Medicine {
	return &model.Medicine{
		Dosage:     medicine.Dosage,
		ExpiryDate: medicine.ExpiryDate.Time,
	}
}
