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

func (mc *medicineController) HandlerCreateMedicine(ctx *gin.Context) {
	var newMedicine models.CreateMedicineDTO

	if err := ctx.ShouldBindJSON(&newMedicine); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	medicine, err := mc.MedicineService.CreateMedicine(ctx, newMedicine)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, medicine)
}

func (mc *medicineController) HandlerGetMedicines(ctx *gin.Context) {
	medicines, err := mc.MedicineService.GetMedicines(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
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
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, medicine)
}

func (mc *medicineController) HandlerUpdateMedicine(ctx *gin.Context) {
	id, ok := getMedicineID(ctx)
	if !ok {
		return
	}

	var updateMed models.UpdateMedicineDTO

	if err := ctx.ShouldBindJSON(&updateMed); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	medicine, err := mc.MedicineService.UpdateMedicine(ctx, id, updateMed)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
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
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func getMedicineID(ctx *gin.Context) (uuid.UUID, bool) {
	medID := ctx.Param("medID")
	id, err := uuid.Parse(medID)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return uuid.Nil, false
	}

	return id, true
}
