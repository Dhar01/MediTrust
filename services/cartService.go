package service

import (
	"context"
	"medicine-app/repository"
)

type cartService struct {
	repo repository.CartRepository
}

// CartService defines the business logic interface for cart management
// @Description Interface for cart-related business logic
type CartService interface {
	AddToCart(ctx context.Context) error
	GetCart(ctx context.Context) error
	UpdateCart(ctx context.Context) error
	RemoveItemFromCart(ctx context.Context) error
}

func NewCartService(repo repository.CartRepository) CartService {
	if repo == nil {
		panic("repository can't be nil")
	}

	return &cartService{
		repo: repo,
	}
}

func (cs *cartService) AddToCart(ctx context.Context) error {
	return nil
}

func (cs *cartService) GetCart(ctx context.Context) error {
	return nil
}

func (cs *cartService) UpdateCart(ctx context.Context) error {
	return nil
}

func (cs *cartService) RemoveItemFromCart(ctx context.Context) error {
	return nil
}
