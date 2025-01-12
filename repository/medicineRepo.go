package repository

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models"

	"github.com/google/uuid"
)

type medicineRepository struct {
	DB *database.Queries
}

func NewMedicineRepository(db *database.Queries) models.MedicineRepository {
	return &medicineRepository{
		DB: db,
	}
}

func (mr *medicineRepository) Create(ctx context.Context, newMed models.Medicine) (models.Medicine, error) {
	medicine, err := mr.DB.CreateMedicine(ctx, database.CreateMedicineParams{
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

func (mr *medicineRepository) Delete(ctx context.Context, medID uuid.UUID) error {
	if err := mr.DB.DeleteMedicine(ctx, medID); err != nil {
		return err
	}

	return nil
}

func (mr *medicineRepository) Update(ctx context.Context, med models.Medicine) (models.Medicine, error) {
	medicine, err := mr.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
		ID:           med.ID,
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

func (mr *medicineRepository) FindByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error) {
	medicine, err := mr.DB.GetMedicine(ctx, medID)
	if err != nil {
		return models.Medicine{}, err
	}

	return toMedicineDomain(medicine), nil
}

func (mr *medicineRepository) FindAll(ctx context.Context) ([]models.Medicine, error) {
	var medList []models.Medicine

	medicines, err := mr.DB.GetMedicines(ctx)
	if err != nil {
		return medList, err
	}

	for _, medicine := range medicines {
		medList = append(medList, toMedicineDomain(medicine))
	}

	return medList, nil
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
