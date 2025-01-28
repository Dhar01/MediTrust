package controllers

import (
	"log"
	"medicine-app/models"
	"net/http"
	"time"

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
	refreshToken, ok := getRefreshToken(ctx)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	token, err := ctrl.GeneralService.GenerateToken(ctx, refreshToken)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusUnauthorized, errorMsg(err))
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
	if err := ctrl.GeneralService.RevokeRefreshToken(ctx, refreshToken); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "can't revoke refresh token"})
		return
	}

	ctx.SetCookie(models.TokenRefresh, models.TokenNull, -1, models.RootPath, models.DomainName, true, true)
	ctx.Status(http.StatusNoContent)
}

func getRefreshToken(ctx *gin.Context) (string, bool) {
	refreshToken, err := ctx.Cookie(models.TokenRefresh)
	if err != nil {
		ctx.Error(err)
		return "", false
	}

	if refreshToken == "" {
		return "", false
	}

	return refreshToken, true
}
