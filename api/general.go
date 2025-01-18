package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	"medicine-app/repository"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

func GeneralRoutes(router *gin.RouterGroup, cfg *config.Config) {
	generalRepo := repository.NewGeneralRepository(cfg.DB)
	generalService := service.NewGeneralService(generalRepo)
	generalCtrl := controllers.NewController(generalService, cfg.Platform)

	router.POST("/reset", generalCtrl.HandlerReset)
}
