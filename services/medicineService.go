package service

import (
	"context"
	"errors"
	"medicine-app/internal/database"
	"medicine-app/models/db"
	"medicine-app/models/dto"

	"github.com/google/uuid"
)

type medicineService struct {
	DB *database.Queries
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

func NewMedicineService(db *database.Queries) MedicineService {
	if db == nil {
		panic("database can't be nil")
	}

	return &medicineService{
		DB: db,
	}
}

func (ms *medicineService) CreateMedicine(ctx context.Context, newMed dto.CreateMedicineDTO) (db.Medicine, error) {
	medicine, err := ms.DB.CreateMedicine(ctx, database.CreateMedicineParams{
		Name:         newMed.Name,
		Dosage:       newMed.Dosage,
		Description:  newMed.Description,
		Manufacturer: newMed.Manufacturer,
		Price:        newMed.Price,
		Stock:        newMed.Stock,
	})

	if err != nil {
		return wrapMedicineError(err)
	}

	return toMedicineDomain(medicine), nil
}

func (ms *medicineService) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	if err := ms.DB.DeleteMedicine(ctx, medID); err != nil {
		return err
	}

	return nil
}

func (ms *medicineService) GetMedicines(ctx context.Context) ([]db.Medicine, error) {
	var medList []db.Medicine

	medicines, err := ms.DB.GetMedicines(ctx)
	if err != nil {
		return medList, err
	}

	for _, medicine := range medicines {
		medList = append(medList, toMedicineDomain(medicine))
	}

	return medList, nil
}

func (ms *medicineService) GetMedicineByID(ctx context.Context, medID uuid.UUID) (db.Medicine, error) {
	medicine, err := ms.DB.GetMedicine(ctx, medID)
	if err != nil {
		return wrapMedicineError(err)
	}

	return toMedicineDomain(medicine), nil
}

func (ms *medicineService) UpdateMedicine(ctx context.Context, medID uuid.UUID, med dto.UpdateMedicineDTO) (db.Medicine, error) {
	var emptyMed db.Medicine

	oldMed, err := ms.DB.GetMedicine(ctx, medID)
	if err != nil {
		return emptyMed, errors.New("cannot find medicine")
	}

	medicine, err := ms.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
		ID:           medID,
		Name:         updateField(med.Name, oldMed.Name),
		Description:  updateField(med.Description, oldMed.Description),
		Manufacturer: updateField(med.Manufacturer, oldMed.Manufacturer),
		Dosage:       updateField(med.Dosage, oldMed.Dosage),
		Price:        *updateIntPointerField(med.Price, &oldMed.Price),
		Stock:        *updateIntPointerField(med.Stock, &oldMed.Stock),
	})

	if err != nil {
		return wrapMedicineError(err)
	}

	return toMedicineDomain(medicine), nil
}

func wrapMedicineError(err error) (db.Medicine, error) {
	return db.Medicine{}, err
}

func toMedicineDomain(dbMed database.Medicine) db.Medicine {
	return db.Medicine{
		ID:           dbMed.ID,
		Name:         dbMed.Name,
		Dosage:       dbMed.Dosage,
		Description:  dbMed.Description,
		Manufacturer: dbMed.Manufacturer,
		Price:        dbMed.Price,
		Stock:        dbMed.Stock,
		Updated_at:   dbMed.UpdatedAt,
	}
}
