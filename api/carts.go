package api

import (
	"fmt"
	"medicine-app/config"
	"medicine-app/middleware"
	"medicine-app/models/dto"
	service "medicine-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	cartBase = "/carts"
	cartID   = cartBase + "/:cartID"
	itemID   = cartID + "/items/:itemID"
)

func CartRoute(router *gin.RouterGroup, cfg *config.Config) {
	cartService := service.NewCartService(cfg.DB)
	cartCtrl := NewCartController(cartService)

	// Register route for creating a cart if not exists or add an item if cart exists
	router.POST(cartBase, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerAddToCart)

	// Register route for updating the quantity of an item in the cart
	router.PATCH(itemID, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerUpdateCartItem)

	// Register route for getting the cart data
	router.GET(cartBase, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerGetCart)

	// Register route for removing an item from the cart
	router.DELETE(itemID, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerRemoveItem)

	// Register route for deleting the entire cart
	router.DELETE(cartBase, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerDeleteCart)

}

type cartController struct {
	cartService service.CartService
}

func NewCartController(service service.CartService) *cartController {
	return &cartController{
		cartService: service,
	}
}

func (cc *cartController) HandlerAddToCart(ctx *gin.Context) {
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}

	var item dto.AddItemToCartDTO

	if err := ctx.ShouldBindBodyWithJSON(item); err != nil {
		errorResponse(ctx, http.StatusBadRequest, fmt.Errorf("can't process the request"))
		return
	}

	cartID, err := cc.cartService.AddToCart(ctx, userID, item)
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, fmt.Errorf("can't create cart"))
		return
	}

	ctx.JSON(http.StatusCreated, dto.CartResponseDTO{
		CartID:  cartID,
		Message: "item added successfully",
	})
}

func (cc *cartController) HandlerUpdateCartItem(ctx *gin.Context) {
	// extract cartID from the URL
	cartID, ok := getParamID(ctx, "cartID")
	if !ok {
		return
	}

	// extract itemID from the URL
	itemID, ok := getParamID(ctx, "itemID")
	if !ok {
		return
	}

	var quantity dto.QuantityControlDTO
	if err := ctx.ShouldBindBodyWithJSON(quantity); err != nil {
		errorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	// calling the service layer to update cart item quantity
	if err := cc.cartService.UpdateCartItem(ctx.Request.Context(), cartID, itemID, quantity.Quantity); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("failed to update item quantity: %v", err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (cc *cartController) HandlerGetCart(ctx *gin.Context) {
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}

	cart, err := cc.cartService.GetCart(ctx.Request.Context(), userID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("cart details not found"))
		return
	}

	ctx.JSON(http.StatusOK, cart)
}

// HandlerRemoveItem will remove an item using cartID and itemID
//
//	@Summary		Remove an item from the cart
//	@Description	Remove an item from the cart
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//	@Param			cartID	path		string					true	"cartID"
//	@Param			itemID	path		string					true	"medID"
//	@Success		204		{string}	string					"status no content"
//	@Failure		400		{object}	dto.ErrorResponseDTO	"cartID/userID not found"
//	@Failure		500		{object}	dto.ErrorResponseDTO	"can't remove the item"
//	@Router			/carts/:cartID/items/:itemID [delete]
func (cc *cartController) HandlerRemoveItem(ctx *gin.Context) {
	cartID, ok := getParamID(ctx, "cartID")
	if !ok {
		return
	}

	itemID, ok := getParamID(ctx, "itemID")
	if !ok {
		return
	}

	if err := cc.cartService.RemoveItemFromCart(ctx.Request.Context(), cartID, itemID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("can't remove the item"))
		return
	}

	ctx.Status(http.StatusNoContent)
}

// HandlerDeleteCart will delete the cart instance using userID
//
//	@Summary		Delete a cart using userID
//	@Description	Delete a cart using userID
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//	@Success		204	{string}	string					"status no content"
//	@Failure		401	{object}	dto.ErrorResponseDTO	"The user is not authorized"
//	@Failure		500	{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/carts [delete]
func (cc *cartController) HandlerDeleteCart(ctx *gin.Context) {
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}

	if err := cc.cartService.DeleteCart(ctx.Request.Context(), userID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("cart not found"))
		return
	}

	ctx.Status(http.StatusNoContent)
}
