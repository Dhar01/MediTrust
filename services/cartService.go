package service

import (
	"context"
	"errors"
	"medicine-app/internal/database"
	"medicine-app/models/db"
	"medicine-app/models/dto"

	"github.com/google/uuid"
)

var (
	errCartNotFound = errors.New("cart not found")
)

type cartService struct {
	DB *database.Queries
}

// CartService defines the business logic interface for cart management
type CartService interface {
	AddToCart(ctx context.Context, userID uuid.UUID, item dto.AddItemToCartDTO) (uuid.UUID, error)
	GetCart(ctx context.Context, userID uuid.UUID) (db.Cart, error)

	UpdateCart(ctx context.Context) error

	RemoveItemFromCart(ctx context.Context, cartID, itemID uuid.UUID) error
	DeleteCart(ctx context.Context, userID uuid.UUID) error
}

func NewCartService(db *database.Queries) CartService {
	if db == nil {
		panic("db can't be nil")
	}

	return &cartService{
		DB: db,
	}
}

func (cs *cartService) AddToCart(ctx context.Context, userID uuid.UUID, item dto.AddItemToCartDTO) (uuid.UUID, error) {

	// cartID, exists := cs.repo.GetCartByID(ctx, userID)
	// if !exists {
	// 	var err error
	// 	cartID, err = cs.repo.CreateCart(ctx, userID)
	// 	if err != nil {
	// 		return uuid.Nil, err
	// 	}
	// }
	// item := db.CartItem{
	// 	CartID:   cartID,
	// 	MedID:    cartItem.MedID,
	// 	Quantity: cartItem.Quantity,
	// }
	// id, err := cs.repo.AddToCart(ctx, item)
	// if err != nil {
	// 	return uuid.Nil, err
	// }

	cartID, err := cs.DB.GetCartByUserID(ctx, userID)
	if err != nil {
		var err error
		cartID, err = cs.DB.CreateCart(ctx, userID)
		if err != nil {
			return uuid.Nil, err
		}
	}

	id, err := cs.DB.AddItemToCart(ctx, database.AddItemToCartParams{
		MedicineID: item.MedID,
		CartID:     cartID,
		Quantity:   item.Quantity,
	})

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (cs *cartService) GetCart(ctx context.Context, userID uuid.UUID) (db.Cart, error) {

	// cart, err := cs.repo.GetCart(ctx, userID)
	// if err != nil {
	// 	return wrapEmptyCartError(err)
	// }
	// for _, item := range cart.Items {
	// 	item.Price = item.Price * item.Quantity
	// }

	cartItems, err := cs.DB.GetCart(ctx, userID)
	if err != nil {
		return wrapEmptyCartError(err)
	}

	cart, err := convertToCart(cartItems)
	if err != nil {
		return wrapEmptyCartError(err)
	}

	return cart, nil
}

func (cs *cartService) UpdateCart(ctx context.Context) error {
	return nil
}

func (cs *cartService) RemoveItemFromCart(ctx context.Context, cartID, itemID uuid.UUID) error {
	if err := cs.DB.RemoveCartItem(ctx, database.RemoveCartItemParams{
		MedicineID: itemID,
		CartID:     cartID,
	}); err != nil {
		return err
	}

	return nil
}

func (cs *cartService) DeleteCart(ctx context.Context, userID uuid.UUID) error {
	return cs.DB.DeleteCart(ctx, userID)
}

func wrapEmptyCartError(err error) (db.Cart, error) {
	return db.Cart{}, err
}

func convertToCart(cartInfo []database.GetCartRow) (db.Cart, error) {
	if len(cartInfo) == 0 {
		return wrapEmptyCartError(errCartNotFound)
	}

	cart := db.Cart{
		ID:         cartInfo[0].CartID,
		Created_At: cartInfo[0].CreatedAt.Time,
		Items:      []db.CartItem{},
	}

	for _, info := range cartInfo {
		// medicine ID is not valid have to be valid
		if info.MedicineID.Valid {
			cart.Items = append(cart.Items, db.CartItem{
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
