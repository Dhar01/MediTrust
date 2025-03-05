package controllers

import (
	"fmt"
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

func (cc *cartController) HandlerCreateCart(ctx *gin.Context) {
	_, ok := getUserID(ctx)
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("userID not found"))
		return
	}

}
