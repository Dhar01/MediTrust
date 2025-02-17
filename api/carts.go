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

	router.POST(cartBase, middleware.AdminAuth(cfg.SecretKey), cartCtrl.HandlerCreateCart)
}
