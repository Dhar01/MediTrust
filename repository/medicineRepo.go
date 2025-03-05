package repository

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models/db"

	"github.com/google/uuid"
)

type medicineRepository struct {
	DB *database.Queries
}

// MedicineRepository defines the DB operations for medicines
// @Description Interface for medicine database transactions
type MedicineRepository interface {
	Create(ctx context.Context, med db.Medicine) (db.Medicine, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, med db.Medicine) (db.Medicine, error)
	FindByID(ctx context.Context, id uuid.UUID) (db.Medicine, error)
	FindAll(ctx context.Context) ([]db.Medicine, error)
}

func NewMedicineRepository(db *database.Queries) MedicineRepository {
	return &medicineRepository{
		DB: db,
	}
}

func (mr *medicineRepository) Create(ctx context.Context, newMed db.Medicine) (db.Medicine, error) {
	medicine, err := mr.DB.CreateMedicine(ctx, database.CreateMedicineParams{
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

func (mr *medicineRepository) Delete(ctx context.Context, medID uuid.UUID) error {
	if err := mr.DB.DeleteMedicine(ctx, medID); err != nil {
		return err
	}

	return nil
}

func (mr *medicineRepository) Update(ctx context.Context, med db.Medicine) (db.Medicine, error) {
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
		return wrapMedicineError(err)
	}

	return toMedicineDomain(medicine), nil
}

func (mr *medicineRepository) FindByID(ctx context.Context, medID uuid.UUID) (db.Medicine, error) {
	medicine, err := mr.DB.GetMedicine(ctx, medID)
	if err != nil {
		return wrapMedicineError(err)
	}

	return toMedicineDomain(medicine), nil
}

func (mr *medicineRepository) FindAll(ctx context.Context) ([]db.Medicine, error) {
	var medList []db.Medicine

	medicines, err := mr.DB.GetMedicines(ctx)
	if err != nil {
		return medList, err
	}

	for _, medicine := range medicines {
		medList = append(medList, toMedicineDomain(medicine))
	}

	return medList, nil
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
