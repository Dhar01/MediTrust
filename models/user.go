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

// SignUpUser represents the request body for user registration
// @Description Contains required fields for creating a new user
type SignUpUser struct {
	Name     Name   `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"SecurePass123"`
	Age      int32  `json:"age" binding:"required,gte=18" example:"25"`
	Phone    string `json:"phone" binding:"required,len=11" example:"01234567891"` // for BD phone
}

// SignUpResponse represents the response after a successful user registration
// @Description Contains the unique ID of the newly created user
type SignUpResponse struct {
	ID uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"` // unique ID of the user
}

// LogIn represents the request body for user login
// @Description Contains credentials required for authentication
type LogIn struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"SecurePass123"`
}

// TokenResponseDTO represents the access and refresh tokens returned upon login
// @Description Contains JWT tokens used for authentication
type TokenResponseDTO struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	RefreshToken string `example:"dXNlciByZWZyZXNoIHRva2Vu"`
}

// ServerResponse represents a generic response containing an access token
// @Description Used for authentication responses
type ServerResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
}

// UserResponseDTO represents the response structure for user details
// @Description Contains user profile information returned from API
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

// UpdateUserDTO represents the request body for updating a user profile
// @Description Contains optional fields for updating user information
type UpdateUserDTO struct {
	Name    Name     `json:"name,omitempty" binding:"omitempty"`
	Email   string   `json:"email,omitempty" binding:"omitempty,email" example:"user@example.com"`
	Age     int32    `json:"age,omitempty" binding:"omitempty,gte=18" example:"18"`
	Phone   string   `json:"phone,omitempty" binding:"omitempty,len=11" example:"01234567891"`
	Address *Address `json:"address,omitempty" binding:"omitempty" `
}

// Name represents the user's full name
// @Description Contains first and last name fields
type Name struct {
	FirstName string `json:"firstname,omitempty" binding:"omitempty,min=4" example:"John"`
	LastName  string `json:"lastname,omitempty" binding:"omitempty,min=4" example:"Doe"`
}

// Address represents a user's physical address
// @Description Contains details of the user's location
type Address struct {
	Country       string `json:"country,omitempty" binding:"omitempty" example:"Bangladesh"`
	City          string `json:"city,omitempty" binding:"omitempty" example:"Dhaka"`
	StreetAddress string `json:"street_address,omitempty" binding:"omitempty" example:"123 Main Street"`
	PostalCode    string `json:"postal_code,omitempty" binding:"omitempty" example:"1207"`
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
