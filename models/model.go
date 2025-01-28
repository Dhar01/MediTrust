package models

import (
	"context"

	"github.com/google/uuid"
)

var (
	Admin    string = "admin"
	Customer string = "customer"

	Email string = "email"
	Phone string = "phone"

	// Change these two name
	DomainName  string = "www.domain.com"
	CompanyName string = "company"

	RootPath     string = "/"
	TokenRefresh string = "refresh_token"
	TokenNull    string = ""
)

type ReqToken struct {
	RefreshToken string `json:"refresh_token"`
}

type GeneralService interface {
	ResetMedicineService(ctx context.Context) error
	ResetUserService(ctx context.Context) error
	ResetAddressService(ctx context.Context) error
	GenerateToken(ctx context.Context, refreshToken string) (TokenResponseDTO, error)
	RevokeRefreshToken(ctx context.Context, refreshToken string) error
}

type GeneralRepository interface {
	ResetMedicineRepo(ctx context.Context) error
	ResetUserRepo(ctx context.Context) error
	ResetAddressRepo(ctx context.Context) error
	RevokeToken(ctx context.Context, token string) error
	CreateRefreshToken(ctx context.Context, token string, id uuid.UUID) error
	FindUserFromToken(ctx context.Context, token string) (User, error)
}
