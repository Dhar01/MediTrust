package repository

import (
	"context"
	"database/sql"
	"medicine-app/internal/database"
	"medicine-app/models"

	"github.com/google/uuid"
)

type userRepository struct {
	DB *database.Queries
}

func NewUserRepository(db *database.Queries) models.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	person, err := ur.DB.CreateUser(ctx, database.CreateUserParams{
		FirstName: user.Name.FirstName,
		LastName:  user.Name.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Age:       user.Age,
	})

	if err != nil {
		return models.User{}, err
	}

	postalCode := sql.NullString{
		String: user.Address.PostalCode,
		Valid:  true,
	}

	addr, err := ur.DB.CreateUserAddress(ctx, database.CreateUserAddressParams{
		UserID:        person.ID,
		Country:       user.Address.Country,
		City:          user.Address.City,
		StreetAddress: user.Address.StreetAddress,
		PostalCode:    postalCode,
	})

	if err != nil {
		return models.User{}, err
	}

	return toUserDomain(person, addr), nil
}

func (ur *userRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	if err := ur.DB.DeleteUser(ctx, userID); err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Update(ctx context.Context, user models.User) (models.User, error) {
	person, err := ur.DB.UpdateUser(ctx, database.UpdateUserParams{
		FirstName: user.Name.FirstName,
		LastName:  user.Name.LastName,
		Email:     user.Email,
		Age:       user.Age,
		Phone:     user.Phone,
		ID:        user.ID,
	})

	if err != nil {
		return models.User{}, err
	}

	address, err := ur.DB.UpdateAddress(ctx, database.UpdateAddressParams{
		Country:       user.Address.Country,
		City:          user.Address.City,
		StreetAddress: user.Address.StreetAddress,
		PostalCode: sql.NullString{
			String: user.Address.PostalCode,
			Valid:  true,
		},
	})

	if err != nil {
		return models.User{}, err
	}

	return toUserDomain(person, address), err
}

func (ur *userRepository) FindByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	user, err := ur.DB.GetUserByID(ctx, userID)
	if err != nil {
		return models.User{}, err
	}

	address, err := ur.DB.GetAddress(ctx, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return toUserDomain(user, address), nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (models.User, error) {
	user, err := ur.DB.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	address, err := ur.DB.GetAddress(ctx, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return toUserDomain(user, address), nil
}

func (ur *userRepository) FindByPhone(ctx context.Context, phone string) (models.User, error) {
	user, err := ur.DB.GetUserByPhone(ctx, phone)
	if err != nil {
		return models.User{}, err
	}

	address, err := ur.DB.GetAddress(ctx, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return toUserDomain(user, address), nil
}

func toUserDomain(dbUser database.User, address database.UserAddress) models.User {
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
		Address: models.Address{
			Country:       address.Country,
			City:          address.City,
			StreetAddress: address.StreetAddress,
			PostalCode:    address.PostalCode.String,
		},
	}
}
