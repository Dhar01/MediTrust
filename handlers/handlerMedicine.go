package medicines

import (
	"errors"
	"medicine-app/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var errNoIdProvided = errors.New("no ID provided")

type Medicine struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Dosage       string    `json:"dosage"`
	Manufacturer string    `json:"manufacturer"`
	Price        int32     `json:"price"`
	Stock        int32     `json:"stock"`
	Created_at   time.Time
	Updated_at   time.Time
}

func (medApp *MedicineApp) CreateMedicine(ctx *gin.Context) {
	var newMedicine Medicine

	if err := ctx.ShouldBindJSON(&newMedicine); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	medInfo, err := medApp.DB.CreateMedicine(ctx, database.CreateMedicineParams{
		Name:         newMedicine.Name,
		Dosage:       newMedicine.Dosage,
		Manufacturer: newMedicine.Manufacturer,
		Price:        newMedicine.Price,
		Stock:        newMedicine.Stock,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
	}

	ctx.JSON(http.StatusOK, medInfo)
}

func (medApp *MedicineApp) GetMedicine(ctx *gin.Context) {
	medID, err := getMedID(ctx)

	if err == errNoIdProvided {
		medApp.getAllMedicines(ctx)
	} else if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	} else {
		medApp.getSingleMedicine(ctx, medID)
	}
}

func (medApp *MedicineApp) getSingleMedicine(ctx *gin.Context, medID uuid.UUID) {
	medicine, err := medApp.DB.GetMedicine(ctx, medID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
	}

	ctx.JSON(http.StatusOK, medicine)
}

func (medApp *MedicineApp) getAllMedicines(ctx *gin.Context) {
	medicines, err := medApp.DB.GetMedicines(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, medicines)
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

func (medApp *MedicineApp) DeleteMedicine(ctx *gin.Context) {
	// expect pathValue will contain medicineID
	medID, err := getMedID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	if err := medApp.DB.DeleteMedicine(ctx, medID); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (medApp *MedicineApp) UpdateMedicine(ctx *gin.Context) {
	medID, err := getMedID(ctx)
	if medID == uuid.Nil || err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	var medUpdate Medicine

	if err := ctx.ShouldBindJSON(&medUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	medInfo, err := medApp.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
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
