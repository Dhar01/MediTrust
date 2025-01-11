package controllers

import (
	"net/http"

	"medicine-app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type medicineController struct {
	MedicineService models.MedicineService
}

func NewMedicineController(service models.MedicineService) *medicineController {
	return &medicineController{
		MedicineService: service,
	}
}

func (mc *medicineController) HandlerCreateMedicineHandler(ctx *gin.Context) {
	var newMedicine models.Medicine

	if err := ctx.ShouldBindJSON(&newMedicine); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	medicine, err := mc.MedicineService.CreateMedicine(ctx, newMedicine)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusCreated, medicine)
}

func (mc *medicineController) HandlerGetMedicines(ctx *gin.Context) {
	medicines, err := mc.MedicineService.GetMedicines(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, medicines)
}

func (mc *medicineController) HandlerGetMedicineByID(ctx *gin.Context) {
	id, ok := getMedicineID(ctx)
	if !ok {
		return
	}

	medicine, err := mc.MedicineService.GetMedicineByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, medicine)
}

func (mc *medicineController) HandlerUpdateMedicine(ctx *gin.Context) {
	id, ok := getMedicineID(ctx)
	if !ok {
		return
	}

	var updateMed models.Medicine

	if err := ctx.ShouldBindJSON(&updateMed); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	medicine, err := mc.MedicineService.UpdateMedicine(ctx, id, updateMed)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusAccepted, medicine)
}

func (mc *medicineController) HandlerDeleteMedicine(ctx *gin.Context) {
	id, ok := getMedicineID(ctx)
	if !ok {
		return
	}

	if err := mc.MedicineService.DeleteMedicine(ctx, id); err != nil {
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
