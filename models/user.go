package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      Name      `json:"name"`
	Email     string    `json:"email"`
	Age       int32     `json:"age"`
	Phone     string    `json:"phone"`
	Address   Address   `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserDTO struct {
	Name    Name    `json:"name" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Age     int32   `json:"age" binding:"required"`
	Phone   string  `json:"phone" binding:"required"`
	Address Address `json:"address" binding:"required"`
}

type UpdateUserDTO struct {
	Name    *Name    `json:"name,omitempty"`
	Email   string   `json:"email,omitempty"`
	Age     *int32   `json:"age,omitempty"`
	Phone   string   `json:"phone,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Admin struct {
	User
	Role            string
	CanManageUsers  bool
	CanManageOrders bool
	CanManageStore  bool
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

type UserService interface {
	CreateUser(ctx context.Context, user CreateUserDTO) (User, error)
	FindUserByID(ctx context.Context, userID uuid.UUID) (User, error)
	FindUserByKey(ctx context.Context, key, value string) (User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user UpdateUserDTO) (User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

type UserRepository interface {
	Create(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	Update(ctx context.Context, user User) (User, error)
	FindByID(ctx context.Context, userID uuid.UUID) (User, error)
	FindUser(ctx context.Context, key, value string) (User, error)
}
