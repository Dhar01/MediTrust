package main

import (
	"log"
	"medicine-app/config"
	ctrl "medicine-app/controllers"
	repo "medicine-app/repository"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

var (
	apiBase = "/api/v1"

	medicineBase     = apiBase + "/medicines"
	medicineBaseByID = medicineBase + "/:medID"

	usersBase    = apiBase + "/users"
	userBaseByID = usersBase + "/:userID"
)

func main() {
	cfg := config.NewConfig()

	medRepo := repo.NewMedicineRepository(cfg.DB)

	medService := service.NewMedicineService(medRepo)
	medCtrl := ctrl.NewMedicineController(medService)

	resetCtrl := ctrl.NewController(cfg.DB, cfg.Platform)

	userService := service.NewUserService(cfg.DB)
	userCtrl := ctrl.NewUserController(userService)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	// medicines
	router.GET(medicineBase, medCtrl.HandlerGetMedicines)
	router.GET(medicineBaseByID, medCtrl.HandlerGetMedicineByID)
	router.PUT(medicineBaseByID, medCtrl.HandlerUpdateMedicine)
	router.POST(medicineBase, medCtrl.HandlerCreateMedicine)
	router.DELETE(medicineBaseByID, medCtrl.HandlerDeleteMedicine)

	// users
	router.GET(userBaseByID, userCtrl.HandlerGetUserByID)
	router.PUT(userBaseByID, userCtrl.HandlerUpdateUser)
	router.POST(usersBase, userCtrl.HandlerCreateUser)
	router.DELETE(userBaseByID, userCtrl.HandlerDeleteUser)

	router.POST("/reset", resetCtrl.HandlerReset)

	port := ":8080"

	if err := router.Run(port); err != nil {
		log.Fatalf("cant' run in port %s: %v", port, err)
	}
}
