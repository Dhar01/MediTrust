package controllers

import (
	"log"
	"medicine-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type controller struct {
	GeneralService models.GeneralService
	Platform       string
}

func NewController(service models.GeneralService, platform string) *controller {
	return &controller{
		GeneralService: service,
		Platform:       platform,
	}
}

func (ctrl *controller) HandlerReset(ctx *gin.Context) {
	log.Println(ctrl.Platform)

	if ctrl.Platform != models.Dev {
		ctx.Status(http.StatusForbidden)
		return
	}

	if err := ctrl.GeneralService.ResetMedicineService(ctx.Request.Context()); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	log.Println("Medicine database reset successfully!")

	if err := ctrl.GeneralService.ResetAddressService(ctx.Request.Context()); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	log.Println("Address database reset successfully!")

	if err := ctrl.GeneralService.ResetUserService(ctx.Request.Context()); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	log.Println("User database reset successfully!")

	ctx.Status(http.StatusNoContent)
}

func (ctrl *controller) HandlerRefresh(ctx *gin.Context) {
	refreshToken, ok := getRefreshToken(ctx)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	token, err := ctrl.GeneralService.GenerateToken(ctx.Request.Context(), refreshToken)
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	ctx.SetCookie(models.TokenRefresh, token.RefreshToken, int(time.Hour*7*24), models.RootPath, models.DomainName, true, true)
	ctx.JSON(http.StatusCreated, models.ServerResponse{
		AccessToken: token.AccessToken,
	})
}

func (ctrl *controller) HandlerRevoke(ctx *gin.Context) {
	refreshToken, ok := getRefreshToken(ctx)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return
	}
	if err := ctrl.GeneralService.RevokeRefreshToken(ctx.Request.Context(), refreshToken); err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	ctx.SetCookie(models.TokenRefresh, models.TokenNull, -1, models.RootPath, models.DomainName, true, true)
	ctx.Status(http.StatusNoContent)
}

func getRefreshToken(ctx *gin.Context) (string, bool) {
	refreshToken, err := ctx.Cookie(models.TokenRefresh)
	if err != nil {
		return "", false
	}

	if refreshToken == "" {
		return "", false
	}

	return refreshToken, true
}

func errorResponse(ctx *gin.Context, code int, err error) {
	ctx.Error(err)
	ctx.JSON(code, models.ErrorResponse{
		Message: err.Error(),
		Code:    code,
	})
}
