package repository

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models"
)

type adminRepository struct {
	DB *database.Queries
}

func NewAdminRepository(db *database.Queries) models.AdminRepository {
	return &adminRepository{
		DB: db,
	}
}

func (ar *adminRepository) PermissionUpdate(ctx context.Context) error {
	return nil
}
