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

	log.Println("Medicine database reset successfully!")

	if err := ctrl.GeneralService.ResetAddressService(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	log.Println("Address database reset successfully!")

	if err := ctrl.GeneralService.ResetUserService(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	log.Println("User database reset successfully!")

	ctx.Status(http.StatusNoContent)
}

func (ctrl *controller) HandlerRefresh(ctx *gin.Context) {
	var reqToken models.ReqToken

	if err := ctx.ShouldBindJSON(&reqToken); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	token, err := ctrl.GeneralService.GenerateToken(ctx, reqToken.RefreshToken)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusUnauthorized, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusCreated, token)
}

func (ctrl *controller) HandlerRevoke(ctx *gin.Context) {
	if err := ctrl.GeneralService.RevokeRefreshToken(ctx, ctx.Request.Header); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "can't revoke refresh token"})
		return
	}

	ctx.Status(http.StatusOK)
}
