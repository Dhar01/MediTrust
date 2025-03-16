package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	"medicine-app/middleware"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

const (
	cartBase = "/carts"
	cartID   = cartBase + "/:cartID"
	itemID   = cartID + "items/:itemID"
)

func CartRoute(router *gin.RouterGroup, cfg *config.Config) {
	cartService := service.NewCartService(cfg.DB)
	cartCtrl := controllers.NewCartController(cartService)

	// create a cart or add an item
	router.POST(cartBase, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerAddToCart)

	// update the quantity of an item
	router.PATCH(itemID, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerUpdateCartItem)

	// get the cart data
	router.GET(cartBase, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerGetCart)

	// remove an item from the cart
	router.DELETE(itemID, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerRemoveItem)

	// delete the entire cart
	router.DELETE(cartBase, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerDeleteCart)

}
