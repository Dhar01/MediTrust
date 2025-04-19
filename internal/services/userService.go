package services

import (
	"context"
	"errors"
	"medicine-app/internal/errs"
	"medicine-app/internal/repository"
	"medicine-app/models"

	"github.com/google/uuid"
)

type UserService interface {
	FetchUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
	UpdateUserInfoByID(ctx context.Context, id uuid.UUID, updateInfo models.UpdateUserRequest) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	if repo == nil {
		panic("user repository can't be empty/nil")
	}

	return &userService{
		userRepo: repo,
	}
}

func (srv *userService) FetchUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := srv.userRepo.FetchUserByID(ctx, id)
	if err != nil {
		return wrapUserErr(err)
	}

	return user, nil
}

func (srv *userService) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	return wrapUserSpecErr(srv.userRepo.DeleteUserByID(ctx, id))
}

func (srv *userService) UpdateUserInfoByID(ctx context.Context, id uuid.UUID, updateInfo models.UpdateUserRequest) (*models.User, error) {
	oldInfo, err := srv.userRepo.FetchUserByID(ctx, id)
	if err != nil {
		return wrapUserErr(err)
	}

	user, err := srv.userRepo.UpdateUserInfo(ctx, models.User{
		Name: models.FullName{
			FirstName: updateField(updateInfo.Name.FirstName, oldInfo.Name.FirstName),
			LastName:  updateField(updateInfo.Name.LastName, oldInfo.Name.LastName),
		},
		Age:   *updateIntPointerField(&updateInfo.Age, &oldInfo.Age),
		Phone: updateField(updateInfo.Phone, oldInfo.Phone),
		Email: updateField(updateInfo.Email, oldInfo.Email),
	})
	if err != nil {
		return wrapUserErr(err)
	}

	return user, nil
}

func wrapUserErr(err error) (*models.User, error) {
	return nil, wrapUserSpecErr(err)
}

func wrapUserSpecErr(err error) error {
	if errors.Is(err, errs.ErrNotFound) {
		return errs.ErrUserNotExist
	}

	return err
}
