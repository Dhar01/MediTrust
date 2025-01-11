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
	person, err := us.DB.CreateUser(ctx, database.CreateUserParams{
		FirstName: user.Name.FirstName,
		LastName:  user.Name.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Age:       user.Age,
	})

	if err != nil {
		return models.User{}, err
	}

	return toUserDomain(person), nil
}

func (us *userService) FindUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	user, err := us.DB.GetUserByID(ctx, userID)
	if err != nil {
		return models.User{}, err
	}

	return toUserDomain(user), nil
}

func (us *userService) FindUserByPhone(ctx context.Context, phone string) (models.User, error) {
	user, err := us.DB.GetUserByPhone(ctx, phone)
	if err != nil {
		return models.User{}, err
	}

	return toUserDomain(user), nil
}

func (us *userService) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	user, err := us.DB.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	return toUserDomain(user), nil
}

func (us *userService) UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) (models.User, error) {
	person, err := us.DB.UpdateUser(ctx, database.UpdateUserParams{
		FirstName: user.Name.FirstName,
		LastName: user.Name.LastName,
		Email: user.Email,
		Age: user.Age,
		Phone: user.Phone,
		ID: userID,
	})

	if err != nil {
		return models.User{}, err
	}

	return toUserDomain(person), nil
}

func (us *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if err := us.DB.DeleteUser(ctx, userID); err != nil {
		return err
	}

	return nil
}

func toUserDomain(dbUser database.User) models.User {
	return models.User{
		ID: dbUser.ID,
		Name: models.Name{
			FirstName: dbUser.FirstName,
			LastName:  dbUser.LastName,
		},
		Email:     dbUser.Email,
		Age:       dbUser.Age,
		Phone:     dbUser.Phone,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}
