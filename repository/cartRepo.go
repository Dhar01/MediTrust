package repository

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models"
)

type cartRepository struct {
	DB *database.Queries
}

func NewCartRepository(db *database.Queries) models.CartRepository {
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
