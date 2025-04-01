package api

import (
	"log"
	"medicine-app/config"
	"medicine-app/models"
	"medicine-app/models/dto"
	service "medicine-app/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GeneralRoutes(router *gin.RouterGroup, cfg *config.Config) {
	generalService := service.NewGeneralService(cfg.DB, cfg.SecretKey)
	generalCtrl := NewController(generalService, *cfg)

	router.POST("/reset", generalCtrl.HandlerReset)
	router.POST("/refresh", generalCtrl.HandlerRefresh)
	router.POST("/revoke", generalCtrl.HandlerRevoke)
}

type controller struct {
	generalService service.GeneralService
	conf           config.Config
}

func NewController(service service.GeneralService, cfg config.Config) *controller {
	return &controller{
		generalService: service,
		conf:           cfg,
	}
}

func (ctrl *controller) HandlerReset(ctx *gin.Context) {
	log.Println(ctrl.conf.Platform)

	if ctrl.conf.Platform != models.Dev {
		ctx.Status(http.StatusForbidden)
		return
	}

	if err := ctrl.generalService.ResetMedicineService(ctx.Request.Context()); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	log.Println("Medicine database reset successfully!")

	if err := ctrl.generalService.ResetAddressService(ctx.Request.Context()); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	log.Println("Address database reset successfully!")

	if err := ctrl.generalService.ResetUserService(ctx.Request.Context()); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	log.Println("User database reset successfully!")

	ctx.Status(http.StatusNoContent)
}

func (ctrl *controller) HandlerRefresh(ctx *gin.Context) {
	refreshToken, ok := getRefreshToken(ctx)
	if !ok {
		return
	}

	token, err := ctrl.generalService.GenerateToken(ctx.Request.Context(), refreshToken)
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	ctx.SetCookie(models.TokenRefresh, token.RefreshToken, int(time.Hour*7*24), models.RootPath, ctrl.conf.Domain, true, true)
	ctx.JSON(http.StatusCreated, dto.ServerResponseDTO{
		AccessToken: token.AccessToken,
	})
}

// HandlerRevoke revokes the refresh token and logs the user out.
//
//	@Summary		Revoke refresh token
//	@Description	This endpoint revokes the refresh token, effectively logging them out.
//	             The refresh token is retrieved from the cookie and invalidated.
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Success		204	"Refresh token revoked successfully"
//	@Failure		401	{object}	dto.ErrorResponseDTO	"Unauthorized – Invalid or missing refresh token"
//	@Router			/revoke [post]
//	@Security		ApiKeyAuth
func (ctrl *controller) HandlerRevoke(ctx *gin.Context) {
	refreshToken, ok := getRefreshToken(ctx)
	if !ok {
		return
	}

	if err := ctrl.generalService.RevokeRefreshToken(ctx.Request.Context(), refreshToken); err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	ctx.SetCookie(models.TokenRefresh, models.TokenNull, -1, models.RootPath, ctrl.conf.Domain, true, true)
	ctx.Status(http.StatusNoContent)
}
