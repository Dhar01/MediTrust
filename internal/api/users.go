package api

import (
	"context"
	"errors"
	user_gen "medicine-app/internal/api/users_gen"
	"medicine-app/internal/errs"
	"medicine-app/internal/services"
	"medicine-app/models"

	"github.com/oapi-codegen/runtime/types"
)

type userAPI struct {
	userService services.UserService
}

func newUserAPI(srv services.UserService) *userAPI {
	if srv == nil {
		panic("user service can't be empty/nil")
	}

	return &userAPI{
		userService: srv,
	}
}

var _ user_gen.StrictServerInterface = (*userAPI)(nil)

func (api *userAPI) DeleteUserByID(ctx context.Context, request user_gen.DeleteUserByIDRequestObject) (user_gen.DeleteUserByIDResponseObject, error) {
	if err := api.userService.DeleteUserByID(ctx, request.UserID); err != nil {
		return user_gen.UnauthorizedAccessErrorResponse{}, err
	}

	return user_gen.DeleteUserByID204Response{}, nil
}

func (api *userAPI) FetchUserInfoByID(ctx context.Context, request user_gen.FetchUserInfoByIDRequestObject) (user_gen.FetchUserInfoByIDResponseObject, error) {
	user, err := api.userService.FetchUserByID(ctx, request.UserID)
	if errors.Is(err, errs.ErrUserNotExist) {
		return user_gen.BadRequestErrorResponse{}, err
	}

	if err != nil {
		return user_gen.InternalServerErrorResponse{}, err
	}

	return user_gen.FetchUserInfoByID200JSONResponse(
		user_gen.FetchUserInfoResponse{
			Name: &user_gen.FullName{
				FirstName: &user.Name.FirstName,
				LastName:  &user.Name.LastName,
			},
			Age:      &user.Age,
			Phone:    &user.Phone,
			Email:    (*types.Email)(&user.Email),
			Role:     &user.Role,
			IsActive: &user.IsActive,
		}), nil
}

func (api *userAPI) UpdateUserInfoByID(ctx context.Context, request user_gen.UpdateUserInfoByIDRequestObject) (user_gen.UpdateUserInfoByIDResponseObject, error) {
	user, err := api.userService.UpdateUserInfoByID(ctx, request.UserID, models.UpdateUserRequest{
		Name: models.FullName{
			FirstName: *request.Body.Name.FirstName,
			LastName:  *request.Body.Name.LastName,
		},
		Age:   *request.Body.Age,
		Phone: *request.Body.Phone,
		Address: models.Address{
			City:          *request.Body.Address.City,
			Country:       *request.Body.Address.Country,
			PostalCode:    *request.Body.Address.PostalCode,
			StreetAddress: *request.Body.Address.StreetAddress,
		},
	})

	if errors.Is(err, errs.ErrUserNotExist) {
		return user_gen.BadRequestErrorResponse{}, errs.ErrUserNotExist
	}

	if err != nil {
		return user_gen.InternalServerErrorResponse{}, err
	}

	return user_gen.UpdateUserInfoByID202JSONResponse(
		user_gen.UpdateUserResponse{
			Name: user_gen.FullName{
				FirstName: &user.Name.FirstName,
				LastName:  &user.Name.LastName,
			},
			Age:      user.Age,
			Email:    types.Email(user.Email),
			Phone:    user.Phone,
			IsActive: user.IsActive,
			Role:     user.Role,
		}), nil
}
