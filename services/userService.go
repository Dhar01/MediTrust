package service

import (
	"context"
	"medicine-app/internal/auth"
	"medicine-app/models"
	"time"

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

func (us *userService) SignUpUser(ctx context.Context, user models.SignUpUser) (uuid.UUID, error) {
	pass, err := auth.HashPassword(user.Password)
	if err != nil {
		return uuid.Nil, err
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

	newUser, err := us.Repo.SignUp(ctx, person)
	if err != nil {
		return uuid.Nil, err
	}

	return newUser.ID, nil
}

func (us *userService) LogInUser(ctx context.Context, login models.LogIn) (models.ResponseTokenDTO, error) {
	user, err := us.Repo.FindUser(ctx, models.Email, login.Email)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	if err = auth.CheckPasswordHash(login.Password, user.HashPassword); err != nil {
		return wrapTokenResponseError(err)
	}

	// TODO: need to handle models.Admin type
	accessToken, err := auth.MakeJWT(user.ID, models.Customer, us.Secret, time.Minute*55)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		return wrapTokenResponseError(err)
	}

	if err = us.Repo.CreateRefreshToken(ctx, refreshToken, user.ID); err != nil {
		return wrapTokenResponseError(err)
	}

	return models.ResponseTokenDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
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
			FirstName: updateStrPointerField(user.Name.FirstName, oldInfo.Name.FirstName),
			LastName:  updateStrPointerField(user.Name.LastName, oldInfo.Name.LastName),
		},
		Role:  oldInfo.Role,
		Email: updateField(user.Email, oldInfo.Email),
		Phone: updateField(user.Phone, oldInfo.Phone),
		Age:   *updateIntPointerField(&user.Age, &oldInfo.Age),

		// TODO: need to work on address section

		// Address: models.Address{
		// 	Country:       updateStrPointerField(user.Address.Country, oldInfo.Address.Country),
		// 	City:          updateStrPointerField(user.Address.City, oldInfo.Address.City),
		// 	PostalCode:    updateStrPointerField(user.Address.PostalCode, oldInfo.Address.PostalCode),
		// 	StreetAddress: updateStrPointerField(user.Address.StreetAddress, oldInfo.Address.StreetAddress),
		// },

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

func updateStrPointerField(newValue, oldValue string) string {
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

func wrapTokenResponseError(err error) (models.ResponseTokenDTO, error) {
	return models.ResponseTokenDTO{}, err
}

func wrapUserError(err error) (models.UserResponseDTO, error) {
	return models.UserResponseDTO{}, err
}
