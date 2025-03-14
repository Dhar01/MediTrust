package service

import (
	"context"
	"medicine-app/internal/auth"
	"medicine-app/internal/database"
	"medicine-app/models/dto"
	"time"
)

type generalService struct {
	DB     *database.Queries
	secret string
}

type GeneralService interface {
	ResetMedicineService(ctx context.Context) error
	ResetUserService(ctx context.Context) error
	ResetAddressService(ctx context.Context) error
	GenerateToken(ctx context.Context, refreshToken string) (dto.TokenResponseDTO, error)
	RevokeRefreshToken(ctx context.Context, refreshToken string) error
}

func NewGeneralService(db *database.Queries, secret string) GeneralService {
	if db == nil {
		panic("database can't be nil")
	}

	return &generalService{
		DB:     db,
		secret: secret,
	}
}

func (gs *generalService) ResetMedicineService(ctx context.Context) error {
	return gs.DB.ResetMedicines(ctx)
}

func (gs *generalService) ResetUserService(ctx context.Context) error {
	return gs.DB.ResetUsers(ctx)
}

func (gs *generalService) ResetAddressService(ctx context.Context) error {
	return gs.DB.ResetAddress(ctx)
}

func (gs *generalService) GenerateToken(ctx context.Context, token string) (dto.TokenResponseDTO, error) {
	user, err := gs.DB.GetUserFromRefreshToken(ctx, token)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	accessToken, err := auth.GenerateAccessToken(user.ID, user.Role, gs.secret, time.Minute*15)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	// * While generate access token, new refreshtoken will be generated
	// refreshToken, err := gs.DB.CreateRefreshToken(ctx, database.CreateRefreshTokenParams{
	// 	Refreshtoken: token,
	// 	UserID: id,
	// })

	return dto.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: token,
	}, nil
}

func (gs *generalService) RevokeRefreshToken(ctx context.Context, refreshToken string) error {
	return gs.DB.RevokeRefreshToken(ctx, refreshToken)
}
