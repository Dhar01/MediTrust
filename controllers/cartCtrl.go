package controllers

import (
	"medicine-app/models"

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

}
