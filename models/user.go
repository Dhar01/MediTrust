package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"

	Email string = "email"
	Phone string = "phone"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         Name      `json:"name"`
	Email        string    `json:"email"`
	Age          int32     `json:"age"`
	Phone        string    `json:"phone"`
	Address      Address   `json:"address"`
	HashPassword string
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateUserDTO struct {
	Name    Name    `json:"name" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Age     int32   `json:"age" binding:"required"`
	Phone   string  `json:"phone" binding:"required"`
	Address Address `json:"address" binding:"required"`
}

type SignUpUser struct {
	Name     Name   `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int32  `json:"age" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type UpdateUserDTO struct {
	Name    *Name    `json:"name,omitempty"`
	Email   string   `json:"email,omitempty"`
	Age     *int32   `json:"age,omitempty"`
	Phone   string   `json:"phone,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	Country       string `json:"country"`
	City          string `json:"city"`
	StreetAddress string `json:"street_address"`
	PostalCode    string `json:"postal_code"`
}

type Name struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(ctx context.Context, user CreateUserDTO) (User, error)
	FindUserByID(ctx context.Context, userID uuid.UUID) (User, error)
	FindUserByKey(ctx context.Context, key, value string) (User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user UpdateUserDTO) (User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	SignUpUser(ctx context.Context, user SignUpUser) error
	LogInUser(ctx context.Context, login LogIn) error
}

type UserRepository interface {
	Create(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	Update(ctx context.Context, user User) (User, error)
	FindByID(ctx context.Context, userID uuid.UUID) (User, error)
	FindUser(ctx context.Context, key, value string) (User, error)
	SignUp(ctx context.Context, user User) error
}
