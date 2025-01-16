package models

import (
	"context"

	"github.com/google/uuid"
)

type Permission string

const (
	PermissionManageUsers  Permission = "manage_users"
	PermissionManageOrders Permission = "manage_orders"
	PermissionManageStore  Permission = "manage_store"
)

type Admin struct {
	User
	Role        string
	Permissions []Permission
}

type CreateAdminDTO struct {
	CreateUserDTO
	Role        string       `json:"role" binding:"required"`
	Permissions []Permission `json:"permissions" binding:"required"`
}

type UpdateAdminDTO struct {
	UpdateUserDTO
	Role        *string      `json:"role,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}

type AdminService interface {
	UserService
	UpdatePermissions(ctx context.Context, id uuid.UUID, permissions []Permission) error
}

type AdminRepository interface {
	UserRepository
	PermissionUpdate(ctx context.Context) error
}