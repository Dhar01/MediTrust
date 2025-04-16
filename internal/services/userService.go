package services

import (
	"context"
	"medicine-app/config"
	users "medicine-app/internal/api/users_gen"
	"medicine-app/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime/types"
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

func (us *userService) DeleteUserByID(ctx context.Context, request users.DeleteUserByIDRequestObject) (users.DeleteUserByIDResponseObject, error) {
	if err := us.DB.DeleteUser(ctx, request.UserID); err != nil {
		return users.DeleteUserByID401Response{}, err
	}

	return users.DeleteUserByID204Response{}, nil
}

func (us *userService) FetchUserInfoByID(ctx context.Context, request users.FetchUserInfoByIDRequestObject) (users.FetchUserInfoByIDResponseObject, error) {
	user, err := us.DB.GetUserByID(ctx, request.UserID)
	if err != nil {
		return users.FetchUserInfoByID404Response{}, err
	}

	return users.FetchUserInfoByID200JSONResponse(users.FetchUserInfoResponse{
		Name: &users.FullName{
			FirstName: &user.FirstName,
			LastName:  &user.LastName,
		},
		Age:      &user.Age,
		Email:    (*types.Email)(&user.Email),
		IsActive: &user.Verified,
		Phone:    &user.Phone,
		Role:     &user.Role,
	}), nil

}

func (us *userService) UpdateUserInfoByID(ctx context.Context, request users.UpdateUserInfoByIDRequestObject) (users.UpdateUserInfoByIDResponseObject, error) {
	oldInfo, err := us.DB.GetUserByID(ctx, request.UserID)
	if err != nil {
		return users.UpdateUserInfoByID401Response{}, err
	}

	updateInfo, err := us.DB.UpdateUser(ctx, database.UpdateUserParams{
		FirstName: updateField(*request.Body.Name.FirstName, oldInfo.FirstName),
		LastName:  updateField(*request.Body.Name.LastName, oldInfo.LastName),
		Phone:     updateField(*request.Body.Phone, oldInfo.Phone),
		Age:       *updateIntPointerField(request.Body.Age, &oldInfo.Age),
	})
	if err != nil {
		return users.UpdateUserInfoByID500Response{}, err
	}

	return users.UpdateUserInfoByID202JSONResponse(users.UpdateUserResponse{
		Name: users.FullName{
			FirstName: &updateInfo.FirstName,
			LastName:  &updateInfo.LastName,
		},
		Email:    types.Email(updateInfo.Email),
		IsActive: updateInfo.Verified,
		Phone:    updateInfo.Phone,
		Age:      updateInfo.Age,
		Role:     updateInfo.Role,
	}), nil
}

func updateField(newValue, oldValue string) string {
	if newValue == "" {
		return oldValue
	}

	return newValue
}

func updateIntPointerField(newValue, oldValue *int32) *int32 {
	if newValue == nil {
		return oldValue
	}

	return newValue
}
