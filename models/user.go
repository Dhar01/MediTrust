package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// planning to use gin's default binding for validation

// User represents the user entity stored in the database
// @Description User entity contains details about a user
type User struct {
	ID           uuid.UUID
	Name         Name
	Email        string
	Age          int32
	Role         string
	Phone        string
	Address      Address
	Exist        bool
	HashPassword string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SignUpUser struct {
	Name     Name   `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Age      int32  `json:"age" binding:"required,gte=18"`
	Phone    string `json:"phone" binding:"required,len=11"` // for BD phone
}

type SignUpResponse struct {
	ID uuid.UUID `json:"id"`
}

type LogIn struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type TokenResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string
}

type ServerResponse struct {
	AccessToken string `json:"access_token"`
}

type UserResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      Name      `json:"name"`
	Email     string    `json:"email"`
	Age       int32     `json:"age"`
	Phone     string    `json:"phone"`
	Address   Address   `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserDTO struct {
	Name    Name     `json:"name,omitempty" binding:"omitempty"`
	Email   string   `json:"email,omitempty" binding:"omitempty,email"`
	Age     int32    `json:"age,omitempty" binding:"omitempty,gte=18"`
	Phone   string   `json:"phone,omitempty" binding:"omitempty,len=11"`
	Address *Address `json:"address,omitempty" binding:"omitempty"`
}
type Name struct {
	FirstName string `json:"firstname,omitempty" binding:"omitempty,min=4"`
	LastName  string `json:"lastname,omitempty" binding:"omitempty,min=4"`
}

type Address struct {
	Country       string `json:"country,omitempty" binding:"omitempty"`
	City          string `json:"city,omitempty" binding:"omitempty"`
	StreetAddress string `json:"street_address,omitempty" binding:"omitempty"`
	PostalCode    string `json:"postal_code,omitempty" binding:"omitempty"`
}

// UserService defines the business logic interface for user management
// @Description Interface for user-related business logic
type UserService interface {
	SignUpUser(ctx context.Context, user SignUpUser) (SignUpResponse, error) // should act as CreateUser
	LogInUser(ctx context.Context, login LogIn) (TokenResponseDTO, error)
	FindUserByID(ctx context.Context, userID uuid.UUID) (UserResponseDTO, error)
	FindUserByKey(ctx context.Context, key, value string) (UserResponseDTO, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user UpdateUserDTO) (UserResponseDTO, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	LogoutUser(ctx context.Context, id uuid.UUID) error
	SetVerifiedUser(ctx context.Context, token string) error
}

// UserRepository defines the DB operations for users
// @Description Interface for user database transactions
type UserRepository interface {
	SignUp(ctx context.Context, user User) (User, error) // should act as CreateUser
	Delete(ctx context.Context, userID uuid.UUID) error
	Update(ctx context.Context, user User) (User, error)
	FindByID(ctx context.Context, userID uuid.UUID) (User, error)
	FindUser(ctx context.Context, key, value string) (User, error)
	CreateRefreshToken(ctx context.Context, token string, id uuid.UUID) error
	FindUserFromToken(ctx context.Context, token string) (User, error)
	CountAvailableUsers(ctx context.Context) (int, error)
	GetUserRole(ctx context.Context, id uuid.UUID) (string, error)
	Logout(ctx context.Context, id uuid.UUID) error
	GetVerification(ctx context.Context, id uuid.UUID) (bool, error)
	SetVerification(ctx context.Context, id uuid.UUID) error
}
