// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	AdminID      uuid.UUID
	IsSuperAdmin bool
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}

type Medicine struct {
	ID           uuid.UUID
	Name         string
	Dosage       string
	Description  string
	Manufacturer string
	Price        int32
	Stock        int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RefreshToken struct {
	Refreshtoken string
	UserID       uuid.UUID
	ExpiresAt    time.Time
	RevokedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type User struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	Age          int32
	Email        string
	Phone        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserAddress struct {
	UserID        uuid.UUID
	Country       string
	City          string
	StreetAddress string
	PostalCode    sql.NullString
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
