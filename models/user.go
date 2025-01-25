package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	Admin    string = "admin"
	Customer string = "customer"

	Email string = "email"
	Phone string = "phone"
)

// planning to use gin's default binding for validation

type User struct {
	ID           uuid.UUID
	Name         Name
	Email        string
	Age          int32
	Phone        string
	Address      Address
	HashPassword string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

/*
DTO's for interacting with external request and responses
*/
type SignUpUser struct {
	Name     Name   `json:"name" binding:"required,dive"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Age      int32  `json:"age" binding:"required,gte=18"`
	Phone    string `json:"phone" binding:"required,len=11"` // for BD phone
}

type LogIn struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type ResponseTokenDTO struct {
	AccessToken  string
	RefreshToken string
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
	Name    *Name    `json:"name,omitempty" binding:"omitempty,dive"`
	Email   string   `json:"email,omitempty" biding:"omitempty,email"`
	Age     *int32   `json:"age,omitempty" binding:"omitempty,gte=18"`
	Phone   string   `json:"phone,omitempty" binding:"omitempty,len=11"`
	Address *Address `json:"address,omitempty" binding:"omitempty,dive"`
}
type Name struct {
	FirstName string `json:"firstname" binding:"required,min=4"`
	LastName  string `json:"lastname" binding:"required,min=4"`
}
type Address struct {
	Country       string `json:"country" binding:"required"`
	City          string `json:"city" binding:"required"`
	StreetAddress string `json:"street_address" binding:"required"`
	PostalCode    string `json:"postal_code" binding:"required"`
}

type UserService interface {
	SignUpUser(ctx context.Context, user SignUpUser) error // should act as CreateUser
	LogInUser(ctx context.Context, login LogIn) (ResponseTokenDTO, error)
	FindUserByID(ctx context.Context, userID uuid.UUID) (UserResponseDTO, error)
	FindUserByKey(ctx context.Context, key, value string) (UserResponseDTO, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user UpdateUserDTO) (UserResponseDTO, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

type UserRepository interface {
	SignUp(ctx context.Context, user User) error // should act as CreateUser
	Delete(ctx context.Context, userID uuid.UUID) error
	Update(ctx context.Context, user User) (User, error)
	FindByID(ctx context.Context, userID uuid.UUID) (User, error)
	FindUser(ctx context.Context, key, value string) (User, error)
	CreateRefreshToken(ctx context.Context, token string, id uuid.UUID) error
	FindUserFromToken(ctx context.Context, token string) (User, error)
}
