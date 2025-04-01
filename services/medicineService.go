package service

import (
	"context"
	"errors"
	med "medicine-app/internal/api/medicines_gen"
	"medicine-app/internal/database"

	"github.com/google/uuid"
)

type medicineService struct {
	DB *database.Queries
}

// MedicineService defines the business logic interface for medicine management
// @Description Interface for medicine-related business logic
type MedicineService interface {
	CreateMedicine(ctx context.Context, medicine med.CreateMedicineDTO) (med.Medicine, error)
	DeleteMedicine(ctx context.Context, medID med.MedicineID) error
	UpdateMedicine(ctx context.Context, medID med.MedicineID, medicine med.UpdateMedicineDTO) (med.Medicine, error)
	GetMedicines(ctx context.Context) ([]med.Medicine, error)
	GetMedicineByID(ctx context.Context, medID med.MedicineID) (med.Medicine, error)
}

func NewMedicineService(db *database.Queries) MedicineService {
	if db == nil {
		panic("database can't be nil")
	}

	return &medicineService{
		DB: db,
	}
}

func (ms *medicineService) CreateMedicine(ctx context.Context, newMed med.CreateMedicineDTO) (med.Medicine, error) {
	medicine, err := ms.DB.CreateMedicine(ctx, database.CreateMedicineParams{
		Name:         newMed.Name,
		Dosage:       newMed.Dosage,
		Description:  newMed.Description,
		Manufacturer: newMed.Manufacturer,
		Price:        int32(newMed.Price),
		Stock:        int32(newMed.Stock),
	})

	if err != nil {
		return wrapMedicineError(err)
	}

	return toMedicineDomain(medicine), nil
}

func (ms *medicineService) DeleteMedicine(ctx context.Context, medID med.MedicineID) error {
	if err := ms.DB.DeleteMedicine(ctx, medID); err != nil {
		return err
	}

	return nil
}

func (ms *medicineService) GetMedicines(ctx context.Context) ([]med.Medicine, error) {
	var medList []med.Medicine

	medicines, err := ms.DB.GetMedicines(ctx)
	if err != nil {
		return medList, err
	}

	for _, medicine := range medicines {
		medList = append(medList, toMedicineDomain(medicine))
	}

	return medList, nil
}

func (ms *medicineService) GetMedicineByID(ctx context.Context, medID med.MedicineID) (med.Medicine, error) {
	medicine, err := ms.DB.GetMedicine(ctx, medID)
	if err != nil {
		return wrapMedicineError(err)
	}

	return toMedicineDomain(medicine), nil
}

func (ms *medicineService) UpdateMedicine(ctx context.Context, medID med.MedicineID, medicine med.UpdateMedicineDTO) (med.Medicine, error) {
	var emptyMed med.Medicine

	oldMed, err := ms.DB.GetMedicine(ctx, medID)
	if err != nil {
		return emptyMed, errors.New("cannot find medicine")
	}

	newMed, err := ms.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
		ID:           medID,
		Name:         updateField(*medicine.Name, oldMed.Name),
		Description:  updateField(*medicine.Description, oldMed.Description),
		Manufacturer: updateField(*medicine.Manufacturer, oldMed.Manufacturer),
		Dosage:       updateField(*medicine.Dosage, oldMed.Dosage),
		Price:        *updateIntPointerField(medicine.Price, &oldMed.Price),
		Stock:        *updateIntPointerField(medicine.Stock, &oldMed.Stock),
	})

	if err != nil {
		return wrapMedicineError(err)
	}

	return toMedicineDomain(newMed), nil
}

func wrapMedicineError(err error) (med.Medicine, error) {
	return med.Medicine{}, err
}

func toMedicineDomain(dbMed database.Medicine) med.Medicine {
	return med.Medicine{
		Id:           (*uuid.UUID)(&dbMed.ID),
		Name:         &dbMed.Name,
		Dosage:       &dbMed.Dosage,
		Description:  &dbMed.Description,
		Manufacturer: &dbMed.Manufacturer,
		Price:        &dbMed.Price,
		Stock:        &dbMed.Stock,
	}
}
