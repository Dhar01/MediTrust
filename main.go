package main

import (
	models "medicine-app/models"
	routes "medicine-app/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	store := models.NewMedicineStore()

	routes.MedicineRoutes(route, store)

	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// listen and serve on :8080
	route.Run()
}
