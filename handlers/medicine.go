package handlers

import (
	"medicine-app/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Medicine struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Dosage       string    `json:"dosage"`
	Manufacturer string    `json:"manufacturer"`
	Price        int       `json:"price"`
	Stock        int       `json:"stock"`
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

	if err := medApp.DB.CreateMedicine(ctx, database.CreateMedicineParams{
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

func (medApp MedicineApp) GetMedicine(ctx *gin.Context) {

	medicines, err := medApp.DB.GetMedicines(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorPass(err))
		return
	}

	ctx.JSON(http.StatusOK, medicines)
}

func errorPass(err error) gin.H {
	return gin.H{"error": err.Error()}
}
