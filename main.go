package main

import (
	"log"
	"medicine-app/config"
	ctrl "medicine-app/controllers"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	medService := service.NewMedicineService(cfg.DB)
	medCtrl := ctrl.NewMedicineController(medService)

	resetCtrl := ctrl.NewController(cfg.DB, cfg.Platform)

	userService := service.NewUserService(cfg.DB)
	userCtrl := ctrl.NewUserController(userService)


	router := gin.Default()
	router.SetTrustedProxies(nil)

	// medicines
	router.GET("/medicines", medCtrl.HandlerGetMedicines)
	router.GET("/medicines/:medID", medCtrl.HandlerGetMedicineByID)
	router.PUT("/medicines/:medID", medCtrl.HandlerUpdateMedicine)
	router.POST("/medicines", medCtrl.HandlerCreateMedicine)
	router.DELETE("/medicines/:medID", medCtrl.HandlerDeleteMedicine)

	// users
	router.GET("/users/:userID", userCtrl.HandlerGetUserByID)
	router.PUT("/users/:userID", userCtrl.HandlerUpdateUser)
	router.POST("/users", userCtrl.HandlerCreateUser)
	router.DELETE("/users/:userID", userCtrl.HandlerDeleteUser)

	router.POST("/reset", resetCtrl.HandlerReset)

	port := ":8080"

	if err := router.Run(port); err != nil {
		log.Fatalf("cant' run in port %s: %v", port, err)
	}
}
