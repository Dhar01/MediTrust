package dto

import (
	"medicine-app/models/db"
	"time"

	"github.com/google/uuid"
)

// UserResponseDTO represents the response structure for user details
// @Description Contains user profile information returned from API
type UserResponseDTO struct {
	ID        uuid.UUID  `json:"id"`
	Name      db.Name    `json:"name"`
	Email     string     `json:"email"`
	Age       int32      `json:"age"`
	Phone     string     `json:"phone"`
	Address   db.Address `json:"address"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// UpdateUserDTO represents the request body for updating a user profile
// @Description Contains optional fields for updating user information
type UpdateUserDTO struct {
	Name    db.Name     `json:"name,omitempty" binding:"omitempty"`
	Email   string      `json:"email,omitempty" binding:"omitempty,email" example:"user@example.com"`
	Age     int32       `json:"age,omitempty" binding:"omitempty,gte=18" example:"18"`
	Phone   string      `json:"phone,omitempty" binding:"omitempty,len=11" example:"01234567891"`
	Address *db.Address `json:"address,omitempty" binding:"omitempty" `
}

// SignUpUser represents the request body for user registration
// @Description Contains required fields for creating a new user
type SignUpUserDTO struct {
	Name     db.Name `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email" example:"user@example.com"`
	Password string  `json:"password" binding:"required,min=8" example:"SecurePass123"`
	Age      int32   `json:"age" binding:"required,gte=18" example:"25"`
	Phone    string  `json:"phone" binding:"required,len=11" example:"01234567891"` // for BD phone
}
