package model

import (
	"time"

	"github.com/google/uuid"
)

// Medicine struct defines the structure of medicine type of product entity.
type Medicine struct {
	Product
	Dosage     string
	ExpiryDate time.Time
}

// MedicineRequest defines the request structure for medicine
type MedicineRequest struct {
	Name         string      `json:"name"  validate:"required,min=3"`
	Dosage       string      `json:"dosage" validate:"required"`
	Manufacturer string      `json:"manufacturer" validate:"required"`
	Description  string      `json:"description"`
	Price        int32       `json:"price" validate:"required,gt=0"`
	Cost         int32       `json:"cost" validate:"required,gt=0"`
	Stock        int32       `json:"stock" validate:"required,gt=0"`
	Type         ProductType `json:"type" validate:"required,oneof=medicine"`
	ExpiryDate   time.Time   `json:"expiry_date" validate:"required"`
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

// MedicineAdminResponse structures the response for admins
type MedicineAdminResponse struct {
	ID           uuid.UUID   `json:"id"`
	Name         string      `json:"name"`
	Manufacturer string      `json:"manufacturer"`
	Description  string      `json:"description"`
	Price        int32       `json:"price"`
	Cost         int32       `json:"cost"`
	Stock        int32       `json:"stock"`
	Type         ProductType `json:"type"`
	Dosage       string      `json:"dosage"`
	ExpiryDate   time.Time   `json:"expiry_date"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

// GenericResponse method build the response for general users
func (m *Medicine) GenericResponse() any {
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

// AdminResponse method build the response for admins
func (m *Medicine) AdminResponse() any {
	return MedicineAdminResponse{
		ID:           m.ID,
		Name:         m.Name,
		Manufacturer: m.Manufacturer,
		Description:  m.Description,
		Price:        m.Price,
		Cost:         m.Cost,
		Stock:        m.Stock,
		Type:         m.Type,
		Dosage:       m.Dosage,
		ExpiryDate:   m.ExpiryDate,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

// BuildProduct method builds the full product
func (m *Medicine) BuildProduct(product Product) {
	m.Product = product
}
