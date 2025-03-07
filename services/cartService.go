package service

import (
	"context"
	"medicine-app/models/db"
	"medicine-app/repository"

	"github.com/google/uuid"
)

type cartService struct {
	repo repository.CartRepository
}

// CartService defines the business logic interface for cart management
type CartService interface {
	AddToCart(ctx context.Context) error
	GetCart(ctx context.Context, userID uuid.UUID) (db.Cart, error)
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

func (cs *cartService) GetCart(ctx context.Context, userID uuid.UUID) (db.Cart, error) {
	cart, err := cs.repo.GetCart(ctx, userID)
	if err != nil {
		return wrapEmptyCartError(err)
	}

	for _, item := range cart.Items {
		item.Price = item.Price * item.Quantity
	}

	return cart, nil
}

func (cs *cartService) UpdateCart(ctx context.Context) error {
	return nil
}

func (cs *cartService) RemoveItemFromCart(ctx context.Context) error {
	return nil
}

func wrapEmptyCartError(err error) (db.Cart, error) {
	return db.Cart{}, err
}