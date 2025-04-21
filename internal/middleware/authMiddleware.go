package middleware

import (
	"fmt"
	"medicine-app/config"
	med_gen "medicine-app/internal/api/medicines_gen"
	"medicine-app/internal/auth"
	"medicine-app/internal/errs"
	"medicine-app/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// type ctxKey string

// const (
// 	UserIDKey   ctxKey = "userID"
// 	UserRoleKey ctxKey = "role"
// )

// func AdminAuth(secret string) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		_, role, err := getUserAuth(ctx, secret)
// 		if err != nil {
// 			wrapAuthError(ctx, err, "invalid or expired token")
// 			return
// 		}
// 		if role != models.Admin {
// 			wrapAuthError(ctx, fmt.Errorf("invalid user - not admin"), "user not supported")
// 			return
// 		}
// 		ctx.Next()
// 	}
// }

// func IsLoggedIn(secret string) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		id, role, err := getUserAuth(ctx, secret)
// 		if err != nil || id == uuid.Nil || role == "" {
// 			wrapAuthError(ctx, err, "invalid or expired token")
// 			return
// 		}
// 		ctx.Set("user_id", id)
// 		ctx.Set("role", role)
// 		ctx.Next()
// 	}
// }

type middleware struct {
	cfg *config.Config
}

func NewMiddleware(cfg *config.Config) *middleware {
	if cfg == nil {
		panic("configuration can't be empty/nil")
	}

	return &middleware{
		cfg: cfg,
	}
}

func (m *middleware) RequireLoggedIn(next med_gen.StrictHandlerFunc, operationID string) med_gen.StrictHandlerFunc {
	// admin protected
	protected := map[string]bool{
		"FetchMedicineList":      false,
		"FetchMedicineByID":      false,
		"CreateNewMedicine":      true,
		"DeleteMedicineByID":     true,
		"UpdateMedicineInfoByID": true,
	}

	if !protected[operationID] {
		return next
	}

	return func(ctx echo.Context, request any) (any, error) {
		id, role, err := m.getUserAuth(ctx)
		if err != nil {
			wrapErr(err)
		}

		if role == "" {
			return wrapErr(fmt.Errorf("role not found"))
		}

		if role != models.Admin {
			return wrapErr(errs.ErrUserNotAdmin)
		}

		if id == uuid.Nil {
			return wrapErr(fmt.Errorf("id not found"))
		}


		ctx.Set("user_id", id)
		ctx.Set("role", role)

		return next(ctx, request)
	}
}

// func (m *middle) ONLogin(next public_gen.StrictHandlerFunc, operationID string) public_gen.StrictHandlerFunc {
// 	return func(ctx echo.Context, request any) (any, error) {
// 		resp, err := next(ctx, request)
// 		if err != nil {
// 			wrapErr(err)
// 		}
// 		cookie := new(http.Cookie)
// 		cookie.Name = "refresh_token"
// 		// cookie.Value = resp.()
// 	}
// }

// func IsLoggedInMiddleware() medicines.StrictMiddlewareFunc {
// 	return func (
// 		next func(ctx *gin.Context, request interface{}) (interface{}, error),
// 		operationID string,
// 	) func (ctx *gin.Context, request interface{}) (interface{}, error) {
// 		return func(ctx *gin.Context, request interface{}) (interface{}, error) {
// 			start := time.Now()
// 			resp, err := next(ctx, request)
// 			return resp, err
// 		}
// 	}
// }

func (m *middleware) getUserAuth(ctx echo.Context) (uuid.UUID, string, error) {
	token, err := auth.GetBearerToken(ctx.Request().Header)
	if err != nil {
		return wrapNilError(err)
	}

	id, role, err := auth.ValidateAccessToken(token, m.cfg.SecretKey)
	if err != nil {
		return wrapNilError(err)
	}

	return id, role, nil
}

func wrapNilError(err error) (uuid.UUID, string, error) {
	return uuid.Nil, "", err
}

func wrapErr(err error) (any, error) {
	return nil, err
}
