package controllers

import (
	"fmt"
	"medicine-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cartController struct {
	CartService models.CartService
}

func NewCartController(service models.CartService) *cartController {
	return &cartController{
		CartService: service,
	}
}

func (cc *cartController) HandlerCreateCart(ctx *gin.Context) {
	_, ok := getUserID(ctx)
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("UserID not found!"))
		return
	}


}
