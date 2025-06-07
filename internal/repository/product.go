package repository

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/internal/database/model"

	"github.com/google/uuid"
)

// ProductRepo defines the interaction methods with the database
type ProductRepo interface {
	Create(ctx context.Context, product model.Product) (*model.Product, error)
	Update(ctx context.Context, product model.Product) (*model.Product, error)
	Delete(ctx context.Context, productID uuid.UUID) error
	FetchByID(ctx context.Context, productID uuid.UUID) (*model.Product, error)
	FetchByName(ctx context.Context, name string) (*model.Product, error)
	FetchList(ctx context.Context) ([]model.Product, error)
}

var _ ProductRepo = (*PostgresRepo)(nil)

type PostgresRepo struct {
	db *database.Queries
}

// NewPostgresRepo will return a new Postgres connection to interact with postgres.
func NewPostgresRepo(db *database.Queries) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

// Create creates a new product instance on the database
func (r *PostgresRepo) Create(ctx context.Context, product model.Product) (*model.Product, error) {
	result, err := r.db.CreateProduct(ctx, database.CreateProductParams{
		Name:         product.Name,
		Manufacturer: product.Manufacturer,
		Description:  product.Description,
		Price:        product.Price,
		Cost:         product.Cost,
		Stock:        product.Stock,
		Type:         toDBProductType(product.Type),
	})

	if err != nil {
		return nil, err
	}

	return toProduct(result), nil
}

// Update updates a product instance on the database
func (r *PostgresRepo) Update(ctx context.Context, product model.Product) (*model.Product, error) {
	updatedProduct, err := r.db.UpdateProduct(ctx, database.UpdateProductParams{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Manufacturer: product.Manufacturer,
		Price:        product.Price,
		Cost:         product.Cost,
		Stock:        product.Stock,
		Type:         toDBProductType(product.Type),
	})

	if err != nil {
		return nil, err
	}

	return toProduct(updatedProduct), nil
}

// Delete deletes a product instance by its productID
func (r *PostgresRepo) Delete(ctx context.Context, productID uuid.UUID) error {
	if err := r.db.DeleteProduct(ctx, productID); err != nil {
		return err
	}

	return nil
}

// FetchByID fetches a product instance by its productID
func (r *PostgresRepo) FetchByID(ctx context.Context, productID uuid.UUID) (*model.Product, error) {
	product, err := r.db.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return toProduct(product), nil
}

// FetchByName returns error if a product with the same name doesn't exist
func (r *PostgresRepo) FetchByName(ctx context.Context, name string) (*model.Product, error) {
	product, err := r.db.GetProductByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return toProduct(product), nil
}

// FetchList fetches the list of product information
func (r *PostgresRepo) FetchList(ctx context.Context) ([]model.Product, error) {
	products, err := r.db.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]model.Product, 0, len(products))

	for _, product := range products {
		results = append(results, *toProduct(product))
	}

	return results, nil
}

// helper: toProduct will convert database.Product type to model.Product type
func toProduct(product database.Product) *model.Product {
	return &model.Product{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Manufacturer: product.Manufacturer,
		Price:        product.Price,
		Cost:         product.Cost,
		Stock:        product.Stock,
		Type:         toModelProductType(product.Type),
		CreatedAt:    product.CreatedAt.Time,
		UpdatedAt:    product.UpdatedAt.Time,
	}
}

// toDBProductType helper mapped the type to the database.ProductType
func toDBProductType(t model.ProductType) database.ProductType {
	return database.ProductType(t)
}

// toModelProductType helper mapped the type to the model.ProductType
func toModelProductType(t database.ProductType) model.ProductType {
	return model.ProductType(t)
}
