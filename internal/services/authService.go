package services

import (
	"context"
	"medicine-app/config"
	auth "medicine-app/internal/api/auth_gen"
	"medicine-app/internal/database"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, cfg *config.Config) {
	authserver := newAuthServer(cfg.DB)
	handlers := auth.NewStrictHandler(authserver, []auth.StrictMiddlewareFunc{})
	auth.RegisterHandlers(router, handlers)
}

type authService struct {
	DB *database.Queries
}

var _ auth.StrictServerInterface = (*authService)(nil)

func newAuthServer(db *database.Queries) *authService {
	if db == nil {
		panic("database can't be nil")
	}

	return &authService{
		DB: db,
	}
}

func (as *authService) LogInUser(ctx context.Context, request auth.LogInUserRequestObject) (auth.LogInUserResponseObject, error)

func (as *authService) PostAuthLogout(ctx context.Context, request auth.PostAuthLogoutRequestObject) (auth.PostAuthLogoutResponseObject, error)

func (as *authService) RequestPasswordReset(ctx context.Context, request auth.RequestPasswordResetRequestObject) (auth.RequestPasswordResetResponseObject, error)

func (as *authService) UpdatePasswordReset(ctx context.Context, request auth.UpdatePasswordResetRequestObject) (auth.UpdatePasswordResetResponseObject, error)

func (as *authService) UserSignUpHandler(ctx context.Context, request auth.UserSignUpHandlerRequestObject) (auth.UserSignUpHandlerResponseObject, error)

func (as *authService) VerifySignedUpUser(ctx context.Context, request auth.VerifySignedUpUserRequestObject) (auth.VerifySignedUpUserResponseObject, error)
