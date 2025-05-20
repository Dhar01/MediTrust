package product

import (
	"context"
	"errors"

	"medicine-app/internal/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type pgMedicineRepo struct {
	DB *database.Queries
}

func newMedicineRepo(db *database.Queries) medicineRepository {
	if db == nil {
		panic("database can't be nil")
	}

	return &pgMedicineRepo{
		DB: db,
	}
}

func (repo *pgMedicineRepo) Create(ctx context.Context, med medicine) (*medicine, error) {
	created, err := repo.DB.CreateMedicine(ctx, database.CreateMedicineParams{
		Name:         med.Name,
		Dosage:       med.Dosage,
		Description:  med.Description,
		Manufacturer: med.Manufacturer,
		Price:        med.Price,
		Stock:        med.Stock,
	})
	if err != nil {
		return nil, wrapRepoErr(err)
	}

	return toMedicineDomain(created), nil
}

func (repo *pgMedicineRepo) Update(ctx context.Context, med medicine) (*medicine, error) {
	updated, err := repo.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
		ID:           med.ID,
		Name:         med.Name,
		Dosage:       med.Dosage,
		Manufacturer: med.Manufacturer,
		Description:  med.Description,
		Price:        med.Price,
		Stock:        med.Stock,
	})
	if err != nil {
		return nil, wrapRepoErr(err)
	}

	return toMedicineDomain(updated), nil
}

func (repo *pgMedicineRepo) Delete(ctx context.Context, medID uuid.UUID) error {
	return wrapRepoErr(repo.DB.DeleteMedicine(ctx, medID))
}

func (repo *pgMedicineRepo) FetchByID(ctx context.Context, medID uuid.UUID) (*medicine, error) {
	med, err := repo.DB.GetMedicineByID(ctx, medID)
	if err != nil {
		return nil, wrapRepoErr(err)
	}

	return toMedicineDomain(med), nil
}

func (repo *pgMedicineRepo) FetchByName(ctx context.Context, name string) error {
	if err := repo.DB.GetMedicineByName(ctx, name); err != nil {
		return err
	}

	return nil
}

func (repo *pgMedicineRepo) FetchList(ctx context.Context) ([]*medicine, error) {
	_, err := repo.DB.GetMedicines(ctx)
	if err != nil {
		return nil, wrapRepoErr(err)
	}

	return []*medicine{}, nil
}

func toMedicineDomain(dbMed database.Medicine) *medicine {
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

func wrapRepoErr(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return errNotFound
	}

	// handle other errors: constraint violation messages, duplicate key, etc..

	return err
}
