package repository

import (
	"context"
	"medicine-app/internal/database"

	"github.com/google/uuid"
)

type verificationRepository struct {
	DB *database.Queries
}

func NewVerificationRepository(db *database.Queries) VerificationRepository {
	return &verificationRepository{
		DB: db,
	}
}

func (vr *verificationRepository) GetUserRole(ctx context.Context, id uuid.UUID) (string, error) {
	role, err := vr.DB.GetRole(ctx, id)
	if err != nil {
		return "", err
	}

	return role, nil
}

func (vr *verificationRepository) GetVerification(ctx context.Context, id uuid.UUID) (bool, error) {
	ok, err := vr.DB.GetVerified(ctx, id)
	if err != nil {
		return ok, err
	}

	return ok, nil
}

func (vr *verificationRepository) SetVerification(ctx context.Context, id uuid.UUID) error {
	if err := vr.DB.SetVerified(ctx, id); err != nil {
		return err
	}

	return nil
}
