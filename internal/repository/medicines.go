package repository

import (
	"context"
	"medicine-app/internal/database/medicine/medDB"
	"medicine-app/models"

	"github.com/google/uuid"
)

type MedicineRepository interface {
	CreateMedicine(ctx context.Context, med models.CreateMedicineDTO) (models.Medicine, error)
	DeleteMedicine(ctx context.Context, medID uuid.UUID) error
	UpdateMedicine(ctx context.Context, med models.UpdateMedicineDTO) (models.Medicine, error)
	FetchMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error)
}

type MedRepoImpl struct {
	DB medDB.Queries
}

func NewMedicineRepo(db medDB.Queries) MedicineRepository {
	return MedRepoImpl{
		DB: db,
	}
}

func (m MedRepoImpl) CreateMedicine(ctx context.Context, med models.CreateMedicineDTO) (models.Medicine, error) {
	return models.Medicine{}, nil
}

func (repo MedRepoImpl) UpdateMedicine(ctx context.Context, med models.UpdateMedicineDTO) (models.Medicine, error) {
	return models.Medicine{}, nil
}

func (repo MedRepoImpl) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	return nil
}

func (repo MedRepoImpl) FetchMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error) {
	return models.Medicine{}, nil
}
