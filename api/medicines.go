package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	middleware "medicine-app/middleware"
	repo "medicine-app/repository"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

const (
	medicineBase     = "/medicines"
	medicineBaseByID = medicineBase + "/:medID"
)

func MedicineRoutes(router *gin.RouterGroup, cfg *config.Config) {
	medRepo := repo.NewMedicineRepository(cfg.DB)
	medService := service.NewMedicineService(medRepo)
	medCtrl := controllers.NewMedicineController(medService)

	// GET route for medicines - PUBLIC
	router.GET(medicineBase, medCtrl.HandlerGetMedicines)
	router.GET(medicineBaseByID, medCtrl.HandlerGetMedicineByID)

	// Admin-Only routes
	adminGroup := router.Group(medicineBase, middleware.AdminAuth(cfg.SecretKey))
	{
		// POST route for medicines
		adminGroup.POST(medicineBase, medCtrl.HandlerCreateMedicine)

		// PUT route for medicines
		adminGroup.PUT(medicineBaseByID, medCtrl.HandlerUpdateMedicine)

		// DELETE route for medicine
		adminGroup.DELETE(medicineBaseByID, medCtrl.HandlerDeleteMedicine)
	}
}
