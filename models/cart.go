package models

import (
	"context"
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

// * Cart Service *//

// CartService defines the business logic interface for cart management
// @Description Interface for cart-related business logic
type CartService interface {
	AddToCart(ctx context.Context) error
	GetCart(ctx context.Context) error
	UpdateCart(ctx context.Context) error
	RemoveItemFromCart(ctx context.Context) error
}

// * Cart Repository * //

// CartRepository defines the DB operations for cart
// @Description Interface for cart database transactions
type CartRepository interface {
	CreateCart(ctx context.Context) error
	AddToCart(ctx context.Context) error
	GetCart(ctx context.Context) error
	UpdateCart(ctx context.Context) error
	DeleteFromCart(ctx context.Context) error
}
