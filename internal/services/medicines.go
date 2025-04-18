package services

import (
	"context"
	"medicine-app/internal/repository"
	"medicine-app/models"

	"github.com/google/uuid"
)

type MedService interface {
	CreateMedicine(ctx context.Context, med models.CreateMedicineDTO) (models.Medicine, error)
	UpdateMedicine(ctx context.Context, med models.UpdateMedicineDTO) (models.Medicine, error)
	DeleteMedicine(ctx context.Context, medID uuid.UUID) error
	FetchMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error)
}

type MedicineServiceImpl struct {
	medicineRepo repository.MedicineRepository
}

func NewMedicineService(repo repository.MedicineRepository) MedService {
	return MedicineServiceImpl{
		medicineRepo: repo,
	}
}

func (srv MedicineServiceImpl) CreateMedicine(ctx context.Context, med models.CreateMedicineDTO) (models.Medicine, error) {
	return models.Medicine{}, nil
}

func (srv MedicineServiceImpl) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	return nil
}

func (srv MedicineServiceImpl) UpdateMedicine(ctx context.Context, med models.UpdateMedicineDTO) (models.Medicine, error) {
	return models.Medicine{}, nil
}

func (srv MedicineServiceImpl) FetchMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error) {
	return models.Medicine{}, nil
}
