package product

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	errNotFound        = errors.New("medicine not found")
	errInvalidMedicine = errors.New("invalid medicine data")
)

type medicine struct {
	ID           uuid.UUID
	Name         string
	Manufacturer string
	Dosage       string
	Description  string
	Price        int32
	Stock        int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type medicineRequest struct {
	Name         string  `json:"name" validate:"required"`
	Manufacturer string  `json:"manufacturer" validate:"required"`
	Dosage       string  `json:"dosage" validate:"required"`
	Description  string  `json:"description"`
	Price        float64 `json:"price" validate:"required,gte=0"`
	Stock        int32   `json:"stock" validate:"required,gte=0"`
}

type medicineResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Manufacturer string    `json:"manufacturer"`
	Dosage       string    `json:"dosage"`
	Description  string    `json:"description,omitempty"`
	Price        float64   `json:"price"`
	Stock        int32     `json:"stock"`
}

type medicineService interface {
	CreateMedicine(ctx context.Context, med medicine) (*medicine, error)
	UpdateMedicine(ctx context.Context, medID uuid.UUID, med medicine) (*medicine, error)
	DeleteMedicine(ctx context.Context, medID uuid.UUID) error
	FetchMedicineByID(ctx context.Context, medID uuid.UUID) (*medicine, error)
	FetchMedicineList(ctx context.Context) ([]medicine, error)
}

type medicineRepository interface {
	Create(ctx context.Context, med medicine) (*medicine, error)
	Delete(ctx context.Context, medID uuid.UUID) error
	Update(ctx context.Context, med medicine) (*medicine, error)
	FetchByID(ctx context.Context, medID uuid.UUID) (*medicine, error)
	FetchByName(ctx context.Context, name string) error
	FetchList(ctx context.Context) ([]*medicine, error)
}
