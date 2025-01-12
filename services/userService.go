package service

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models"

	"github.com/google/uuid"
)

type userService struct {
	Repo models.UserRepository
}

func NewUserService(repo models.UserRepository) models.UserService {
	if repo == nil {
		panic("repo can't be nil")
	}

	return &userService{
		Repo: repo,
	}
}

func (us *userService) CreateUser(ctx context.Context, user models.CreateUserDTO) (models.User, error) {
	return us.Repo.Create(ctx, models.User{})
}

func (us *userService) FindUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	return us.Repo.FindByID(ctx, userID)
}

func (us *userService) FindUserByPhone(ctx context.Context, phone string) (models.User, error) {
	return us.Repo.FindByPhone(ctx, phone)
}

func (us *userService) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	return us.Repo.FindByEmail(ctx, email)
}

func (us *userService) UpdateUser(ctx context.Context, userID uuid.UUID, user models.UpdateUserDTO) (models.User, error) {
	return us.Repo.Update(ctx, models.User{})
}

func (us *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return us.Repo.Delete(ctx, userID)
}
