package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"medicine-app/internal/auth"
	"medicine-app/internal/database"
	"medicine-app/models"
	"medicine-app/models/db"
	"medicine-app/models/dto"

	"github.com/google/uuid"
)

var errUserExist = errors.New("user exist")

type userProfileService struct {
	DB     *database.Queries
	secret string
}

// UserService defines the business logic interface for user management
// @Description Interface for user-related business logic
type UserProfileService interface {
	FindUserByID(ctx context.Context, userID uuid.UUID) (dto.UserResponseDTO, error)
	FindUserByKey(ctx context.Context, key, value string) (dto.UserResponseDTO, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user dto.UpdateUserDTO) (dto.UserResponseDTO, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

func NewUserProfileService(db *database.Queries, secret string) UserProfileService {
	if db == nil {
		panic("repo can't be nil")
	}

	return &userProfileService{
		DB:     db,
		secret: secret,
	}
}

// NewUser constructor
func NewUser(signUp dto.SignUpUserDTO, role string) (db.User, error) {
	hashedPass, err := auth.HashPassword(signUp.Password)
	if err != nil {
		return db.User{}, err
	}

	return db.User{
		Name:         signUp.Name,
		Email:        signUp.Email,
		Age:          signUp.Age,
		Phone:        signUp.Phone,
		HashPassword: hashedPass,
		Role:         role,
	}, nil
}

func (us *userProfileService) UpdateUser(ctx context.Context, userID uuid.UUID, user dto.UpdateUserDTO) (dto.UserResponseDTO, error) {
	oldInfo, err := us.DB.GetUserByID(ctx, userID)
	// log.Printf("OldInfo: %+v", oldInfo)
	if err != nil {
		return wrapUserError(err)
	}

	person := db.User{
		ID: userID,
		Name: db.Name{
			FirstName: updateField(user.Name.FirstName, oldInfo.FirstName),
			LastName:  updateField(user.Name.LastName, oldInfo.LastName),
		},
		Role:  oldInfo.Role,
		Email: updateField(user.Email, oldInfo.Email),
		Phone: updateField(user.Phone, oldInfo.Phone),
		Age:   *updateIntPointerField(&user.Age, &oldInfo.Age),
		// Address: setAddress(user.Address, &oldInfo.Address),
	}

	userUpdate, err := us.DB.UpdateUser(ctx, database.UpdateUserParams{
		FirstName: person.Name.FirstName,
		LastName:  person.Name.LastName,
		Email:     person.Email,
		Age:       person.Age,
		Phone:     person.Phone,
		ID:        person.ID,
	})
	if err != nil {
		return wrapUserError(err)
	}

	_, err = us.DB.CheckAddressExist(ctx, person.ID)
	if err != nil {
		return wrapUserError(err)
	}

	// var address database.UserAddress
	// if !exists {
	// 	address, err = us.DB.CreateUserAddress(ctx, database.CreateUserAddressParams{
	// 		UserID:        user.ID,
	// 		Country:       user.Address.Country,
	// 		City:          user.Address.City,
	// 		StreetAddress: user.Address.StreetAddress,
	// 		PostalCode:    toNullString(user.Address.PostalCode),
	// 	})
	// } else {
	// 	address, err = us.DB.UpdateAddress(ctx, database.UpdateAddressParams{
	// 		UserID:        user.ID,
	// 		Country:       user.Address.Country,
	// 		City:          user.Address.City,
	// 		StreetAddress: user.Address.StreetAddress,
	// 		PostalCode:    toNullString(user.Address.PostalCode),
	// 	})
	// }

	return toUserDTODomain(userUpdate), nil
}

func (us *userProfileService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if err := us.DB.DeleteUser(ctx, userID); err != nil {
		return err
	}

	return nil
}

func (us *userProfileService) FindUserByID(ctx context.Context, userID uuid.UUID) (dto.UserResponseDTO, error) {
	user, err := us.DB.GetUserByID(ctx, userID)
	if err != nil {
		return wrapUserError(err)
	}

	return toUserDTODomain(user), nil
}

func (us *userProfileService) FindUserByKey(ctx context.Context, key, value string) (dto.UserResponseDTO, error) {
	var user database.User
	var err error

	switch key {
	case models.Email:
		user, err = us.DB.GetUserByEmail(ctx, value)
	case models.Phone:
		user, err = us.DB.GetUserByPhone(ctx, value)
	default:
		return wrapUserError(fmt.Errorf("unsupported lookup key %s", key))
	}

	if err != nil {
		return wrapUserError(err)
	}

	return toUserDTODomain(user), nil
}

func toUserDTODomain(user database.User) dto.UserResponseDTO {
	return dto.UserResponseDTO{
		ID: user.ID,
		Name: db.Name{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Email:     user.Email,
		Age:       user.Age,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,

		// Address: db.Address{
		// 	Country:       user.,
		// 	City:          user.City,
		// 	StreetAddress: user.StreetAddress,
		// 	PostalCode:    user.PostalCode,
		// },

	}
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

func wrapTokenResponseError(err error) (dto.TokenResponseDTO, error) {
	return dto.TokenResponseDTO{}, err
}

func wrapUserError(err error) (dto.UserResponseDTO, error) {
	return dto.UserResponseDTO{}, err
}

func wrapSignUpError(err error) (dto.SignUpResponseDTO, error) {
	return dto.SignUpResponseDTO{}, err
}

func setAddress(address, oldAddress *db.Address) db.Address {
	if address == nil {
		address = oldAddress
	}

	return db.Address{
		Country:       address.Country,
		City:          address.City,
		StreetAddress: address.StreetAddress,
		PostalCode:    address.PostalCode,
	}

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
