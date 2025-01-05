package controllers

import (
	"net/http"

	"medicine-app/internal/database"
	"medicine-app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type medicineController struct {
	DB *database.Queries
}

func NewMedicineController(db *database.Queries) *medicineController {
	if db == nil {
		panic("database can't be nil")
	}

	return &medicineController{
		DB: db,
	}
}

func (mc *medicineController) CreateMedicineHandler(ctx *gin.Context) {
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

func (mc *medicineController) GetMedicines(ctx *gin.Context) {
	medicines, err := mc.DB.GetMedicines(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, medicines)
}

func (mc *medicineController) GetMedicineByID(ctx *gin.Context) {
	id, ok := getMedicineID(ctx)
	if !ok {
		return
	}

	medicine, err := mc.DB.GetMedicine(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, medicine)
}

func (mc *medicineController) UpdateMedicine(ctx *gin.Context) {
	id, ok := getMedicineID(ctx)
	if !ok {
		return
	}

	var updateMed models.MedicineBody

	if err := ctx.ShouldBindJSON(&updateMed); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	medicine, err := mc.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
		ID:           id,
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

func (mc *medicineController) DeleteMedicine(ctx *gin.Context) {
	id, ok := getMedicineID(ctx)
	if !ok {
		return
	}

	if err := mc.DB.DeleteMedicine(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func getMedicineID(ctx *gin.Context) (uuid.UUID, bool) {
	medID := ctx.Param("medID")
	id, err := uuid.Parse(medID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return uuid.Nil, false
	}

	return id, true
}
