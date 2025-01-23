package service

import (
	"context"
	"medicine-app/models"

	"github.com/google/uuid"
)

type adminService struct {
	Repo models.AdminRepository
}

func NewAdminService(repo models.AdminRepository) models.AdminService {
	return &adminService{
		Repo: repo,
	}
}

func (as *adminService) UpdatePermissions(ctx context.Context, id uuid.UUID, permissions []models.Permission) error {
	return nil
}
