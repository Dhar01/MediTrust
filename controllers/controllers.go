package controllers

import (
	"log"
	"medicine-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var dev = "dev"

type controller struct {
	GeneralService models.GeneralService
	Platform       string
}

func errorMsg(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func NewController(service models.GeneralService, platform string) *controller {
	return &controller{
		GeneralService: service,
		Platform:       platform,
	}
}

func (ctrl *controller) HandlerReset(ctx *gin.Context) {
	log.Println(ctrl.Platform)

	if ctrl.Platform != dev {
		ctx.Status(http.StatusForbidden)
		return
	}

	if err := ctrl.GeneralService.ResetMedicineService(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	if err := ctrl.GeneralService.ResetAddressService(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	if err := ctrl.GeneralService.ResetUserService(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
