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

	// GET route for medicines
	router.GET(medicineBase, medCtrl.HandlerGetMedicines)
	router.GET(medicineBaseByID, medCtrl.HandlerGetMedicineByID)

	// POST route for medicines - admin only
	router.POST(medicineBase, middleware.AdminAuth(cfg.SecretKey), medCtrl.HandlerCreateMedicine)

	// PUT route for medicines - Admin only
	router.PUT(medicineBaseByID, middleware.AdminAuth(cfg.SecretKey), medCtrl.HandlerUpdateMedicine)

	// DELETE route for medicine - Admin only
	router.DELETE(medicineBaseByID, middleware.AdminAuth(cfg.SecretKey), medCtrl.HandlerDeleteMedicine)
}
