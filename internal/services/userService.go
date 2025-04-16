package services

import (
	"context"
	"medicine-app/config"
	users "medicine-app/internal/api/users_gen"
	"medicine-app/internal/database"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, cfg *config.Config) {
	userServer := newUserServer(cfg.DB)
	handlers := users.NewStrictHandler(userServer, []users.StrictMiddlewareFunc{})
	users.RegisterHandlers(router, handlers)
}

type userService struct {
	DB *database.Queries
}

var _ users.StrictServerInterface = (*userService)(nil)

func newUserServer(db *database.Queries) *userService {
	if db == nil {
		panic("database can't be nil")
	}

	return &userService{
		DB: db,
	}
}

func (us *userService) DeleteUserByID(ctx context.Context, request users.DeleteUserByIDRequestObject) (users.DeleteUserByIDResponseObject, error)

func (us *userService) FetchUserInfoByID(ctx context.Context, request users.FetchUserInfoByIDRequestObject) (users.FetchUserInfoByIDResponseObject, error)

func (us *userService) UpdateUserInfoByID(ctx context.Context, request users.UpdateUserInfoByIDRequestObject) (users.UpdateUserInfoByIDResponseObject, error)
