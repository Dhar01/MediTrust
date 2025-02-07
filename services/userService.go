package service

import (
	"context"
	"errors"
	"medicine-app/internal/auth"
	"medicine-app/models"

	"github.com/google/uuid"
)

var errUserExist = errors.New("user exist")

type userProfileService struct {
	UserRepo models.UserRepository
	Secret   string
}

func NewUserProfileService(userRepo models.UserRepository, secret string) models.UserProfileService {
	if userRepo == nil {
		panic("repo can't be nil")
	}

	return &userProfileService{
		UserRepo: userRepo,
		Secret:   secret,
	}
}

// NewUser constructor
func NewUser(signUp models.SignUpUser, role string) (models.User, error) {
	hashedPass, err := auth.HashPassword(signUp.Password)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Name:         signUp.Name,
		Email:        signUp.Email,
		Age:          signUp.Age,
		Phone:        signUp.Phone,
		HashPassword: hashedPass,
		Role:         role,
	}, nil
}

func (us *userProfileService) UpdateUser(ctx context.Context, userID uuid.UUID, user models.UpdateUserDTO) (models.UserResponseDTO, error) {
	oldInfo, err := us.UserRepo.FindByID(ctx, userID)
	// log.Printf("OLDINFO: %+v", oldInfo)
	if err != nil {
		return wrapUserError(err)
	}

	person := models.User{
		ID: userID,
		Name: models.Name{
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

func (us *userProfileService) FindUserByID(ctx context.Context, userID uuid.UUID) (models.UserResponseDTO, error) {
	user, err := us.UserRepo.FindByID(ctx, userID)
	if err != nil {
		return wrapUserError(err)
	}

	return toUserDTODomain(user), nil
}

// FindUser by KEY. Key should be either Email or Phone.
func (us *userProfileService) FindUserByKey(ctx context.Context, key, value string) (models.UserResponseDTO, error) {
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

func toUserDTODomain(user models.User) models.UserResponseDTO {
	return models.UserResponseDTO{
		ID: user.ID,
		Name: models.Name{
			FirstName: user.Name.FirstName,
			LastName:  user.Name.LastName,
		},
		Email:     user.Email,
		Age:       user.Age,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Address: models.Address{
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

func wrapTokenResponseError(err error) (models.TokenResponseDTO, error) {
	return models.TokenResponseDTO{}, err
}

func wrapUserError(err error) (models.UserResponseDTO, error) {
	return models.UserResponseDTO{}, err
}

func wrapSignUpError(err error) (models.SignUpResponse, error) {
	return models.SignUpResponse{}, err
}

func setAddress(address, oldAddress *models.Address) models.Address {
	if address == nil {
		address = oldAddress
	}

	return models.Address{
		Country:       address.Country,
		City:          address.City,
		StreetAddress: address.StreetAddress,
		PostalCode:    address.PostalCode,
	}

}
