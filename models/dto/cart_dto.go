package dto

import "github.com/google/uuid"

type CartResponseDTO struct {
	CartID  uuid.UUID `json:"cartID" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Message string    `json:"message"`
}

type AddItemToCartDTO struct {
	CartID   uuid.UUID
	MedID    uuid.UUID `json:"medID" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid" `
	Quantity int32     `json:"quantity"`
}
