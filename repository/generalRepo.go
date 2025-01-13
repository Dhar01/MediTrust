package repository

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models"
)

type generalRepository struct {
	DB *database.Queries
}

func NewGeneralRepository(db *database.Queries) models.GeneralRepository {
	return &generalRepository{
		DB: db,
	}
}

func (gr *generalRepository) ResetMedicineRepo(ctx context.Context) error {
	return gr.DB.ResetMedicines(ctx)
}

func (gr *generalRepository) ResetUserRepo(ctx context.Context) error {
	return gr.DB.ResetUsers(ctx)
}

func (gr *generalRepository) ResetAddressRepo(ctx context.Context) error {
	return gr.DB.ResetAddress(ctx)
}
