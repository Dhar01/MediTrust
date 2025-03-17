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
	itemID   = cartID + "/items/:itemID"
)

func CartRoute(router *gin.RouterGroup, cfg *config.Config) {
	cartService := service.NewCartService(cfg.DB)
	cartCtrl := controllers.NewCartController(cartService)

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
