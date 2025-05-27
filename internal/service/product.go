package service

import (
	"context"
	"medicine-app/internal/database/model"

	"github.com/gofrs/uuid"
)

type ProductService interface {
	Create(ctx context.Context, product model.ProductRequest) (model.Product, error)
	Update(ctx context.Context, productID uuid.UUID, product model.Product) (model.Product, error)
	Delete(ctx context.Context, productID uuid.UUID) error
	FetchProductByID(ctx context.Context, productID uuid.UUID) (model.Product, error)
	FetchProductByName(ctx context.Context) (model.Product, error)
	FetchProducts(ctx context.Context) ([]model.Product, error)
}
