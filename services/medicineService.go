package service

import (
	"context"
	"errors"
	"medicine-app/models"

	"github.com/google/uuid"
)

type medicineService struct {
	Repo models.MedicineRepository
}

func NewMedicineService(repo models.MedicineRepository) models.MedicineService {
	if repo == nil {
		panic("repository can't be nil")
	}

	return &medicineService{
		Repo: repo,
	}
}

func (ms *medicineService) CreateMedicine(ctx context.Context, newMed models.CreateMedicineDTO) (models.Medicine, error) {
	medicine := models.Medicine{
		Name:         newMed.Name,
		Dosage:       newMed.Dosage,
		Description:  newMed.Description,
		Manufacturer: newMed.Manufacturer,
		Price:        newMed.Price,
		Stock:        newMed.Stock,
	}

	return ms.Repo.Create(ctx, medicine)
}

func (ms *medicineService) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	return ms.Repo.Delete(ctx, medID)
}

func (ms *medicineService) GetMedicines(ctx context.Context) ([]models.Medicine, error) {
	return ms.Repo.FindAll(ctx)
}

func (ms *medicineService) GetMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error) {
	return ms.Repo.FindByID(ctx, medID)
}

func (ms *medicineService) UpdateMedicine(ctx context.Context, medID uuid.UUID, med models.UpdateMedicineDTO) (models.Medicine, error) {
	var emptyMed models.Medicine

	oldMed, err := ms.Repo.FindByID(ctx, medID)
	if err != nil {
		return emptyMed, errors.New("cannot find medicine")
	}

	medicine := models.Medicine{
		ID:           medID,
		Name:         updateField(med.Name, oldMed.Name),
		Description:  updateField(med.Description, oldMed.Description),
		Manufacturer: updateField(med.Manufacturer, oldMed.Manufacturer),
		Dosage:       updateField(med.Dosage, oldMed.Dosage),
		Price:        *updateIntPointerField(med.Price, &oldMed.Price),
		Stock:        *updateIntPointerField(med.Stock, &oldMed.Stock),
	}

	return ms.Repo.Update(ctx, medicine)
}
