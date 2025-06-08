package model

import (
	"time"

	"github.com/google/uuid"
)

type Response interface {
	GenericResponse() any
	AdminResponse() any
}

// ProductType defines the type for products entity on the store
type ProductType string

// Constants for product types
const (
	ProductMedicine ProductType = "medicine"
	ProductMedical  ProductType = "medical_instrument"
)

// Product defines the product entity for the store.
type Product struct {
	ID           uuid.UUID
	Name         string
	Manufacturer string
	Description  string
	Price        int32
	Cost         int32
	Stock        int32
	Type         ProductType
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// IsTypeMedicine returns true if the product type is medicine
func (p *Product) IsTypeMedicine() bool {
	return p.Type == ProductMedicine
}

// IsTypeMedicalInstrument returns true if the product type is medicine_instrument
func (p *Product) IsTypeMedicalInstrument() bool {
	return p.Type == ProductMedical
}