package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// ProductType defines the type for products available on the store
type ProductType string

const (
	ProductMedicine ProductType = "medicine"
	ProductMedical  ProductType = "medical_instrument"
)

// Product defines the product entity for the store
type Product struct {
	ID           uuid.UUID
	Name         string
	Manufacturer string
	Description  string
	Price        int32
	Cost         int32
	Stock        int32
	Type         ProductType
	Metadata     any
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// MedicineMetadata defines the "medicine" type of product entity
type MedicineMetadata struct {
	Dosage     string    `json:"dosage"`
	ExpiryDate time.Time `json:"expiry_date"`
}

// IsTypeMedicine returns true if the product type is medicine
func (p *Product) IsTypeMedicine() bool {
	return p.Type == ProductMedicine
}

// IsTypeMedicalInstrument returns true if the product type is medicine_instrument
func (p *Product) IsTypeMedicalInstrument() bool {
	return p.Type == ProductMedical
}

// ProductRequest defines the request structure for product entity
type ProductRequest struct {
	Name         string          `json:"name" validate:"required"`
	Manufacturer string          `json:"manufacturer" validate:"required"`
	Description  string          `json:"description"`
	Price        int32           `json:"price" validate:"required"`
	Cost         int32           `json:"cost" validate:"required"`
	Stock        int32           `json:"stock" validate:"required"`
	Type         ProductType     `json:"type" validate:"required"`
	Metadata     json.RawMessage `json:"metadata" validate:"required"` // raw JSON payload
}

// ProductResponse defines the response structure for product entity
type ProductResponse struct {
	ID           uuid.UUID   `json:"id"`
	Name         string      `json:"name"`
	Manufacturer string      `json:"manufacturer"`
	Description  string      `json:"description"`
	Price        int32       `json:"price"`
	Type         ProductType `json:"type"`
	Metadata     any         `json:"metadata"` // will be decoded before sending
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}
