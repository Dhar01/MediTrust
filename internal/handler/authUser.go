package handler

import (
	"context"
	"errors"
	"medicine-app/internal/handler/authUser_gen"
	"medicine-app/internal/errs"
	"medicine-app/internal/services"
	"medicine-app/models"

	"github.com/oapi-codegen/runtime/types"
)

type authAPI struct {
	authService   services.AuthService
	publicService services.PublicService
}

func newAuthUserAPI(
	authSrv services.AuthService,
	pubSrv services.PublicService,
) *authAPI {
	if authSrv == nil {
		panic("user service can't be empty/nil")
	}

	if pubSrv == nil {
		panic("public service can't be empty/nil")
	}

	return &authAPI{
		authService:   authSrv,
		publicService: pubSrv,
	}
}

var _ auth_gen.StrictServerInterface = (*authAPI)(nil)

func (api *authAPI) DeleteUserByID(ctx context.Context, request auth_gen.DeleteUserByIDRequestObject) (auth_gen.DeleteUserByIDResponseObject, error) {
	if err := api.authService.DeleteUserByID(ctx, request.UserID); err != nil {
		return auth_gen.UnauthorizedAccessErrorResponse{}, err
	}

	return auth_gen.DeleteUserByID204Response{}, nil
}

func (api *authAPI) FetchUserInfoByID(ctx context.Context, request auth_gen.FetchUserInfoByIDRequestObject) (auth_gen.FetchUserInfoByIDResponseObject, error) {
	user, err := api.publicService.FetchUserByID(ctx, request.UserID)
	if errors.Is(err, errs.ErrUserNotExist) {
		return auth_gen.BadRequestErrorResponse{}, err
	}

	if err != nil {
		return auth_gen.InternalServerErrorResponse{}, err
	}

	return auth_gen.FetchUserInfoByID200JSONResponse(
		auth_gen.FetchUserInfoResponse{
			Name: &auth_gen.FullName{
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

func (api *authAPI) UpdateUserInfoByID(ctx context.Context, request auth_gen.UpdateUserInfoByIDRequestObject) (auth_gen.UpdateUserInfoByIDResponseObject, error) {
	user, err := api.authService.UpdateUserInfoByID(ctx, request.UserID, models.UpdateUserRequest{
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
		return auth_gen.BadRequestErrorResponse{}, errs.ErrUserNotExist
	}

	if err != nil {
		return auth_gen.InternalServerErrorResponse{}, err
	}

	return auth_gen.UpdateUserInfoByID202JSONResponse(
		auth_gen.UpdateUserResponse{
			Name: auth_gen.FullName{
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
