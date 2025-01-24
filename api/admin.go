package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	repo "medicine-app/repository"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

const adminBase = "/admins"

func AdminRoutes(router *gin.RouterGroup, cfg *config.Config) {
	adminRepo := repo.NewAdminRepository(cfg.DB)
	adminService := service.NewAdminService(adminRepo)
	adminCtrl := controllers.NewAdminController(adminService)

	// Get Route for admins
	router.GET(adminBase, adminCtrl.GetCarts)
}
