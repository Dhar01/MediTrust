package controllers

import (
	"fmt"
	"medicine-app/models/dto"
	service "medicine-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	cartID, ok := getParamID(ctx, "cartID")
	if !ok {
		return
	}

	itemID, ok := getParamID(ctx, "itemID")
	if !ok {
		return
	}

	var quantity dto.QuantityControlDTO

	if err := ctx.ShouldBindBodyWithJSON(quantity); err != nil {
		errorResponse(ctx, http.StatusBadRequest, fmt.Errorf("can't process the request"))
		return
	}

	if err := cc.cartService.UpdateCartItem(ctx.Request.Context(), cartID, itemID, quantity.Quantity); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	ctx.Status(http.StatusOK)
}

// HandlerGetCart will fetch the data of a cart by UserID
//
//	@Summary		Fetch the data of a cart by userID
//	@Description	Fetch the data of a cart using userID
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	db.Cart					"details of a cart"
//	@Failure		401	{object}	dto.ErrorResponseDTO	"The user is not authorized"
//	@Failure		500	{object}	dto.ErrorResponseDTO	'Internal	server	error"
//	@Router			/cart [get]
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

func (cc *cartController) HandlerRemoveItem(ctx *gin.Context) {
	itemID, ok := getParamID(ctx, "itemID")
	if !ok {
		return
	}

	cartID, ok := getParamID(ctx, "cartID")
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
//	@Router			/cart [delete]
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
