package repository

import (
	"context"
	"errors"
	"medicine-app/internal/database/medicine/medDB"
	"medicine-app/internal/errs"
	"medicine-app/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type MedicineRepository interface {
	CreateMedicine(ctx context.Context, med models.CreateMedicineDTO) (models.Medicine, error)
	DeleteMedicine(ctx context.Context, medID uuid.UUID) error
	UpdateMedicine(ctx context.Context, med models.Medicine) (models.Medicine, error)
	FetchMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error)
}

type medRepo struct {
	DB *medDB.Queries
}

func NewMedicineRepo(db *medDB.Queries) MedicineRepository {
	return &medRepo{
		DB: db,
	}
}

func (repo *medRepo) CreateMedicine(ctx context.Context, med models.CreateMedicineDTO) (models.Medicine, error) {
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

func (repo *medRepo) UpdateMedicine(ctx context.Context, med models.Medicine) (models.Medicine, error) {
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

func (repo *medRepo) DeleteMedicine(ctx context.Context, medID uuid.UUID) error {
	return wrapMedSpecErr(repo.DB.DeleteMedicine(ctx, medID))
}

func (repo *medRepo) FetchMedicineByID(ctx context.Context, medID uuid.UUID) (models.Medicine, error) {
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
	return models.Medicine{}, wrapMedSpecErr(err)
}

func wrapMedSpecErr(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return errs.ErrNotFound
	}

	return err
}
