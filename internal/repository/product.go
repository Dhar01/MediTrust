package repository

import (
	"context"
	"medicine-app/internal/database/model"

	"github.com/google/uuid"
)

type ProductRepo interface {
	Create(ctx context.Context, product model.Product) (*model.Product, error)
	Update(ctx context.Context, product model.Product) (*model.Product, error)
	Delete(ctx context.Context, productID uuid.UUID) error

	FetchByID(ctx context.Context, productID uuid.UUID) (*model.Product, error)
	FetchByName(ctx context.Context, name string) error
	FetchList(ctx context.Context) ([]model.Product, error)
}
