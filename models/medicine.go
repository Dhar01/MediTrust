package models

import (
	"time"

	"github.com/google/uuid"
)

// Medicine Model

type Medicine struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Dosage       string    `json:"dosage"`
	Manufacturer string    `json:"manufacturer"`
	Price        int32     `json:"price"`
	Stock        int32     `json:"stock"`
	Created_at   time.Time
	Updated_at   time.Time
}
