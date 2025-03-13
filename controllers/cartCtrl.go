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
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("userID not found"))
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
		CartID: cartID,
		Message: "item added successfully",
	})
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
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("UserID not found"))
		return
	}

	cart, err := cc.cartService.GetCart(ctx.Request.Context(), userID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("cart details not found"))
		return
	}

	ctx.JSON(http.StatusOK, cart)
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
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("UserID not found"))
		return
	}

	if err := cc.cartService.DeleteCart(ctx.Request.Context(), userID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("cart not found"))
		return
	}

	ctx.Status(http.StatusNoContent)
}
