package models

import "github.com/google/uuid"

const CompanyName = ""

var (
	Admin    string = "admin"
	Customer string = "customer"
	Dev      string = "dev"
)

// CreateMedicineDTO defines model for CreateMedicineDTO.
type CreateMedicineDTO struct {
	Description  string
	Dosage       string
	Manufacturer string
	Name         string
	Price        int32
	Stock        int32
}

// Medicine defines model for Medicine.
type Medicine struct {
	Description  string
	Dosage       string
	Id           uuid.UUID
	Manufacturer string
	Name         string
	Price        int32
	Stock        int32
}

// UpdateMedicineDTO defines model for UpdateMedicineDTO.
type UpdateMedicineDTO struct {
	Description  string
	Dosage       string
	Manufacturer string
	Name         string
	Price        int32
	Stock        int32
}
