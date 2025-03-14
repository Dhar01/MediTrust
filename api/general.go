package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

func GeneralRoutes(router *gin.RouterGroup, cfg *config.Config) {
	generalService := service.NewGeneralService(cfg.DB, cfg.SecretKey)
	generalCtrl := controllers.NewController(generalService, *cfg)

	router.POST("/reset", generalCtrl.HandlerReset)
	router.POST("/refresh", generalCtrl.HandlerRefresh)
	router.POST("/revoke", generalCtrl.HandlerRevoke)
}
