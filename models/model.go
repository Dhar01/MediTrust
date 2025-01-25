package models

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type GeneralService interface {
	ResetMedicineService(ctx context.Context) error
	ResetUserService(ctx context.Context) error
	ResetAddressService(ctx context.Context) error
	GenerateToken(ctx context.Context, headers http.Header) (ResponseTokenDTO, error)
	RevokeRefreshToken(ctx context.Context, headers http.Header) error
}

type GeneralRepository interface {
	ResetMedicineRepo(ctx context.Context) error
	ResetUserRepo(ctx context.Context) error
	ResetAddressRepo(ctx context.Context) error
	RevokeToken(ctx context.Context, token string) error
	CreateRefreshToken(ctx context.Context, token string, id uuid.UUID) error
	FindUserFromToken(ctx context.Context, token string) (User, error)
}
