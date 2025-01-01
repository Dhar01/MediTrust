package main

import (
	"medicine-app/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	medicineApp := MedicineApp{
		Router: r,
	}

	r.POST("/medicine", medicineApp.CreateMedicine)

	r.Run(":8080")

	// listen and serve on :8080

}

type Medicine struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Dosage       string `json:"dosage"`
	Manufacturer string `json:"manufacturer"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	Created_at   time.Time
	Updated_at   time.Time
}

type MedicineApp struct {
	DB     *database.Queries
	Router *gin.Engine
}

func (medApp MedicineApp) CreateMedicine(ctx *gin.Context) {
	var newMedicine Medicine

	if err := ctx.ShouldBindJSON(&newMedicine); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := medApp.DB.CreateMedicine(ctx.Request.Context(), database.CreateMedicineParams{
		Name:         newMedicine.Name,
		Dosage:       newMedicine.Dosage,
		Manufacturer: newMedicine.Manufacturer,
		Price:        int32(newMedicine.Price),
		Stock:        int32(newMedicine.Stock),
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, newMedicine)
}

// func (medApp MedicineApp) GetMedicine(ctx *gin.Context) {
// }
