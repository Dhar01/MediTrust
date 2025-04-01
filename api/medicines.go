package api

import (
	"net/http"

	"medicine-app/config"
	med "medicine-app/internal/api/medicines_gen"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

func MedicineRoutes(router *gin.RouterGroup, cfg *config.Config) {
	medService := service.NewMedicineService(cfg.DB)
	medServer := newMedicineServer(medService)
	med.RegisterHandlers(router, medServer)
}

/*

The APIs and controllers are bind together.
Belows are mostly internal.

*/

type medicineServer struct {
	medicineService service.MedicineService
}

func newMedicineServer(service service.MedicineService) *medicineServer {
	return &medicineServer{
		medicineService: service,
	}
}

func (mc *medicineServer) CreateNewMedicine(ctx *gin.Context) {
	var newMedicine med.CreateMedicineDTO

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

func (mc *medicineServer) FetchMedicineList(ctx *gin.Context) {
	medicines, err := mc.medicineService.GetMedicines(ctx.Request.Context())
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, medicines)
}

func (mc *medicineServer) FetchMedicineByID(ctx *gin.Context, medID med.MedicineID) {
	medicine, err := mc.medicineService.GetMedicineByID(ctx.Request.Context(), medID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, medicine)
}

func (mc *medicineServer) UpdateMedicineInfoByID(ctx *gin.Context, medID med.MedicineID) {
	var updateMed med.UpdateMedicineDTO

	if err := ctx.ShouldBindJSON(&updateMed); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	medicine, err := mc.medicineService.UpdateMedicine(ctx.Request.Context(), medID, updateMed)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusAccepted, medicine)
}

func (mc *medicineServer) DeleteMedicineByID(ctx *gin.Context, medID med.MedicineID) {
	if err := mc.medicineService.DeleteMedicine(ctx.Request.Context(), medID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
