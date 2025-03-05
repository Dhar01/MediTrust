package repository

import (
	"context"
	"medicine-app/internal/database"
)

type cartRepository struct {
	DB *database.Queries
}

// CartRepository defines the DB operations for cart
// @Description Interface for cart database transactions
type CartRepository interface {
	CreateCart(ctx context.Context) error
	AddToCart(ctx context.Context) error
	GetCart(ctx context.Context) error
	UpdateCart(ctx context.Context) error
	DeleteFromCart(ctx context.Context) error
}

func NewCartRepository(db *database.Queries) CartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (cr *cartRepository) CreateCart(ctx context.Context) error {
	return nil
}

func (cr *cartRepository) AddToCart(ctx context.Context) error {
	return nil
}

func (cr *cartRepository) GetCart(ctx context.Context) error {
	return nil
}

func (cr *cartRepository) UpdateCart(ctx context.Context) error {
	return nil
}

func (cr *cartRepository) DeleteFromCart(ctx context.Context) error {
	return nil
}
