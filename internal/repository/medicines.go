package repository

import (
	"context"
	"medicine-app/internal/database/medicine/medDB"
	"medicine-app/models"

	"github.com/google/uuid"
)

type MedicineRepository interface {
	CreateMedicine(ctx context.Context, med models.CreateMedicineDTO) (models.Medicine, error)
	DeleteMedicine(ctx context.Context, medID uuid.UUID) error
	UpdateMedicine(ctx context.Context, med models.Medicine) (models.Medicine, error)
	FetchMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error)
}

type MedRepoImpl struct {
	DB medDB.Queries
}

func NewMedicineRepo(db medDB.Queries) MedicineRepository {
	return MedRepoImpl{
		DB: db,
	}
}

func (repo MedRepoImpl) CreateMedicine(ctx context.Context, med models.CreateMedicineDTO) (models.Medicine, error) {
	medicine, err := repo.DB.CreateMedicine(ctx, medDB.CreateMedicineParams{
		Name:         med.Name,
		Dosage:       med.Dosage,
		Description:  med.Description,
		Manufacturer: med.Manufacturer,
		Price:        med.Price,
		Stock:        med.Stock,
	})

	if err != nil {
		return wrapMedicineErr(err)
	}

	return toMedicineDomain(medicine), nil
}

func (repo MedRepoImpl) UpdateMedicine(ctx context.Context, med models.Medicine) (models.Medicine, error) {
	updatedMedicine, err := repo.DB.UpdateMedicine(ctx, medDB.UpdateMedicineParams{
		ID:           med.Id,
		Name:         med.Name,
		Dosage:       med.Dosage,
		Manufacturer: med.Manufacturer,
		Description:  med.Description,
		Price:        med.Price,
		Stock:        med.Stock,
	})

	if err != nil {
		return wrapMedicineErr(err)
	}

	return toMedicineDomain(updatedMedicine), nil
}

func (repo MedRepoImpl) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	return repo.DB.DeleteMedicine(ctx, medID)
}

func (repo MedRepoImpl) FetchMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error) {
	medicine, err := repo.DB.GetMedicine(ctx, medID)
	if err != nil {
		return wrapMedicineErr(err)
	}

	return toMedicineDomain(medicine), nil
}

func toMedicineDomain(dbMed medDB.Medicine) models.Medicine {
	return models.Medicine{
		Id:           dbMed.ID,
		Name:         dbMed.Name,
		Dosage:       dbMed.Dosage,
		Description:  dbMed.Description,
		Manufacturer: dbMed.Manufacturer,
		Price:        dbMed.Price,
		Stock:        dbMed.Stock,
	}
}

func wrapMedicineErr(err error) (models.Medicine, error) {
	return models.Medicine{}, err
}
