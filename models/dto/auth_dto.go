package dto

import "github.com/google/uuid"

// SignUpResponse represents the response after a successful user registration
// @Description Contains the unique ID of the newly created user
type SignUpResponseDTO struct {
	ID uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"` // unique ID of the user
}

// LogIn represents the request body for user login
// @Description Contains credentials required for authentication
type LogInDTO struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"SecurePass123"`
}

// RequestResetPass represents the request body for password reset
// @Description when a user forgets his/her password, they can hit this endpoint to reset
type RequestResetPassDTO struct {
	Email string `json:"email" binding:"required,email" example:"user@example.com"`
}

type ResetPassDTO struct {
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
type ServerResponseDTO struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
}
