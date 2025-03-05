package service

import (
	"context"
	"medicine-app/models"

	"github.com/google/uuid"
)

type adminService struct {
	repo models.AdminRepository
}

func NewAdminService(repo models.AdminRepository) models.AdminService {
	return &adminService{
		repo: repo,
	}
}

func (as *adminService) UpdatePermissions(ctx context.Context, id uuid.UUID, permissions []models.Permission) error {
	return nil
}
