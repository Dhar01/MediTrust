package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Medicine Model

type Medicine struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Dosage       string    `json:"dosage"`
	Manufacturer string    `json:"manufacturer"`
	Price        int32     `json:"price"`
	Stock        int32     `json:"stock"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type MedicineBody struct {
	ID           uuid.UUID `json:"medID"`
	Name         string    `json:"name"`
	Dosage       string    `json:"dosage"`
	Description  string    `json:"description"`
	Manufacturer string    `json:"manufacturer"`
	Price        int32     `json:"price"`
	Stock        int32     `json:"stock"`
}

type MedicineID struct {
	ID uuid.UUID `json:"medID"`
}

type MedicineService interface {
	CreateMedicine(ctx context.Context, medicine Medicine) (Medicine, error)
	DeleteMedicine(ctx context.Context, medID uuid.UUID) error
	UpdateMedicine(ctx context.Context, medID uuid.UUID, med Medicine) (Medicine, error)
	GetMedicines(ctx context.Context) ([]Medicine, error)
	GetMedicineByID(ctx context.Context, medID uuid.UUID) (Medicine, error)
}
