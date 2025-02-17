package service

import (
	"context"
	"medicine-app/models"
)

type cartService struct {
	Repo models.CartRepository
}

func NewCartService(repo models.CartRepository) models.CartService {
	if repo == nil {
		panic("repository can't be nil")
	}

	return &cartService{
		Repo: repo,
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
