package service

import (
	"context"
	"medicine-app/internal/auth"
	"medicine-app/models/dto"
	"medicine-app/repository"
	"time"
)

type generalService struct {
	repo   repository.GeneralRepository
	secret string
}

type GeneralService interface {
	ResetMedicineService(ctx context.Context) error
	ResetUserService(ctx context.Context) error
	ResetAddressService(ctx context.Context) error
	GenerateToken(ctx context.Context, refreshToken string) (dto.TokenResponseDTO, error)
	RevokeRefreshToken(ctx context.Context, refreshToken string) error
}

func NewGeneralService(genRepo repository.GeneralRepository, secret string) GeneralService {
	return &generalService{
		repo:   genRepo,
		secret: secret,
	}
}

func (gs *generalService) ResetMedicineService(ctx context.Context) error {
	return gs.repo.ResetMedicineRepo(ctx)
}

func (gs *generalService) ResetUserService(ctx context.Context) error {
	return gs.repo.ResetUserRepo(ctx)
}

func (gs *generalService) ResetAddressService(ctx context.Context) error {
	return gs.repo.ResetAddressRepo(ctx)
}

func (gs *generalService) GenerateToken(ctx context.Context, refreshToken string) (dto.TokenResponseDTO, error) {
	user, err := gs.repo.FindUserFromToken(ctx, refreshToken)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	accessToken, err := auth.GenerateAccessToken(user.ID, user.Role, gs.secret, time.Minute*15)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	return dto.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (gs *generalService) RevokeRefreshToken(ctx context.Context, refreshToken string) error {
	return gs.repo.RevokeToken(ctx, refreshToken)
}
