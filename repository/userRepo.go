package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"medicine-app/internal/database"
	"medicine-app/models"
	"medicine-app/models/db"

	"github.com/google/uuid"
)

type userRepository struct {
	DB *database.Queries
}

func NewUserRepository(db *database.Queries) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) SignUp(ctx context.Context, user db.User) (db.User, error) {
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

func (ur *userRepository) Update(ctx context.Context, user db.User) (db.User, error) {
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

func (ur *userRepository) UpdatePassword(ctx context.Context, password string, id uuid.UUID) error {
	if err := ur.DB.ResetPassword(ctx, database.ResetPasswordParams{
		PasswordHash: password,
		ID:           id,
	}); err != nil {
		return err
	}

	return nil
}

// FindUser by KEY. Key should be either Email or Phone.
func (ur *userRepository) FindUser(ctx context.Context, key, value string) (db.User, error) {
	var user database.User
	var err error

	switch key {
	case models.Email:
		user, err = ur.DB.GetUserByEmail(ctx, value)
	case models.Phone:
		user, err = ur.DB.GetUserByPhone(ctx, value)
	default:
		return db.User{}, fmt.Errorf("unsupported lookup key %s", key)
	}

	if err != nil {
		return wrapUserError(err)
	}

	return toUser(user), nil
}

func (ur *userRepository) FindByID(ctx context.Context, userID uuid.UUID) (db.User, error) {
	user, err := ur.DB.GetUserByID(ctx, userID)
	if err != nil {
		return wrapUserError(err)
	}

	// log.Printf("DBUSER: %+v", user)

	return toUser(user), nil
}

func (ur *userRepository) CountAvailableUsers(ctx context.Context) (int, error) {
	count, err := ur.DB.CountUsers(ctx)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func toUserDomain(dbUser database.User, address database.UserAddress) db.User {
	return db.User{
		ID: dbUser.ID,
		Name: db.Name{
			FirstName: dbUser.FirstName,
			LastName:  dbUser.LastName,
		},
		Email:     dbUser.Email,
		Age:       dbUser.Age,
		Phone:     dbUser.Phone,
		Role:      dbUser.Role,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Address: db.Address{
			Country:       address.Country,
			City:          address.City,
			StreetAddress: address.StreetAddress,
			PostalCode:    address.PostalCode.String,
		},
	}
}

func wrapUserError(err error) (db.User, error) {
	return db.User{}, err
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

func toUser(dbUser database.User) db.User {
	return db.User{
		ID: dbUser.ID,
		Name: db.Name{
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
