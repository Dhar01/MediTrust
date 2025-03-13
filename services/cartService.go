package service

import (
	"context"
	"medicine-app/models/db"
	"medicine-app/models/dto"
	"medicine-app/repository"

	"github.com/google/uuid"
)

type cartService struct {
	repo repository.CartRepository
}

// CartService defines the business logic interface for cart management
type CartService interface {
	AddToCart(ctx context.Context, userID uuid.UUID, item dto.AddItemToCartDTO) (uuid.UUID, error)
	GetCart(ctx context.Context, userID uuid.UUID) (db.Cart, error)
	UpdateCart(ctx context.Context) error
	RemoveItemFromCart(ctx context.Context) error
	DeleteCart(ctx context.Context, userID uuid.UUID) error
}

func NewCartService(repo repository.CartRepository) CartService {
	if repo == nil {
		panic("repository can't be nil")
	}

	return &cartService{
		repo: repo,
	}
}

func (cs *cartService) AddToCart(ctx context.Context, userID uuid.UUID, cartItem dto.AddItemToCartDTO) (uuid.UUID, error) {
	cartID, exists := cs.repo.GetCartByID(ctx, userID)
	if !exists {
		var err error
		cartID, err = cs.repo.CreateCart(ctx, userID)
		if err != nil {
			return uuid.Nil, err
		}
	}

	item := db.CartItem{
		CartID:   cartID,
		MedID:    cartItem.MedID,
		Quantity: cartItem.Quantity,
	}

	id, err := cs.repo.AddToCart(ctx, item)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
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

func (cs *cartService) DeleteCart(ctx context.Context, userID uuid.UUID) error {
	return cs.repo.DeleteCart(ctx, userID)
}

func wrapEmptyCartError(err error) (db.Cart, error) {
	return db.Cart{}, err
}
