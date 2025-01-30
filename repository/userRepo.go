package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

func (ur *userRepository) SignUp(ctx context.Context, user models.User) (models.User, error) {
	newUser, err := ur.DB.CreateUser(ctx, database.CreateUserParams{
		FirstName:    user.Name.FirstName,
		LastName:     user.Name.LastName,
		Email:        user.Email,
		Role:         user.Role,
		Age:          user.Age,
		Phone:        user.Phone,
		PasswordHash: user.HashPassword,
	})

	if err != nil {
		return wrapUserError(err)
	}

	return toUser(newUser), nil
}

func (ur *userRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	if err := ur.DB.DeleteUser(ctx, userID); err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Update(ctx context.Context, user models.User) (models.User, error) {
	log.Printf("User ID before update: %v", user.ID)

	person, err := ur.DB.UpdateUser(ctx, database.UpdateUserParams{
		FirstName: user.Name.FirstName,
		LastName:  user.Name.LastName,
		Email:     user.Email,
		Age:       user.Age,
		Phone:     user.Phone,
		ID:        user.ID,
	})

	if err != nil {
		return wrapUserError(err)
	}

	// TODO: need to work on address sector
	log.Printf("updated person ID: %+v", person.ID)

	exists, err := ur.DB.CheckAddressExist(ctx, person.ID)
	if err != nil {
		return wrapUserError(err)
	}

	var address database.UserAddress

	if !exists {
		address, err = ur.DB.CreateUserAddress(ctx, database.CreateUserAddressParams{
			UserID:        user.ID,
			Country:       user.Address.Country,
			City:          user.Address.City,
			StreetAddress: user.Address.StreetAddress,
			PostalCode:    toNullString(user.Address.PostalCode),
		})
	} else {
		address, err = ur.DB.UpdateAddress(ctx, database.UpdateAddressParams{
			UserID:        user.ID,
			Country:       user.Address.Country,
			City:          user.Address.City,
			StreetAddress: user.Address.StreetAddress,
			PostalCode:    toNullString(user.Address.PostalCode),
		})
	}

	log.Printf("address user ID: %+v", address.UserID)
	if err != nil {
		return wrapUserError(err)
	}

	return toUserDomain(person, address), nil

	// return toUser(person), nil
}

// FindUser by KEY. Key should be either Email or Phone.
func (ur *userRepository) FindUser(ctx context.Context, key, value string) (models.User, error) {
	var user database.User
	var err error

	switch key {
	case models.Email:
		user, err = ur.DB.GetUserByEmail(ctx, value)
	case models.Phone:
		user, err = ur.DB.GetUserByPhone(ctx, value)
	default:
		return models.User{}, fmt.Errorf("unsupported lookup key %s", key)
	}

	if err != nil {
		return wrapUserError(err)
	}

	return toUser(user), nil
}

func (ur *userRepository) FindByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	user, err := ur.DB.GetUserByID(ctx, userID)
	if err != nil {
		return wrapUserError(err)
	}

	// log.Printf("DBUSER: %+v", user)

	return toUser(user), nil
}

func (ur *userRepository) CreateRefreshToken(ctx context.Context, token string, id uuid.UUID) error {
	return ur.DB.CreateRefreshToken(ctx, database.CreateRefreshTokenParams{
		Refreshtoken: token,
		UserID:       id,
	})
}

func (ur *userRepository) FindUserFromToken(ctx context.Context, token string) (models.User, error) {
	user, err := ur.DB.GetUserFromRefreshToken(ctx, token)
	if err != nil {
		return wrapUserError(err)
	}

	return ur.userWithAddress(ctx, user)
}

func (ur *userRepository) Logout(ctx context.Context, id uuid.UUID) error {
	if err := ur.DB.RevokeTokenByID(ctx, id); err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) CountAvailableUsers(ctx context.Context) (int, error) {
	count, err := ur.DB.CountUsers(ctx)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (ur *userRepository) GetUserRole(ctx context.Context, id uuid.UUID) (string, error) {
	role, err := ur.DB.GetRole(ctx, id)
	if err != nil {
		return "", err
	}

	return role, nil
}

func (ur *userRepository) GetVerification(ctx context.Context, id uuid.UUID) (bool, error) {
	ok, err := ur.DB.GetVerified(ctx, id)
	if err != nil {
		return ok, err
	}

	return ok, nil
}

func (ur *userRepository) SetVerification(ctx context.Context, id uuid.UUID) error {
	if err := ur.DB.SetVerified(ctx, id); err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) userWithAddress(ctx context.Context, user database.User) (models.User, error) {
	address, err := ur.DB.GetAddress(ctx, user.ID)
	if err != nil {
		return wrapUserError(err)
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
		Role:      dbUser.Role,
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

func wrapUserError(err error) (models.User, error) {
	return models.User{}, err
}

func toNullString(value string) sql.NullString {
	if value == "" {
		return sql.NullString{
			Valid: false,
		}
	}
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

func toUser(dbUser database.User) models.User {
	return models.User{
		ID: dbUser.ID,
		Name: models.Name{
			FirstName: dbUser.FirstName,
			LastName:  dbUser.LastName,
		},
		Email:        dbUser.Email,
		Exist:        true,
		Age:          dbUser.Age,
		HashPassword: dbUser.PasswordHash,
		Phone:        dbUser.Phone,
		Role:         dbUser.Role,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
	}
}
