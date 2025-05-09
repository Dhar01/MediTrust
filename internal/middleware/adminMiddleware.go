package middleware

import (
	"medicine-app/config"
	"medicine-app/internal/auth"
	"medicine-app/internal/errs"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type HandlerFunc func(ctx echo.Context, request any) (response any, err error)

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

func (m *middleware) IsUser(next HandlerFunc, operationID string) HandlerFunc {
	return func(ctx echo.Context, request any) (any, error) {
		if err := m.requireLoggedIn(ctx); err != nil {
			return wrapErr(err)
		}

		role, ok := ctx.Get("role").(string)
		if !ok {
			return wrapErr(errs.ErrUserRoleNotFound)
		}

		if role != "customer" {
			return wrapErr(errs.ErrUserNotExist)
		}

		return next(ctx, request)
	}
}

func (m *middleware) IsAdmin(next HandlerFunc, operationID string) HandlerFunc {
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
			return wrapErr(errs.ErrUserRoleNotFound)
		}

		if role != "admin" {
			return wrapErr(errs.ErrUserNotAdmin)
		}

		return next(ctx, request)
	}
}

func (m *middleware) requireLoggedIn(ctx echo.Context) error {
	token, err := getToken(ctx)
	if err != nil {
		return err
	}

	id, role, err := m.getUserAuth(token)
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

func (m *middleware) getUserAuth(token string) (uuid.UUID, string, error) {
	id, role, err := auth.ValidateAccessToken(token, m.cfg.SecretKey)
	if err != nil {
		return wrapNilError(err)
	}

	return id, role, nil
}

func getToken(ctx echo.Context) (string, error) {
	token, err := auth.GetBearerToken(ctx.Request().Header)
	if err != nil {
		return "", err
	}

	return token, nil
}

func wrapNilError(err error) (uuid.UUID, string, error) {
	return uuid.Nil, "", err
}

func wrapErr(err error) (any, error) {
	return nil, err
}
