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

type UserRepository interface {
	FetchUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
	UpdateUserInfo(ctx context.Context, updateInfo models.User) (*models.User, error)
}

type userRepo struct {
	DB *userDB.Queries
}

func NewUserRepo(db *userDB.Queries) UserRepository {
	if db == nil {
		panic("userDB can't be empty/nil")
	}

	return &userRepo{
		DB: db,
	}
}

func (repo *userRepo) FetchUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := repo.DB.GetUserByID(ctx, id)
	if err != nil {
		return wrapUserErr(err)
	}

	return toUserDomain(user), nil
}

func (repo *userRepo) UpdateUserInfo(ctx context.Context, updateInfo models.User) (*models.User, error) {
	user, err := repo.DB.UpdateUser(ctx, userDB.UpdateUserParams{
		FirstName: updateInfo.Name.FirstName,
		LastName:  updateInfo.Name.LastName,
		Email:     updateInfo.Email,
		Phone:     updateInfo.Phone,
		Age:       updateInfo.Age,
		ID:        updateInfo.Id,
	})
	if err != nil {
		return wrapUserErr(err)
	}

	return toUserDomain(user), nil
}

func (repo *userRepo) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	return wrapUserSpecErr(repo.DB.DeleteUser(ctx, id))
}

func toUserDomain(dbUser userDB.User) *models.User {
	return &models.User{
		Id: dbUser.ID,
		Name: models.FullName{
			FirstName: dbUser.FirstName,
			LastName:  dbUser.LastName,
		},
		Role:     dbUser.Role,
		Email:    dbUser.Email,
		Phone:    dbUser.Phone,
		Age:      dbUser.Age,
		IsActive: dbUser.Verified,
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
