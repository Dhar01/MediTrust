package middleware

import (
	"medicine-app/internal/auth"
	"medicine-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AdminAuth(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, role, err := getAuth(ctx, secret)
		if err != nil || role != models.Admin {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", id)
		ctx.Next()
	}
}

func IsLoggedIn(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, role, err := getAuth(ctx, secret)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", id)
		ctx.Set("role", role)
		ctx.Next()
	}
}

func getAuth(ctx *gin.Context, secret string) (uuid.UUID, string, error) {
	token, err := auth.GetBearerToken(ctx.Request.Header)
	if err != nil {
		return wrapNilError(err)
	}

	id, role, err := auth.ValidateJWT(token, secret)
	if err != nil {
		return wrapNilError(err)
	}

	return id, role, nil
}

func wrapNilError(err error) (uuid.UUID, string, error) {
	return uuid.Nil, "", err
}
