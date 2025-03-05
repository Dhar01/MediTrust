package db

import (
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
