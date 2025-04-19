package services

import (
	"context"
	"errors"
	"log"
	"medicine-app/internal/errs"
	"medicine-app/internal/repository"
	"medicine-app/models"

	"github.com/google/uuid"
)

type MedService interface {
	CreateMedicine(ctx context.Context, med models.CreateMedicineDTO) (models.Medicine, error)
	UpdateMedicine(ctx context.Context, medID uuid.UUID, med models.UpdateMedicineDTO) (models.Medicine, error)
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
	return srv.medicineRepo.CreateMedicine(ctx, med)
}

func (srv MedicineServiceImpl) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	return srv.medicineRepo.DeleteMedicine(ctx, medID)
}

func (srv MedicineServiceImpl) UpdateMedicine(ctx context.Context, medID uuid.UUID, med models.UpdateMedicineDTO) (models.Medicine, error) {
	oldMedicine, err := srv.medicineRepo.FetchMedicineByID(ctx, medID)
	if err != nil {
		return wrapMedicineErr(err)
	}

	newMedicine, err := srv.medicineRepo.UpdateMedicine(ctx, models.Medicine{
		Id:           medID,
		Name:         updateField(med.Name, oldMedicine.Name),
		Dosage:       updateField(med.Dosage, oldMedicine.Dosage),
		Manufacturer: updateField(med.Manufacturer, oldMedicine.Manufacturer),
		Description:  updateField(med.Description, oldMedicine.Description),
		Price:        *updateIntPointerField(&med.Price, &oldMedicine.Price),
		Stock:        *updateIntPointerField(&med.Price, &oldMedicine.Price),
	})

	log.Println(newMedicine)
	log.Println(err)
	if err != nil {
		return wrapMedicineErr(err)
	}

	return srv.medicineRepo.UpdateMedicine(ctx, newMedicine)
}

func (srv MedicineServiceImpl) FetchMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error) {
	return srv.medicineRepo.FetchMedicineByID(ctx, medID)
}

func updateField(newValue, oldValue string) string {
	if newValue == "" {
		return oldValue
	}

	return newValue
}

func updateIntPointerField(newValue, oldValue *int32) *int32 {
	if newValue == nil {
		return oldValue
	}

	return newValue
}

func wrapMedicineErr(err error) (models.Medicine, error) {
	if errors.Is(err, errs.ErrNotFound) {
		return models.Medicine{}, errs.ErrMedicineNotExist
	}

	return models.Medicine{}, err
}
