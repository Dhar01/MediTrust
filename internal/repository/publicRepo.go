package repository

import (
	"context"
	"errors"
	"medicine-app/internal/database/user/userDB"
	"medicine-app/internal/errs"
	"medicine-app/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PublicRepository interface {
	FetchUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	CountUsers(ctx context.Context) (int64, error)
	CreateUser(ctx context.Context, user models.User) (*models.User, error)

	SaveRefreshToken(ctx context.Context, id uuid.UUID, token string) error
	SetActive(ctx context.Context, id uuid.UUID) error
}

type publicRepo struct {
	DB *userDB.Queries
}

func NewPublicRepo(db *userDB.Queries) PublicRepository {
	if db == nil {
		panic("userDB can't be empty/nil")
	}

	return &publicRepo{
		DB: db,
	}
}

func (repo *publicRepo) FetchUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := repo.DB.GetUserByID(ctx, id)
	if err != nil {
		return wrapUserErr(err)
	}

	return toUserDomain(user), nil
}

func (repo *publicRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := repo.DB.GetUserByEmail(ctx, email)
	if err != nil {
		return wrapUserErr(err)
	}

	return toUserDomain(user), nil
}

func (repo *publicRepo) CountUsers(ctx context.Context) (int64, error) {
	return repo.DB.CountUsers(ctx)
}

func (repo *publicRepo) SaveRefreshToken(ctx context.Context, id uuid.UUID, token string) error {
	if err := repo.DB.CreateRefreshToken(ctx, userDB.CreateRefreshTokenParams{
		Refreshtoken: token,
		UserID:       id,
	}); err != nil {
		return wrapUserSpecErr(err)
	}

	return nil
}

func (repo *publicRepo) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	userInfo, err := repo.DB.CreateUser(ctx, userDB.CreateUserParams{
		FirstName:    user.Name.FirstName,
		LastName:     user.Name.LastName,
		Email:        user.Email,
		Age:          user.Age,
		Phone:        user.Phone,
		Role:         user.Role,
		PasswordHash: user.HashPassword,
	})
	if err != nil {
		return wrapUserErr(err)
	}

	return toUserDomain(userInfo), nil
}

func (repo *publicRepo) SetActive(ctx context.Context, id uuid.UUID) error {
	if err := repo.DB.SetVerified(ctx, id); err != nil {
		return wrapUserSpecErr(err)
	}

	return nil
}

func toUserDomain(dbUser userDB.User) *models.User {
	return &models.User{
		Id: dbUser.ID,
		Name: models.FullName{
			FirstName: dbUser.FirstName,
			LastName:  dbUser.LastName,
		},
		Role:         dbUser.Role,
		Email:        dbUser.Email,
		Phone:        dbUser.Phone,
		Age:          dbUser.Age,
		IsActive:     dbUser.Verified,
		HashPassword: dbUser.PasswordHash,
	}
}

func wrapUserErr(err error) (*models.User, error) {
	return nil, wrapUserSpecErr(err)
}

func wrapUserSpecErr(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return errs.ErrNotFound
	}

	return err
}
