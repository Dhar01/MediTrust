package repository

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models"

	"github.com/google/uuid"
)

type authRepository struct {
	DB *database.Queries
}

func NewAuthRepository(db *database.Queries) models.AuthRepository {
	return &authRepository{
		DB: db,
	}
}

func (ar *authRepository) CreateRefreshToken(ctx context.Context, token string, id uuid.UUID) error {
	return ar.DB.CreateRefreshToken(ctx, database.CreateRefreshTokenParams{
		Refreshtoken: token,
		UserID:       id,
	})
}

func (ar *authRepository) FindUserFromToken(ctx context.Context, token string) (models.User, error) {
	user, err := ar.DB.GetUserFromRefreshToken(ctx, token)
	if err != nil {
		return wrapUserError(err)
	}

	return ar.userWithAddress(ctx, user)
}

func (ar *authRepository) Logout(ctx context.Context, id uuid.UUID) error {
	if err := ar.DB.RevokeTokenByID(ctx, id); err != nil {
		return err
	}

	return nil
}

func (ar *authRepository) userWithAddress(ctx context.Context, user database.User) (models.User, error) {
	address, err := ar.DB.GetAddress(ctx, user.ID)
	if err != nil {
		return wrapUserError(err)
	}

	return toUserDomain(user, address), nil
}
