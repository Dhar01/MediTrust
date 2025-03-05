package controllers

import (
	"net/http"

	"medicine-app/models/dto"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type medicineController struct {
	medicineService service.MedicineService
}

func NewMedicineController(service service.MedicineService) *medicineController {
	return &medicineController{
		medicineService: service,
	}
}

// CreateMedicine creates a medicine information - Admin only
//
//	@Summary		Creates a medicine info - Admin only
//	@Description	Create a new medicine on the store. Only an admin can create a medicine.
//	@Tags			medicines
//	@Accept			json
//	@Produce		json
//	@Param			medicine	body		dto.CreateMedicineDTO	true	"Create medicine details"
//	@Success		201			{object}	db.Medicine				"Medicine created successfully"
//	@Failure		400			{object}	dto.ErrorResponseDTO	"Invalid input"
//	@Failure		500			{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/medicines [post]
func (mc *medicineController) HandlerCreateMedicine(ctx *gin.Context) {
	var newMedicine dto.CreateMedicineDTO

	if err := ctx.ShouldBindJSON(&newMedicine); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	medicine, err := mc.medicineService.CreateMedicine(ctx.Request.Context(), newMedicine)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, medicine)
}

// GetMedicines retrieves a list of medicines
//
//	@Summary		Get all medicines
//	@Description	Fetch a list of available medicines
//	@Tags			medicines
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]db.Medicine			"List of medicines"
//	@Failure		500	{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/medicines [get]
func (mc *medicineController) HandlerGetMedicines(ctx *gin.Context) {
	medicines, err := mc.medicineService.GetMedicines(ctx.Request.Context())
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, medicines)
}

// GetMedicineByID retrieves a medicine by its ID
//
//	@Summary		Get a medicine info by its ID
//	@Description	Fetch information of a medicine by its ID
//	@Tags			medicines
//	@Accept			json
//	@Produce		json
//	@Param			medID	path		string					true	"Medicine ID"
//	@Success		200		{object}	db.Medicine				"Medicine found"
//	@Failure		400		{object}	dto.ErrorResponseDTO	"Invalid input"
//	@Failure		500		{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/medicines/{medID} [get]
//
//	@example
func (mc *medicineController) HandlerGetMedicineByID(ctx *gin.Context) {
	id, ok := getMedicineID(ctx)
	if !ok {
		return
	}

	medicine, err := mc.medicineService.GetMedicineByID(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, medicine)
}

// UpdateMedicineByID updates a medicine information by its ID
//
//	@Summary		Updates a medicine info by its ID
//	@Description	Updates information of a medicine by its ID
//	@Tags			medicines
//	@Accept			json
//	@Produce		json
//	@Param			medID		path		string					true	"Medicine ID"
//	@Param			medicine	body		dto.UpdateMedicineDTO	true	"Updated medicine details"
//	@Success		202			{object}	db.Medicine				"Medicine updated successfully"
//	@Failure		400			{object}	dto.ErrorResponseDTO	"Invalid input"
//	@Failure		500			{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/medicines/{medID} [put]
func (mc *medicineController) HandlerUpdateMedicineByID(ctx *gin.Context) {
	id, ok := getMedicineID(ctx)
	if !ok {
		return
	}

	var updateMed dto.UpdateMedicineDTO

	if err := ctx.ShouldBindJSON(&updateMed); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	medicine, err := mc.medicineService.UpdateMedicine(ctx.Request.Context(), id, updateMed)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusAccepted, medicine)
}

// DeleteMedicineByID deletes a medicine by its ID
//
//	@Summary		Deletes a medicine info by its ID
//	@Description	Deletes information of a medicine by its ID
//	@Tags			medicines
//	@Accept			json
//	@Produce		json
//	@Param			medID	path		string	true	"Medicine ID"
//	@Success		204		{}			"Medicine deleted successfully"
//	@Failure		400		{object}	dto.ErrorResponseDTO	"Invalid input"
//	@Failure		500		{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/medicines/{medID} [delete]
func (mc *medicineController) HandlerDeleteMedicineByID(ctx *gin.Context) {
	id, ok := getMedicineID(ctx)
	if !ok {
		return
	}

	if err := mc.medicineService.DeleteMedicine(ctx.Request.Context(), id); err != nil {
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
