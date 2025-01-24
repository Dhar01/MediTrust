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

// planning to use gin's default binding for validation

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         Name      `json:"name" binding:"required,dive"`
	Email        string    `json:"email" binding:"required,email"`
	Age          int32     `json:"age" binding:"required,gte=18"`
	Phone        string    `json:"phone" binding:"required,len=11"` // for BD phone
	Address      Address   `json:"address" binding:"required,dive"`
	HashPassword string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SignUpUser struct {
	Name     Name   `json:"name" binding:"required,dive"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Age      int32  `json:"age" binding:"required,gte=18"`
	Phone    string `json:"phone" binding:"required,len=11"`
}

type UpdateUserDTO struct {
	Name    *Name    `json:"name,omitempty" binding:"omitempty,dive"`
	Email   string   `json:"email,omitempty" biding:"omitempty,email"`
	Age     *int32   `json:"age,omitempty" binding:"omitempty,gte=18"`
	Phone   string   `json:"phone,omitempty" binding:"omitempty,len=11"`
	Address *Address `json:"address,omitempty" binding:"omitempty,dive"`
}

type Address struct {
	Country       string `json:"country" binding:"required"`
	City          string `json:"city" binding:"required"`
	StreetAddress string `json:"street_address" binding:"required"`
	PostalCode    string `json:"postal_code" binding:"required"`
}

type Name struct {
	FirstName string `json:"firstname" binding:"required,min=4"`
	LastName  string `json:"lastname" binding:"required,min=4"`
}

type LogIn struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserService interface {
	FindUserByID(ctx context.Context, userID uuid.UUID) (User, error)
	FindUserByKey(ctx context.Context, key, value string) (User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user UpdateUserDTO) (User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	SignUpUser(ctx context.Context, user SignUpUser) error // should act as CreateUser
	LogInUser(ctx context.Context, login LogIn) error
}

type UserRepository interface {
	Delete(ctx context.Context, userID uuid.UUID) error
	Update(ctx context.Context, user User) (User, error)
	FindByID(ctx context.Context, userID uuid.UUID) (User, error)
	FindUser(ctx context.Context, key, value string) (User, error)
	SignUp(ctx context.Context, user User) error // should act as CreateUser
	FindPass(ctx context.Context, email string) (string, error)
}
