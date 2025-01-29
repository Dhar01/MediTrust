package middleware

import (
	"fmt"
	"medicine-app/internal/auth"
	"medicine-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AdminAuth(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, role, err := getUserAuth(ctx, secret)
		if err != nil {
			wrapAuthError(ctx, err, "invalid or expired token")
			return
		}

		if role != models.Admin {
			wrapAuthError(ctx, fmt.Errorf("invalid user - not admin"), "user not supported")
			return
		}

		ctx.Set("user_id", id)
		ctx.Next()
	}
}

func IsLoggedIn(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, role, err := getUserAuth(ctx, secret)
		if err != nil || id == uuid.Nil || role == "" {
			wrapAuthError(ctx, err, "invalid or expired token")
			return
		}

		ctx.Set("user_id", id)
		ctx.Set("role", role)
		ctx.Next()
	}
}

func getUserAuth(ctx *gin.Context, secret string) (uuid.UUID, string, error) {
	token, err := auth.GetBearerToken(ctx.Request.Header)
	if err != nil {
		return wrapNilError(err)
	}

	id, role, err := auth.ValidateAccessToken(token, secret)
	if err != nil {
		return wrapNilError(err)
	}

	return id, role, nil
}

func wrapNilError(err error) (uuid.UUID, string, error) {
	return uuid.Nil, "", err
}

func wrapAuthError(ctx *gin.Context, err error, message string) {
	ctx.Error(err)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": message})
}
