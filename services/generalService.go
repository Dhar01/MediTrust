package service

import (
	"context"
	"medicine-app/internal/auth"
	"medicine-app/models"
	"net/http"
	"time"
)

type generalService struct {
	Repo   models.GeneralRepository
	Secret string
}

func NewGeneralService(genRepo models.GeneralRepository, secret string) models.GeneralService {
	return &generalService{
		Repo:   genRepo,
		Secret: secret,
	}
}

func (gs *generalService) ResetMedicineService(ctx context.Context) error {
	return gs.Repo.ResetMedicineRepo(ctx)
}

func (gs *generalService) ResetUserService(ctx context.Context) error {
	return gs.Repo.ResetUserRepo(ctx)
}

func (gs *generalService) ResetAddressService(ctx context.Context) error {
	return gs.Repo.ResetAddressRepo(ctx)
}

func (gs *generalService) GenerateToken(ctx context.Context, refreshToken string) (models.TokenResponseDTO, error) {
	user, err := gs.Repo.FindUserFromToken(ctx, refreshToken)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	accessToken, err := auth.MakeJWT(user.ID, user.Role, gs.Secret, time.Minute*15)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	return models.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (gs *generalService) RevokeRefreshToken(ctx context.Context, headers http.Header) error {
	authToken, err := auth.GetBearerToken(headers)
	if err != nil {
		return err
	}

	return gs.Repo.RevokeToken(ctx, authToken)
}
