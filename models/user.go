package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      Name      `json:"first_name"`
	Email     string    `json:"email"`
	Age       int32     `json:"age"`
	Phone     string    `json:"phone"`
	Address   Address   `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	PostalCode    int32  `json:"postal_code"`
}

type Name struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type UserService interface {
	CreateUser(ctx context.Context, user User) (User, error)
	FindUserByID(ctx context.Context, userID uuid.UUID) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	FindUserByPhone(ctx context.Context, phone string) (User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user User) (User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}
