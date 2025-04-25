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

func (m *middleware) IsAdmin(next med_gen.StrictHandlerFunc, operationID string) med_gen.StrictHandlerFunc {
	// admin protected
	protected := map[string]bool{
		"CreateNewMedicine":      true,
		"DeleteMedicineByID":     true,
		"UpdateMedicineInfoByID": true,
	}

	if needsAdmin, ok := protected[operationID]; !ok || !needsAdmin {
		return next
	}

	return func(ctx echo.Context, request any) (any, error) {
		if err := m.requireLoggedIn(ctx); err != nil {
			return wrapErr(err)
		}

		role, ok := ctx.Get("role").(string)
		if !ok {
			return wrapErr(fmt.Errorf("role not found in context"))
		}

		if role != models.Admin {
			return wrapErr(errs.ErrUserNotAdmin)
		}

		return next(ctx, request)
	}
}

func (m *middleware) requireLoggedIn(ctx echo.Context) error {
	id, role, err := m.getUserAuth(ctx)
	if err != nil {
		return err
	}

	if role == "" {
		return errs.ErrUserRoleNotFound
	}

	if id == uuid.Nil {
		return errs.ErrUserIdNotFound
	}

	ctx.Set("user_id", id)
	ctx.Set("role", role)

	return nil
}

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
