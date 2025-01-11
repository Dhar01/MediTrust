package service

import (
	"context"
	"medicine-app/internal/database"
	"medicine-app/models"

	"github.com/google/uuid"
)

type userService struct {
	DB *database.Queries
}

func NewUserService(db *database.Queries) models.UserService {
	if db == nil {
		panic("database can't be nil")
	}

	return &userService{
		DB: db,
	}
}

func (us *userService) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (us *userService) FindUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	return models.User{}, nil
}

func (us *userService) UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (us *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return nil
}
