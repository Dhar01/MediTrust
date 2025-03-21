package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	middleware "medicine-app/middleware"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

const (
	medicineBase     = "/medicines"
	medicineBaseByID = medicineBase + "/:medID"
)

func MedicineRoutes(router *gin.RouterGroup, cfg *config.Config) {
	medService := service.NewMedicineService(cfg.DB)
	medCtrl := controllers.NewMedicineController(medService)

	// GET route for medicines - PUBLIC
	router.GET(medicineBase, medCtrl.HandlerGetMedicines)
	router.GET(medicineBaseByID, medCtrl.HandlerGetMedicineByID)

	// POST route for medicines
	router.POST(medicineBase, middleware.AdminAuth(cfg.SecretKey), medCtrl.HandlerCreateMedicine)

	// PUT route for medicines
	router.PUT(medicineBaseByID, middleware.AdminAuth(cfg.SecretKey), medCtrl.HandlerUpdateMedicineByID)

	// DELETE route for medicine
	router.DELETE(medicineBaseByID, middleware.AdminAuth(cfg.SecretKey), medCtrl.HandlerDeleteMedicineByID)
}
