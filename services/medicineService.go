package service

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models"

	"github.com/google/uuid"
)

type medicineService struct {
	DB *database.Queries
}

func NewMedicineService(db *database.Queries) models.MedicineService {
	if db == nil {
		panic("database can't be nil")
	}

	return &medicineService{
		DB: db,
	}
}

func (ms *medicineService) CreateMedicine(ctx context.Context, newMed models.Medicine) (models.Medicine, error) {
	medicine, err := ms.DB.CreateMedicine(ctx, database.CreateMedicineParams{
		Name:         newMed.Name,
		Dosage:       newMed.Dosage,
		Description:  newMed.Description,
		Manufacturer: newMed.Manufacturer,
		Price:        newMed.Price,
		Stock:        newMed.Stock,
	})

	if err != nil {
		return models.Medicine{}, err
	}

	return toMedicineDomain(medicine), nil
}

func (ms *medicineService) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	if err := ms.DB.DeleteMedicine(ctx, medID); err != nil {
		return err
	}

	return nil
}

func (ms *medicineService) GetMedicines(ctx context.Context) ([]models.Medicine, error) {
	var medList []models.Medicine

	medicines, err := ms.DB.GetMedicines(ctx)
	if err != nil {
		return medList, err
	}

	for _, medicine := range medicines {
		medList = append(medList, toMedicineDomain(medicine))
	}

	return medList, nil
}

func (ms *medicineService) GetMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error) {
	medicine, err := ms.DB.GetMedicine(ctx, medID)
	if err != nil {
		return models.Medicine{}, err
	}

	return toMedicineDomain(medicine), nil
}

func (ms *medicineService) UpdateMedicine(ctx context.Context, medID uuid.UUID, med models.Medicine) (models.Medicine, error) {
	medicine, err := ms.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
		ID:           medID,
		Name:         med.Name,
		Description:  med.Description,
		Manufacturer: med.Manufacturer,
		Dosage:       med.Dosage,
		Price:        med.Price,
		Stock:        med.Stock,
	})

	if err != nil {
		return models.Medicine{}, err
	}

	return toMedicineDomain(medicine), nil

}

func toMedicineDomain(dbMed database.Medicine) models.Medicine {
	return models.Medicine{
		ID:           dbMed.ID,
		Name:         dbMed.Name,
		Description:  dbMed.Description,
		Manufacturer: dbMed.Manufacturer,
		Price:        dbMed.Price,
		Stock:        dbMed.Stock,
		Updated_at:   dbMed.UpdatedAt,
	}
}
