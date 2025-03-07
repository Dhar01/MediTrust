package repository

import (
	"context"
	"errors"
	"medicine-app/internal/database"
	"medicine-app/models/db"

	"github.com/google/uuid"
)

var (
	errCartNotFound = errors.New("cart not found")
)

type cartRepository struct {
	DB *database.Queries
}

// CartRepository defines the DB operations for cart
// @Description Interface for cart database transactions
type CartRepository interface {
	CreateCart(ctx context.Context) error
	AddToCart(ctx context.Context) error
	GetCart(ctx context.Context, userID uuid.UUID) (db.Cart, error)
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

func (cr *cartRepository) GetCart(ctx context.Context, userID uuid.UUID) (db.Cart, error) {
	cartDetails, err := cr.DB.GetCart(ctx, userID)
	if err != nil {
		return wrapCartError(err)
	}

	return convertToCart(cartDetails)
}

func (cr *cartRepository) UpdateCart(ctx context.Context) error {
	return nil
}

func (cr *cartRepository) DeleteFromCart(ctx context.Context) error {
	return nil
}

func convertToCart(cartInfo []database.GetCartRow) (db.Cart, error) {
	if len(cartInfo) == 0 {
		return wrapCartError(errCartNotFound)
	}

	cart := db.Cart{
		ID:         cartInfo[0].CartID,
		Created_At: cartInfo[0].CreatedAt.Time,
		Items:      []db.CartItems{},
	}

	for _, info := range cartInfo {
		// medicine ID is not valid have to be valid
		if info.MedicineID.Valid {
			cart.Items = append(cart.Items, db.CartItems{
				Serial:   info.ID.Int32,
				MedID:    info.MedicineID.UUID,
				MedName:  info.MedicineName.String,
				Quantity: info.Quantity.Int32,
				Price:    info.Price.Int32,
			})
		}
	}

	return cart, nil
}

func wrapCartError(err error) (db.Cart, error) {
	return db.Cart{}, err
}
