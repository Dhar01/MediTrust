package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"medicine-app/internal/auth"
	"medicine-app/models"
	"medicine-app/utils"
	"time"

	"github.com/google/uuid"
)

var errUserExist = errors.New("user exist")

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

func (us *userService) SignUpUser(ctx context.Context, user models.SignUpUser) (models.SignUpResponse, error) {
	// check if the user exists with the associated email
	available, _ := us.Repo.FindUser(ctx, models.Email, user.Email)
	if available.Exist {
		return wrapSignUpError(errUserExist)
	}

	// * first user will always be admin
	count, err := us.Repo.CountAvailableUsers(ctx)
	if err != nil {
		return wrapSignUpError(err)
	}
	var role string
	if count > 0 {
		role = models.Customer
	} else {
		role = models.Admin
	}

	/*
		TODO: above piece of code need to be replaced later.
		* For the MVP, I wanted create the first user as Admin.
		TODO: I'm planning to create a cmd to create admin.
	*/

	// TODO: need to work on field validation of users (below)

	person, err := NewUser(user, role)
	if err != nil {
		return wrapSignUpError(err)
	}

	newUser, err := us.Repo.SignUp(ctx, person)
	if err != nil {
		return wrapSignUpError(err)
	}

	verifyToken, err := auth.GenerateVerificationToken(newUser.ID, newUser.Role, us.Secret)
	if err != nil {
		return wrapSignUpError(err)
	}

	// TODO: getting slow response for this.
	sendEmail(newUser, verifyToken)

	return models.SignUpResponse{
		ID: newUser.ID,
	}, nil
}

func sendEmail(user models.User, token string) {
	if err := utils.SendVerificationEmail(user.Email, user.Name.FirstName, models.DomainName, token, models.DomainPort); err != nil {
		log.Println(err)
	}
}

func (us *userService) LogInUser(ctx context.Context, login models.LogIn) (models.TokenResponseDTO, error) {
	user, err := us.Repo.FindUser(ctx, models.Email, login.Email)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	ok, err := us.Repo.GetVerification(ctx, user.ID)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	if !ok {
		return wrapTokenResponseError(fmt.Errorf("not verified"))
	}

	if err = auth.CheckPasswordHash(login.Password, user.HashPassword); err != nil {
		return wrapTokenResponseError(err)
	}

	role, err := us.Repo.GetUserRole(ctx, user.ID)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	accessToken, err := auth.GenerateAccessToken(user.ID, role, us.Secret, time.Minute*15)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		return wrapTokenResponseError(err)
	}

	if err = us.Repo.CreateRefreshToken(ctx, refreshToken, user.ID); err != nil {
		return wrapTokenResponseError(err)
	}

	return models.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (us *userService) SetVerifiedUser(ctx context.Context, token string) error {
	id, err := auth.ValidateVerificationToken(token, us.Secret)
	if err != nil {
		return err
	}

	if err = us.Repo.SetVerification(ctx, id); err != nil {
		return err
	}

	return nil
}

func (us *userService) LogoutUser(ctx context.Context, id uuid.UUID) error {
	return us.Repo.Logout(ctx, id)
}

func (us *userService) UpdateUser(ctx context.Context, userID uuid.UUID, user models.UpdateUserDTO) (models.UserResponseDTO, error) {
	oldInfo, err := us.Repo.FindByID(ctx, userID)
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

	userUpdate, err := us.Repo.Update(ctx, person)
	// log.Printf("Userupdate Data: %v", userUpdate)
	if err != nil {
		return wrapUserError(err)
	}

	return toUserDTODomain(userUpdate), nil
}

func (us *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return us.Repo.Delete(ctx, userID)
}

func (us *userService) FindUserByID(ctx context.Context, userID uuid.UUID) (models.UserResponseDTO, error) {
	user, err := us.Repo.FindByID(ctx, userID)
	if err != nil {
		return wrapUserError(err)
	}

	return toUserDTODomain(user), nil
}

// FindUser by KEY. Key should be either Email or Phone.
func (us *userService) FindUserByKey(ctx context.Context, key, value string) (models.UserResponseDTO, error) {
	person, err := us.Repo.FindUser(ctx, key, value)
	if err != nil {
		return wrapUserError(err)
	}

	user, err := us.Repo.FindByID(ctx, person.ID)
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
