package main

import (
	"log"
	"medicine-app/config"
	ctrl "medicine-app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.ConnectDB()

	medCtrl := ctrl.NewMedicineController(cfg.DB)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/medicines", medCtrl.GetMedicines)
	router.GET("/medicines/:medID", medCtrl.GetMedicineByID)
	router.PUT("/medicines/:medID", medCtrl.UpdateMedicine)
	router.POST("/medicines", medCtrl.CreateMedicineHandler)
	router.DELETE("/medicines/:medID", medCtrl.DeleteMedicine)

	port := ":8080"

	if err := router.Run(port); err != nil {
		log.Fatalf("cant' run in port %s: %v", port, err)
	}
}
