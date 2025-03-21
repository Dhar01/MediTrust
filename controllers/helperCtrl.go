package controllers

import (
	"fmt"
	"medicine-app/models"
	"medicine-app/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getRefreshToken(ctx *gin.Context) (string, bool) {
	refreshToken, err := ctx.Cookie(models.TokenRefresh)
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		return "", false
	}

	if refreshToken == "" {
		ctx.Status(http.StatusUnauthorized)
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

func getParamID(ctx *gin.Context, identifier string) (uuid.UUID, bool) {
	id := ctx.Param(identifier)
	actualID, err := uuid.Parse(id)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid %s: %v", identifier, err))
		return uuid.Nil, false
	}

	return actualID, true
}

func getUserID(ctx *gin.Context) (uuid.UUID, bool) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("UserID not found"))
		return uuid.Nil, false
	}

	return userID.(uuid.UUID), true
}

func getRole(ctx *gin.Context) (string, bool) {
	role, exists := ctx.Get("role")
	if !exists {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("role not found"))
		return "", false
	}

	return role.(string), true
}
