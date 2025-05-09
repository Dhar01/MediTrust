package product

import (
	"context"
	"errors"
	"medicine-app/internal/database/medicine/medDB"
	"medicine-app/internal/errs"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type MedicineRepository interface {
	CreateMedicine(ctx context.Context, med medicine) (*medicine, error)
	DeleteMedicine(ctx context.Context, medID uuid.UUID) error
	UpdateMedicine(ctx context.Context, med medicine) (*medicine, error)
	FetchMedicineByID(ctx context.Context, medID uuid.UUID) (*medicine, error)
}

type medRepo struct {
	DB *medDB.Queries
}

// type StrictEchoHandlerFunc func(ctx echo.Context, request interface{}) (response interface{}, err error)

func NewMedicineRepo(db *medDB.Queries) MedicineRepository {
	if db == nil {
		panic("database can't be nil")
	}

	return &medRepo{
		DB: db,
	}
}

func (repo *medRepo) CreateMedicine(ctx context.Context, med medicine) (*medicine, error) {

	// med, err := repo.DB.CreateMedicine(ctx, CreateMedicineParams{
	// 	Name:         med.Name,
	// 	Dosage:       med.Dosage,
	// 	Description:  med.Description,
	// 	Manufacturer: med.Manufacturer,
	// 	Price:        med.Price,
	// 	Stock:        med.Stock,
	// })
	// if err != nil {
	// 	return wrapMedicineErr(err)
	// }
	// return toMedicineDomain(med), nil

	return &medicine{}, nil
}

func (repo *medRepo) UpdateMedicine(ctx context.Context, med medicine) (*medicine, error) {
	updatedMedicine, err := repo.DB.UpdateMedicine(ctx, medDB.UpdateMedicineParams{
		ID:           med.ID,
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

func (repo *medRepo) FetchMedicineByID(ctx context.Context, medID uuid.UUID) (*medicine, error) {
	med, err := repo.DB.GetMedicine(ctx, medID)
	if err != nil {
		return wrapMedicineErr(err)
	}

	return toMedicineDomain(med), nil
}

func toMedicineDomain(dbMed medDB.Medicine) *medicine {
	return &medicine{
		ID:           dbMed.ID,
		Name:         dbMed.Name,
		Dosage:       dbMed.Dosage,
		Description:  dbMed.Description,
		Manufacturer: dbMed.Manufacturer,
		Price:        dbMed.Price,
		Stock:        dbMed.Stock,
	}
}

func wrapMedicineErr(err error) (*medicine, error) {
	return nil, wrapMedSpecErr(err)
}

func wrapMedSpecErr(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return errs.ErrNotFound
	}

	return err
}
