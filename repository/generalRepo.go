package repository

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models/db"

	"github.com/google/uuid"
)

type generalRepository struct {
	DB *database.Queries
}

type GeneralRepository interface {
	ResetMedicineRepo(ctx context.Context) error
	ResetUserRepo(ctx context.Context) error
	ResetAddressRepo(ctx context.Context) error
	RevokeToken(ctx context.Context, token string) error
	CreateRefreshToken(ctx context.Context, token string, id uuid.UUID) error
	FindUserFromToken(ctx context.Context, token string) (db.User, error)
}

func NewGeneralRepository(db *database.Queries) GeneralRepository {
	return &generalRepository{
		DB: db,
	}
}

func (gr *generalRepository) ResetMedicineRepo(ctx context.Context) error {
	return gr.DB.ResetMedicines(ctx)
}

func (gr *generalRepository) ResetUserRepo(ctx context.Context) error {
	return gr.DB.ResetUsers(ctx)
}

func (gr *generalRepository) ResetAddressRepo(ctx context.Context) error {
	return gr.DB.ResetAddress(ctx)
}

func (gr *generalRepository) RevokeToken(ctx context.Context, token string) error {
	return gr.DB.RevokeRefreshToken(ctx, token)
}

func (gr *generalRepository) CreateRefreshToken(ctx context.Context, token string, id uuid.UUID) error {
	return gr.DB.CreateRefreshToken(ctx, database.CreateRefreshTokenParams{
		Refreshtoken: token,
		UserID:       id,
	})
}

func (gr *generalRepository) FindUserFromToken(ctx context.Context, token string) (db.User, error) {
	user, err := gr.DB.GetUserFromRefreshToken(ctx, token)
	if err != nil {
		return wrapUserError(err)
	}

	return toUser(user), nil
}
