package service

import (
	"context"
	"errors"
	"medicine-app/internal/auth"
	"medicine-app/models/db"
	"medicine-app/models/dto"
	repo "medicine-app/repository"

	"github.com/google/uuid"
)

var errUserExist = errors.New("user exist")

type userProfileService struct {
	UserRepo repo.UserRepository
	Secret   string
}


// UserService defines the business logic interface for user management
// @Description Interface for user-related business logic
type UserProfileService interface {
	FindUserByID(ctx context.Context, userID uuid.UUID) (dto.UserResponseDTO, error)
	FindUserByKey(ctx context.Context, key, value string) (dto.UserResponseDTO, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user dto.UpdateUserDTO) (dto.UserResponseDTO, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}


func NewUserProfileService(userRepo repo.UserRepository, secret string) UserProfileService {
	if userRepo == nil {
		panic("repo can't be nil")
	}

	return &userProfileService{
		UserRepo: userRepo,
		Secret:   secret,
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
	oldInfo, err := us.UserRepo.FindByID(ctx, userID)
	// log.Printf("OLDINFO: %+v", oldInfo)
	if err != nil {
		return wrapUserError(err)
	}

	person := db.User{
		ID: userID,
		Name: db.Name{
			FirstName: updateField(user.Name.FirstName, oldInfo.Name.FirstName),
			LastName:  updateField(user.Name.LastName, oldInfo.Name.LastName),
		},
		Role:    oldInfo.Role,
		Email:   updateField(user.Email, oldInfo.Email),
		Phone:   updateField(user.Phone, oldInfo.Phone),
		Age:     *updateIntPointerField(&user.Age, &oldInfo.Age),
		Address: setAddress(user.Address, &oldInfo.Address),
	}

	// log.Printf("UPDATEDUSER: %+v", person)

	userUpdate, err := us.UserRepo.Update(ctx, person)
	// log.Printf("Userupdate Data: %v", userUpdate)
	if err != nil {
		return wrapUserError(err)
	}

	return toUserDTODomain(userUpdate), nil
}

func (us *userProfileService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return us.UserRepo.Delete(ctx, userID)
}

func (us *userProfileService) FindUserByID(ctx context.Context, userID uuid.UUID) (dto.UserResponseDTO, error) {
	user, err := us.UserRepo.FindByID(ctx, userID)
	if err != nil {
		return wrapUserError(err)
	}

	return toUserDTODomain(user), nil
}

// FindUser by KEY. Key should be either Email or Phone.
func (us *userProfileService) FindUserByKey(ctx context.Context, key, value string) (dto.UserResponseDTO, error) {
	person, err := us.UserRepo.FindUser(ctx, key, value)
	if err != nil {
		return wrapUserError(err)
	}

	user, err := us.UserRepo.FindByID(ctx, person.ID)
	if err != nil {
		return wrapUserError(err)
	}

	return toUserDTODomain(user), nil
}

func toUserDTODomain(user db.User) dto.UserResponseDTO {
	return dto.UserResponseDTO{
		ID: user.ID,
		Name: db.Name{
			FirstName: user.Name.FirstName,
			LastName:  user.Name.LastName,
		},
		Email:     user.Email,
		Age:       user.Age,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Address: db.Address{
			Country:       user.Address.Country,
			City:          user.Address.City,
			StreetAddress: user.Address.StreetAddress,
			PostalCode:    user.Address.PostalCode,
		},
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
