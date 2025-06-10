package model

import (
	"time"

	"github.com/google/uuid"
)

// RoleType defines the type for user roles in User entity
type RoleType string

// Constants for user roles
const (
	RoleCustomer RoleType = "customer"
	RoleDoctor   RoleType = "doctor"
	RoleAdmin    RoleType = "admin"
)

// Auth defines the authentication entity for the user
type Auth struct {
	ID             uuid.UUID
	Email          string
	Phone          string
	HashedPassword string
	Role           RoleType
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// IsUserActive checks if the user is active on the database
func (a *Auth) IsUserActive() bool {
	return a.IsActive
}

// IsUserCustomer returns true if the user is customer
func (a *Auth) IsUserCustomer() bool {
	return a.Role == RoleCustomer
}

// IsUserDoctor returns true if the is doctor
func (a *Auth) IsUserDoctor() bool {
	return a.Role == RoleDoctor
}

// IsUserAdmin returns true if the user is admin
func (a *Auth) IsUserAdmin() bool {
	return a.Role == RoleAdmin
}

// AuthRequest defines the request structure for auth entity
type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse defines the response structure for auth entity
type AuthResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
