package service

import (
	"context"
	"errors"
	"medicine-app/internal/auth"
	"medicine-app/models"

	"github.com/google/uuid"
)

type userService struct {
	Repo   models.UserRepository
	Secret string
}

func NewUserService(repo models.UserRepository, secret string) models.UserService {
	if repo == nil {
		panic("repo can't be nil")
	}

	return &userService{
		Repo:   repo,
		Secret: secret,
	}
}

func (us *userService) SignUpUser(ctx context.Context, user models.SignUpUser) error {
	pass, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	person := models.User{
		Name: models.Name{
			FirstName: user.Name.FirstName,
			LastName:  user.Name.LastName,
		},
		Email:        user.Email,
		HashPassword: pass,
		Age:          user.Age,
		Phone:        user.Phone,
	}

	return us.Repo.SignUp(ctx, person)
}

func (us *userService) LogInUser(ctx context.Context, login models.LogIn) error {
	hashPass, err := us.Repo.FindPass(ctx, login.Email)
	if err != nil {
		return err
	}

	if err = auth.CheckPasswordHash(login.Password, hashPass); err != nil {
		return err
	}

	return nil
}

func (us *userService) UpdateUser(ctx context.Context, userID uuid.UUID, user models.UpdateUserDTO) (models.User, error) {
	var emptyUser models.User

	oldInfo, err := us.Repo.FindByID(ctx, userID)
	if err != nil {
		return emptyUser, errors.New("user not found")
	}

	if user.Name == nil {
		user.Name = &oldInfo.Name
	} else {
		if user.Name.FirstName == "" {
			user.Name.FirstName = oldInfo.Name.FirstName
		}

		if user.Name.LastName == "" {
			user.Name.LastName = oldInfo.Name.LastName
		}
	}

	if user.Address == nil {
		user.Address = &oldInfo.Address
	} else {
		if user.Address.Country == "" {
			user.Address.Country = oldInfo.Address.Country
		}

		if user.Address.City == "" {
			user.Address.City = oldInfo.Address.City
		}

		if user.Address.StreetAddress == "" {
			user.Address.StreetAddress = oldInfo.Address.StreetAddress
		}

		if user.Address.PostalCode == "" {
			user.Address.PostalCode = oldInfo.Address.PostalCode
		}
	}

	person := models.User{
		ID: userID,
		Name: models.Name{
			FirstName: user.Name.FirstName,
			LastName:  user.Name.LastName,
		},
		Email: updateField(user.Email, oldInfo.Email),
		Phone: updateField(user.Phone, oldInfo.Phone),
		Age:   *updateIntPointerField(user.Age, &oldInfo.Age),
		Address: models.Address{
			Country:       user.Address.Country,
			City:          user.Address.City,
			PostalCode:    user.Address.PostalCode,
			StreetAddress: user.Address.StreetAddress,
		},
	}

	return us.Repo.Update(ctx, person)
}

func (us *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return us.Repo.Delete(ctx, userID)
}

func (us *userService) FindUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	return us.Repo.FindByID(ctx, userID)
}

func (us *userService) FindUserByKey(ctx context.Context, key, value string) (models.User, error) {
	return us.Repo.FindUser(ctx, key, value)
}

func updateField(newValue, oldValue string) string {
	if newValue == "" {
		return oldValue
	}

	return newValue
}

func updateIntPointerField(newValue, oldValue *int32) *int32 {
	if newValue == nil {
		return oldValue
	}

	return newValue
}
