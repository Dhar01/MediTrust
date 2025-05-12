package product

import (
	"context"
	"errors"
	"fmt"
	"medicine-app/internal/errs"

	"github.com/google/uuid"
)

type medService struct {
	medicineRepo medicineRepository
}

func newMedicineService(repo medicineRepository) medicineService {
	if repo == nil {
		panic("repository can't be nil/empty")
	}

	return &medService{
		medicineRepo: repo,
	}
}

func (srv *medService) CreateMedicine(ctx context.Context, med medicine) (*medicine, error) {
	return srv.medicineRepo.Create(ctx, med)
}

func (srv *medService) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	return srv.medicineRepo.Delete(ctx, medID)
}

func (srv *medService) UpdateMedicine(ctx context.Context, medID uuid.UUID, med medicine) (*medicine, error) {
	existing, err := srv.medicineRepo.FetchByID(ctx, medID)
	if err != nil {
		return nil, fmt.Errorf("could not load medicine %s: %w", medID, err)
	}

	if med.Price < 0 {
		return nil, errInvalidMedicine
	}

	if med.Stock < -1 {
		return nil, errInvalidMedicine
	}

	updated, err := srv.medicineRepo.Update(ctx, medicine{
		ID:           medID,
		Name:         updateField(med.Name, existing.Name),
		Dosage:       updateField(med.Dosage, existing.Dosage),
		Manufacturer: updateField(med.Manufacturer, existing.Manufacturer),
		Description:  updateField(med.Description, existing.Description),
		Price:        *updateIntPointerField(&med.Price, &existing.Price),
		Stock:        *updateIntPointerField(&med.Price, &existing.Price),
	})
	if err != nil {
		return nil, fmt.Errorf("update operation failed: %w", err)
	}

	return updated, nil
}

func (srv *medService) FetchMedicineByID(ctx context.Context, medID uuid.UUID) (*medicine, error) {
	return srv.medicineRepo.FetchByID(ctx, medID)
}

func (srv *medService) FetchMedicineList(ctx context.Context) ([]medicine, error) {
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
