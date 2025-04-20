package repository

import (
	"context"
	"medicine-app/internal/database/general/genDB"
)

type HelperRepository interface {
	ResetDB(ctx context.Context) error
}

type helperRepo struct {
	DB *genDB.Queries
}

func NewHelperRepo(db *genDB.Queries) HelperRepository {
	if db == nil {
		panic("db can't be empty/nil")
	}

	return &helperRepo{
		DB: db,
	}
}

func (repo *helperRepo) ResetDB(ctx context.Context) error {
	if err := repo.DB.ResetUsers(ctx); err != nil {
		return err
	}

	if err := repo.DB.ResetMedicines(ctx); err != nil {
		return err
	}

	if err := repo.DB.ResetCart(ctx); err != nil {
		return err
	}

	if err := repo.DB.ResetAddress(ctx); err != nil {
		return err
	}

	return nil
}
