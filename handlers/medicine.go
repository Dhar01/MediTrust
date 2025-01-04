package handlers

import (
	"errors"
	"medicine-app/internal/database"
	"medicine-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var errNoIdProvided = errors.New("no ID provided")

func errorMsg(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func CreateMedicine(app MedicineApp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newMedicine models.Medicine

		if err := ctx.ShouldBindJSON(&newMedicine); err != nil {
			ctx.JSON(http.StatusBadRequest, errorMsg(err))
			return
		}

		medInfo, err := app.DB.CreateMedicine(ctx, database.CreateMedicineParams{
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

		ctx.JSON(http.StatusOK, medInfo)
	}
}

func GetMedicine(app MedicineApp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		medID, err := getMedID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorMsg(err))
			return
		}

		medicine, err := app.DB.GetMedicine(ctx, medID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorMsg(err))
			return
		}

		ctx.JSON(http.StatusOK, medicine)
	}
}

func GetMedicines(app MedicineApp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		medicines, err := app.DB.GetMedicines(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorMsg(err))
			return
		}

		ctx.JSON(http.StatusOK, medicines)
	}
}

func getMedID(ctx *gin.Context) (uuid.UUID, error) {
	id := ctx.Param("medicineID")
	if id == "" {
		return uuid.Nil, errNoIdProvided
	}

	medID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}

	return medID, nil
}

func DeleteMedicine(app MedicineApp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// expect pathValue will contain medicineID
		medID, err := getMedID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorMsg(err))
			return
		}

		if err := app.DB.DeleteMedicine(ctx, medID); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorMsg(err))
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}

func UpdateMedicine(app MedicineApp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		medID, err := getMedID(ctx)
		if medID == uuid.Nil || err != nil {
			ctx.JSON(http.StatusBadRequest, errorMsg(err))
			return
		}

		var medUpdate models.Medicine

		if err := ctx.ShouldBindJSON(&medUpdate); err != nil {
			ctx.JSON(http.StatusBadRequest, errorMsg(err))
			return
		}

		medInfo, err := app.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
			Name:         medUpdate.Name,
			Dosage:       medUpdate.Dosage,
			Manufacturer: medUpdate.Manufacturer,
			Price:        medUpdate.Price,
			Stock:        medUpdate.Stock,
			ID:           medID,
		})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorMsg(err))
			return
		}

		ctx.JSON(http.StatusOK, medInfo)
	}
}
