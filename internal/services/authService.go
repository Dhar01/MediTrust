package services

import (
	"context"
	"medicine-app/config"
	"medicine-app/internal/repository"
	"medicine-app/models"

	"github.com/google/uuid"
)

type AuthService interface {
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
	UpdateUserInfoByID(ctx context.Context, id uuid.UUID, updateInfo models.UpdateUserRequest) (*models.User, error)
}

type authUserService struct {
	authRepo   repository.AuthRepository
	publicRepo repository.PublicRepository
	cfg        *config.Config
}

func NewAuthUserService(
	authRepo repository.AuthRepository,
	userRepo repository.PublicRepository,
	cfg *config.Config,
) AuthService {
	return &authUserService{
		authRepo:   authRepo,
		publicRepo: userRepo,
		cfg:        cfg,
	}
}

func (srv *authUserService) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	return wrapUserSpecErr(srv.authRepo.DeleteUserByID(ctx, id))
}

func (srv *authUserService) UpdateUserInfoByID(ctx context.Context, id uuid.UUID, updateInfo models.UpdateUserRequest) (*models.User, error) {
	oldInfo, err := srv.publicRepo.FetchUserByID(ctx, id)
	if err != nil {
		return wrapUserErr(err)
	}

	user, err := srv.authRepo.UpdateUserInfo(ctx, models.User{
		Name: models.FullName{
			FirstName: updateField(updateInfo.Name.FirstName, oldInfo.Name.FirstName),
			LastName:  updateField(updateInfo.Name.LastName, oldInfo.Name.LastName),
		},
		Age:   *updateIntPointerField(&updateInfo.Age, &oldInfo.Age),
		Phone: updateField(updateInfo.Phone, oldInfo.Phone),
		Email: updateField(updateInfo.Email, oldInfo.Email),
	})
	if err != nil {
		return wrapUserErr(err)
	}

	return user, nil
}

// func (as *authService) LogInUser(ctx context.Context, request auth_gen.LogInUserRequestObject) (auth_gen.LogInUserResponseObject, error) {
// 	user, err := as.UserDB.GetUserByEmail(ctx, string(request.Body.Email))
// 	if err != nil {
// 		return auth_gen.NotFoundErrorResponse{}, err
// 	}
// 	ok, err := as.UserDB.GetVerified(ctx, user.ID)
// 	if err != nil {
// 		return auth_gen.UnauthorizedAccessErrorResponse{}, err
// 	}
// 	if !ok {
// 		return auth_gen.UnauthorizedAccessErrorResponse{}, err
// 	}
// 	if err = auth.CheckPasswordHash(request.Body.Password, user.PasswordHash); err != nil {
// 		return auth_gen.UnauthorizedAccessErrorResponse{}, err
// 	}
// 	accessToken, err := auth.GenerateAccessToken(user.ID, user.Role, as.Cfg.SecretKey, time.Minute*15)
// 	if err != nil {
// 		return auth_gen.InternalServerErrorResponse{}, err
// 	}
// 	refreshToken, err := auth.GenerateRefreshToken()
// 	if err != nil {
// 		return auth_gen.InternalServerErrorResponse{}, err
// 	}
// 	if err = as.UserDB.CreateRefreshToken(ctx, userDB.CreateRefreshTokenParams{
// 		Refreshtoken: refreshToken,
// 		UserID:       user.ID,
// 	}); err != nil {
// 		return auth_gen.InternalServerErrorResponse{}, err
// 	}
// 	return auth_gen.LogInUser200JSONResponse{
// 		Body: auth_gen.SignInResponse{
// 			AccessToken: accessToken,
// 		},
// 	}, nil
// }

// func (as *authService) PostAuthLogout(ctx context.Context, request auth_gen.PostAuthLogoutRequestObject) (auth_gen.PostAuthLogoutResponseObject, error) {
// 	// if err := as.UserDB.RevokeTokenByID(ctx, uuid.UUID); err != nil {
// 	// 	return auth_gen.UnauthorizedAccessErrorResponse{}, err
// 	// }
// 	return auth_gen.PostAuthLogout200JSONResponse{}, nil
// }

// func (as *authService) RequestPasswordReset(ctx context.Context, request auth_gen.RequestPasswordResetRequestObject) (auth_gen.RequestPasswordResetResponseObject, error) {
// 	user, err := as.UserDB.GetUserByEmail(ctx, string(request.Body.Email))
// 	if err != nil {
// 		return auth_gen.NotFoundErrorResponse{}, err
// 	}
// 	_, err = auth.GenerateVerificationToken(user.ID, user.Role, as.Cfg.SecretKey)
// 	if err != nil {
// 		return auth_gen.InternalServerErrorResponse{}, err
// 	}
// 	// need to implement sendEmail
// 	return auth_gen.RequestPasswordReset202Response{}, nil
// }

// func (as *authService) UpdatePasswordReset(ctx context.Context, request auth_gen.UpdatePasswordResetRequestObject) (auth_gen.UpdatePasswordResetResponseObject, error) {
// 	userID, err := auth.ValidateVerificationToken(request.Params.Token, as.Cfg.SecretKey)
// 	if err != nil {
// 		return auth_gen.UnauthorizedAccessErrorResponse{}, err
// 	}
// 	pass, err := auth.HashPassword(request.Body.Password)
// 	if err != nil {
// 		return auth_gen.InternalServerErrorResponse{}, err
// 	}
// 	if err := as.UserDB.ResetPassword(ctx, userDB.ResetPasswordParams{
// 		PasswordHash: pass,
// 		ID:           userID,
// 	}); err != nil {
// 		return auth_gen.InternalServerErrorResponse{}, err
// 	}
// 	return auth_gen.UpdatePasswordReset202Response{}, nil
// }

// func (as *authService) VerifySignedUpUser(ctx context.Context, request auth_gen.VerifySignedUpUserRequestObject) (auth_gen.VerifySignedUpUserResponseObject, error) {
// 	id, err := auth.ValidateVerificationToken(request.Params.Token, as.Cfg.SecretKey)
// 	if err != nil {
// 		return auth_gen.UnauthorizedAccessErrorResponse{}, err
// 	}
// 	if err = as.UserDB.SetVerified(ctx, id); err != nil {
// 		return auth_gen.InternalServerErrorResponse{}, err
// 	}
// 	return auth_gen.VerifySignedUpUser200Response{}, nil
// }
