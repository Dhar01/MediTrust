package product

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	errNotFound        = errors.New("medicine not found")
	errInvalidMedicine = errors.New("invalid medicine data")
)

func ProductRoutes() {
	repo := newMedicineRepo(nil)
	svc := newMedicineService(repo)
	_ = newMedicineHandler(svc)
}

type medicine struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name         string         `gorm:"type:varchar(100);not null"`
	Manufacturer string         `gorm:"type:varchar(100);not null"`
	Dosage       string         `gorm:"type:varchar(50);not null"`
	Description  string         `gorm:"type:text"`
	Price        int32          `gorm:"not null"`
	Stock        int32          `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type medicineRequest struct {
	Name         string `json:"name" validate:"required"`
	Manufacturer string `json:"manufacturer" validate:"required"`
	Dosage       string `json:"dosage" validate:"required"`
	Description  string `json:"description"`
	Price        int32  `json:"price" validate:"required,gte=0"`
	Stock        int32  `json:"stock" validate:"required,gte=0"`
}

type medicineResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Manufacturer string    `json:"manufacturer"`
	Dosage       string    `json:"dosage"`
	Description  string    `json:"description,omitempty"`
	Price        int32     `json:"price"`
	Stock        int32     `json:"stock"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
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
	FetchList(ctx context.Context) ([]*medicine, error)
}
