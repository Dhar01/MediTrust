package product

import (
	"context"
	"errors"
	"medicine-app/internal/errs"

	"github.com/google/uuid"
)

type medService interface {
	CreateMedicine(ctx context.Context, med medicine) (*medicine, error)
	UpdateMedicine(ctx context.Context, medID uuid.UUID, med medicine) (*medicine, error)
	DeleteMedicine(ctx context.Context, medID uuid.UUID) error
	FetchMedicineByID(ctx context.Context, medID uuid.UUID) (*medicine, error)
	FetchMedicineList(ctx context.Context) ([]medicine, error)
}

type medicineService struct {
	medicineRepo MedicineRepository
}

func NewMedicineService(repo MedicineRepository) medService {
	if repo == nil {
		panic("repository can't be nil/empty")
	}

	return &medicineService{
		medicineRepo: repo,
	}
}

func (srv *medicineService) CreateMedicine(ctx context.Context, med medicine) (*medicine, error) {
	return srv.medicineRepo.CreateMedicine(ctx, med)
}

func (srv *medicineService) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	return srv.medicineRepo.DeleteMedicine(ctx, medID)
}

func (srv *medicineService) UpdateMedicine(ctx context.Context, medID uuid.UUID, med medicine) (*medicine, error) {
	oldMedicine, err := srv.medicineRepo.FetchMedicineByID(ctx, medID)
	if err != nil {
		return wrapServiceErr(err)
	}

	newMedicine, err := srv.medicineRepo.UpdateMedicine(ctx, medicine{
		ID:           medID,
		Name:         updateField(med.Name, oldMedicine.Name),
		Dosage:       updateField(med.Dosage, oldMedicine.Dosage),
		Manufacturer: updateField(med.Manufacturer, oldMedicine.Manufacturer),
		Description:  updateField(med.Description, oldMedicine.Description),
		Price:        *updateIntPointerField(&med.Price, &oldMedicine.Price),
		Stock:        *updateIntPointerField(&med.Price, &oldMedicine.Price),
	})
	if err != nil {
		return wrapServiceErr(err)
	}

	return newMedicine, nil
}

func (srv *medicineService) FetchMedicineByID(ctx context.Context, medID uuid.UUID) (*medicine, error) {
	return srv.medicineRepo.FetchMedicineByID(ctx, medID)
}

func (srv *medicineService) FetchMedicineList(ctx context.Context) ([]medicine, error) {
	return []medicine{}, nil
}

func wrapServiceErr(err error) (*medicine, error) {
	if errors.Is(err, errs.ErrNotFound) {
		return nil, errs.ErrMedicineNotExist
	}

	return nil, err
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
