package service

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models"

	"github.com/google/uuid"
)

type adminService struct {
	DB *database.Queries
}

func NewAdminService(db *database.Queries) models.AdminService {
	if db == nil {
		panic("database can't be nil")
	}

	return &adminService{
		DB: db,
	}
}

func (as *adminService) UpdatePermissions(ctx context.Context, id uuid.UUID, permissions []models.Permission) error {
	return nil
}
