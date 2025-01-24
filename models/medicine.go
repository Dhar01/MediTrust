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

type CreateMedicineDTO struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Dosage       string `json:"dosage" binding:"required"`
	Manufacturer string `json:"manufacturer" binding:"required"`
	Price        int32  `json:"price" binding:"required,min=0"`
	Stock        int32  `json:"stock" binding:"required,min=0"`
}

type UpdateMedicineDTO struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Dosage       string `json:"dosage,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Price        *int32 `json:"price,omitempty" binding:"min=0"`
	Stock        *int32 `json:"stock,omitempty" binding:"min=0"`
}

type MedicineService interface {
	CreateMedicine(ctx context.Context, medicine CreateMedicineDTO) (Medicine, error)
	DeleteMedicine(ctx context.Context, medID uuid.UUID) error
	UpdateMedicine(ctx context.Context, medID uuid.UUID, med UpdateMedicineDTO) (Medicine, error)
	GetMedicines(ctx context.Context) ([]Medicine, error)
	GetMedicineByID(ctx context.Context, medID uuid.UUID) (Medicine, error)
}

type MedicineRepository interface {
	Create(ctx context.Context, med Medicine) (Medicine, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, med Medicine) (Medicine, error)
	FindByID(ctx context.Context, id uuid.UUID) (Medicine, error)
	FindAll(ctx context.Context) ([]Medicine, error)
}
