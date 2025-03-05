package controllers

import (
	"log"
	"medicine-app/models"
	"medicine-app/models/dto"
	service "medicine-app/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type controller struct {
	GeneralService service.GeneralService
	Platform       string
}

func NewController(service service.GeneralService, platform string) *controller {
	return &controller{
		GeneralService: service,
		Platform:       platform,
	}
}

// HandlerReset resets the databases in the development environment.
//
//		@Summary		Reset all databases (development only)
//		@Description	This endpoint resets the medicine, address, and user databases.
//	             It is restricted to the development environment only.
//		@Tags			general
//		@Accept			json
//		@Success		204	"Database reset successfully"
//		@Failure		403	{object}	models.ErrorResponse	"Forbidden – Not allowed outside development environment"
//		@Failure		500	{object}	models.ErrorResponse	"Internal server error"
//		@Router			/reset [post]
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

// HandlerRefresh generates a new access token using a refresh token.
//
//	@Summary		Generate a new access token
//	@Description	This endpoint retrieves the refresh token from the cookie and generates a new access token.
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	models.ServerResponse	"Access token generated successfully"
//	@Failure		401	{object}	models.ErrorResponse	"Unauthorized request"
//	@Router			/refresh [post]
//	@Security		ApiKeyAuth
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
	ctx.JSON(http.StatusCreated, dto.ServerResponseDTO{
		AccessToken: token.AccessToken,
	})
}

// HandlerRevoke revokes the refresh token and logs the user out.
//
//		@Summary		Revoke refresh token
//		@Description	This endpoint revokes the refresh token, effectively logging them out.
//	             The refresh token is retrieved from the cookie and invalidated.
//		@Tags			authentication
//		@Accept			json
//		@Produce		json
//		@Success		204	"Refresh token revoked successfully"
//		@Failure		401	{object}	models.ErrorResponse	"Unauthorized – Invalid or missing refresh token"
//		@Router			/revoke [post]
//		@Security		ApiKeyAuth
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
	ctx.JSON(code, dto.ErrorResponseDTO{
		Message: err.Error(),
		Code:    code,
	})
}
