package model

import (
	"time"

	"github.com/google/uuid"
)

// ProductType defines the type for products entity on the store
type ProductType string

// Constants for product types
const (
	ProductMedicine ProductType = "medicine"
	ProductMedical  ProductType = "medical_instrument"
)

// Product defines the product entity for the store.
type Product struct {
	ID           uuid.UUID   `json:"id"`
	Name         string      `json:"name" validate:"required"`
	Manufacturer string      `json:"manufacturer" validate:"required"`
	Description  string      `json:"description"`
	Price        int32       `json:"price" validate:"required"`
	Cost         int32       `json:"cost" validate:"required"`
	Stock        int32       `json:"stock" validate:"required"`
	Type         ProductType `json:"type" validate:"required"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

// Medicine struct defines the structure of medicine type of product entity.
type Medicine struct {
	Product
	Dosage     string    `json:"dosage"`
	ExpiryDate time.Time `json:"expiry_date"`
}

// MedicineResponse structures the response for general users
type MedicineResponse struct {
	ID           uuid.UUID   `json:"id"`
	Name         string      `json:"name"`
	Manufacturer string      `json:"manufacturer"`
	Description  string      `json:"description"`
	Price        int32       `json:"price"`
	Type         ProductType `json:"type"`
	Dosage       string      `json:"dosage"`
	ExpiryDate   time.Time   `json:"expiry_date"`
}

// GenericResponse method build the response for general users
func (m *Medicine) GenericResponse() MedicineResponse {
	return MedicineResponse{
		ID:           m.ID,
		Name:         m.Name,
		Manufacturer: m.Manufacturer,
		Description:  m.Description,
		Price:        m.Price,
		Dosage:       m.Dosage,
		ExpiryDate:   m.ExpiryDate,
	}
}

// IsTypeMedicine returns true if the product type is medicine
func (p *Product) IsTypeMedicine() bool {
	return p.Type == ProductMedicine
}

// IsTypeMedicalInstrument returns true if the product type is medicine_instrument
func (p *Product) IsTypeMedicalInstrument() bool {
	return p.Type == ProductMedical
}
