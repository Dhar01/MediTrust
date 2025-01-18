package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
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
	router.POST(medicineBase, medCtrl.HandlerCreateMedicine)

	// PUT route for medicines - Admin only
	router.PUT(medicineBaseByID, medCtrl.HandlerUpdateMedicine)

	// DELETE route for medicine - Admin only
	router.DELETE(medicineBaseByID, medCtrl.HandlerDeleteMedicine)
}
