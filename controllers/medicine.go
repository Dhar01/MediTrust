package controllers

import (
	"net/http"

	"medicine-app/internal/database"
	"medicine-app/models"

	"github.com/gin-gonic/gin"
)

type MedicineController struct {
	DB *database.Queries
}

func (mc *MedicineController) CreateMedicineHandler(ctx *gin.Context) {
	var newMedicine models.MedicineBody

	if err := ctx.ShouldBindJSON(&newMedicine); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	medicine, err := mc.DB.CreateMedicine(ctx, database.CreateMedicineParams{
		Name:         newMedicine.Name,
		Dosage:       newMedicine.Dosage,
		Manufacturer: newMedicine.Manufacturer,
		Price:        newMedicine.Price,
		Stock:        newMedicine.Stock,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusCreated, medicine)
}

func (mc *MedicineController) UpdateMedicine(ctx *gin.Context) {
	var updateMed models.MedicineBody

	if err := ctx.ShouldBindJSON(&updateMed); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	medicine, err := mc.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
		ID:           updateMed.ID,
		Name:         updateMed.Name,
		Dosage:       updateMed.Dosage,
		Manufacturer: updateMed.Manufacturer,
		Description:  updateMed.Description,
		Price:        updateMed.Price,
		Stock:        updateMed.Stock,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusAccepted, medicine)
}

func (mc *MedicineController) DeleteMedicine(ctx *gin.Context) {
	var medID models.MedicineID

	if err := ctx.ShouldBindJSON(&medID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	if err := mc.DB.DeleteMedicine(ctx, medID.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
