package routes

import (
	"medicine-app/models"

	"github.com/gin-gonic/gin"
)

func MedicineRoutes(router *gin.Engine, store *models.MedicineStore) {
	router.GET("/medicines", func(ctx *gin.Context) {
	})

	router.GET("/medicines/:id", func(ctx *gin.Context) {

	})

	router.PUT("/medicines/:id", func(ctx *gin.Context) {

	})

	router.DELETE("/medicines/:id", func(ctx *gin.Context) {

	})
}
