package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	"medicine-app/middleware"
	repo "medicine-app/repository"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

const cartBase = "/cart"

func CartRoute(router *gin.RouterGroup, cfg *config.Config) {
	cartRepo := repo.NewCartRepository(cfg.DB)
	cartService := service.NewCartService(cartRepo)
	cartCtrl := controllers.NewCartController(cartService)

	// create a cart
	router.POST(cartBase, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerAddToCart)

	// add an item to the cart
	// router.POST(cartBase+"/add", middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerAddToCart)

	// update the quantity of an item
	// router.PUT(cartBase+"/update", middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerUpdateCart)

	// get the cart data
	router.GET(cartBase, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerGetCart)

	// delete the entire cart
	router.DELETE(cartBase, middleware.IsLoggedIn(cfg.SecretKey), cartCtrl.HandlerDeleteCart)

}
