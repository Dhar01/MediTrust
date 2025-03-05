package repository

import (
	"context"
	"medicine-app/models/db"

	"github.com/google/uuid"
)

// * REPOSITORY LAYER * //

// UserRepository defines the DB operations for users
// @Description Interface for user database transactions
type UserRepository interface {
	SignUp(ctx context.Context, user db.User) (db.User, error) // should act as CreateUser
	Delete(ctx context.Context, userID uuid.UUID) error
	Update(ctx context.Context, user db.User) (db.User, error)
	UpdatePassword(ctx context.Context, password string, id uuid.UUID) error
	FindByID(ctx context.Context, userID uuid.UUID) (db.User, error)
	FindUser(ctx context.Context, key, value string) (db.User, error)
	CountAvailableUsers(ctx context.Context) (int, error)
}

type AuthRepository interface {
	CreateRefreshToken(ctx context.Context, token string, id uuid.UUID) error
	FindUserFromToken(ctx context.Context, token string) (db.User, error)
	Logout(ctx context.Context, id uuid.UUID) error
}

type VerificationRepository interface {
	GetVerification(ctx context.Context, id uuid.UUID) (bool, error)
	SetVerification(ctx context.Context, id uuid.UUID) error
	GetUserRole(ctx context.Context, id uuid.UUID) (string, error)
}
