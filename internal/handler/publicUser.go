package handler

import (
	"context"
	"errors"
	"medicine-app/internal/handler/public_gen"
	"medicine-app/internal/errs"
	"medicine-app/internal/services"
	"medicine-app/models"

	"github.com/google/uuid"
)

type publicAPI struct {
	publicService services.PublicService
}

var _ public_gen.StrictServerInterface = (*publicAPI)(nil)

func newPublicAPI(srv services.PublicService) *publicAPI {
	if srv == nil {
		panic("auth service can't be empty/nil")
	}

	return &publicAPI{
		publicService: srv,
	}
}

func (api *publicAPI) ForgetPasswordHandler(ctx context.Context, request public_gen.ForgetPasswordHandlerRequestObject) (public_gen.ForgetPasswordHandlerResponseObject, error) {
	return public_gen.ForgetPasswordHandler202Response{}, nil
}

func (api *publicAPI) ResetPasswordHandler(ctx context.Context, request public_gen.ResetPasswordHandlerRequestObject) (public_gen.ResetPasswordHandlerResponseObject, error) {
	return public_gen.ResetPasswordHandler202Response{}, nil
}

func (api *publicAPI) LogInUserHandler(ctx context.Context, request public_gen.LogInUserHandlerRequestObject) (public_gen.LogInUserHandlerResponseObject, error) {
	resp, err := api.publicService.LogInUser(ctx, string(request.Body.Email), request.Body.Password)

	if errors.Is(err, errs.ErrUserNotExist) || errors.Is(err, errs.ErrUserNotAuthorized) {
		return public_gen.UnauthorizedAccessErrorResponse{}, err
	} else if errors.Is(err, errs.ErrUserInactive) {
		return public_gen.NotFoundErrorResponse{}, err
	} else if errors.Is(err, errs.ErrInternalServer) {
		return public_gen.InternalServerErrorResponse{}, err
	}

	if err != nil {
		return public_gen.BadRequestErrorResponse{}, err
	}

	return public_gen.LogInUserHandler200JSONResponse{
		Body: public_gen.LoginResponse{
			AccessToken: resp.AccessToken,
		},
		Headers: public_gen.LogInUserHandler200ResponseHeaders{
			SetCookie: resp.RefreshToken,
		},
	}, nil
}

func (api *publicAPI) RegisterUserHandler(ctx context.Context, request public_gen.RegisterUserHandlerRequestObject) (public_gen.RegisterUserHandlerResponseObject, error) {
	id, err := api.publicService.SignUp(ctx, models.User{
		Email:    string(request.Body.Email),
		Password: request.Body.Password,
	})

	if errors.Is(err, errs.ErrEmailAlreadyExists) {
		return public_gen.ConflictErrorResponse{}, errs.ErrEmailAlreadyExists
	}

	if err != nil {
		return public_gen.InternalServerErrorResponse{}, err
	}

	return public_gen.RegisterUserHandler201JSONResponse{
		UserId: (*uuid.UUID)(&id),
	}, nil
}

func (api *publicAPI) VerifyUserHandler(ctx context.Context, request public_gen.VerifyUserHandlerRequestObject) (public_gen.VerifyUserHandlerResponseObject, error) {
	if err := api.publicService.VerifyUser(ctx, request.Params.Token); err != nil {
		return public_gen.UnauthorizedAccessErrorResponse{}, err
	}

	return public_gen.VerifyUserHandler200Response{}, nil
}
