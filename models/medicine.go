package models

import (
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
	Created_at   time.Time
	Updated_at   time.Time
}

type MedicineBody struct {
	Name         string `json:"name"`
	Dosage       string `json:"dosage"`
	Description  string `json:"description"`
	Manufacturer string `json:"manufacturer"`
	Price        int32  `json:"price"`
	Stock        int32  `json:"stock"`
}

type MedicineID struct {
	ID uuid.UUID `json:"medID"`
}

// type MedicineService interface {
// 	CreateMedicine(medicine Medicine) (Medicine, error)
// 	DeleteMedicine(medID uuid.UUID) error
// 	UpdateMedicine(medID uuid.UUID) (Medicine, error)
// 	GetMedicines() ([]Medicine, error)
// 	GetMedicineByID(medID uuid.UUID) (Medicine, error)
// }
