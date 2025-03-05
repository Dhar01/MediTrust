package service

import (
	"context"
	"errors"
	"medicine-app/models/db"
	"medicine-app/models/dto"
	"medicine-app/repository"

	"github.com/google/uuid"
)

type medicineService struct {
	repo repository.MedicineRepository
}

// MedicineService defines the business logic interface for medicine management
// @Description Interface for medicine-related business logic
type MedicineService interface {
	CreateMedicine(ctx context.Context, medicine dto.CreateMedicineDTO) (db.Medicine, error)
	DeleteMedicine(ctx context.Context, medID uuid.UUID) error
	UpdateMedicine(ctx context.Context, medID uuid.UUID, med dto.UpdateMedicineDTO) (db.Medicine, error)
	GetMedicines(ctx context.Context) ([]db.Medicine, error)
	GetMedicineByID(ctx context.Context, medID uuid.UUID) (db.Medicine, error)
}

func NewMedicineService(repo repository.MedicineRepository) MedicineService {
	if repo == nil {
		panic("repository can't be nil")
	}

	return &medicineService{
		repo: repo,
	}
}

func (ms *medicineService) CreateMedicine(ctx context.Context, newMed dto.CreateMedicineDTO) (db.Medicine, error) {
	medicine := db.Medicine{
		Name:         newMed.Name,
		Dosage:       newMed.Dosage,
		Description:  newMed.Description,
		Manufacturer: newMed.Manufacturer,
		Price:        newMed.Price,
		Stock:        newMed.Stock,
	}

	return ms.repo.Create(ctx, medicine)
}

func (ms *medicineService) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	return ms.repo.Delete(ctx, medID)
}

func (ms *medicineService) GetMedicines(ctx context.Context) ([]db.Medicine, error) {
	return ms.repo.FindAll(ctx)
}

func (ms *medicineService) GetMedicineByID(ctx context.Context, medID uuid.UUID) (db.Medicine, error) {
	return ms.repo.FindByID(ctx, medID)
}

func (ms *medicineService) UpdateMedicine(ctx context.Context, medID uuid.UUID, med dto.UpdateMedicineDTO) (db.Medicine, error) {
	var emptyMed db.Medicine

	oldMed, err := ms.repo.FindByID(ctx, medID)
	if err != nil {
		return emptyMed, errors.New("cannot find medicine")
	}

	medicine := db.Medicine{
		ID:           medID,
		Name:         updateField(med.Name, oldMed.Name),
		Description:  updateField(med.Description, oldMed.Description),
		Manufacturer: updateField(med.Manufacturer, oldMed.Manufacturer),
		Dosage:       updateField(med.Dosage, oldMed.Dosage),
		Price:        *updateIntPointerField(med.Price, &oldMed.Price),
		Stock:        *updateIntPointerField(med.Stock, &oldMed.Stock),
	}

	return ms.repo.Update(ctx, medicine)
}
