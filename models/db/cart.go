package db

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID         uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"` // Unique ID of the cart
	UserID     uuid.UUID
	Created_At time.Time `json:"created_at"`
}

type CartItems struct {
	ID       uuid.UUID
	CartID   uuid.UUID
	MedID    uuid.UUID
	Quantity int32
	Price    int32
}
