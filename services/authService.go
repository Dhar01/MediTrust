package service

import (
	"context"
	"fmt"
	"medicine-app/internal/auth"
	"medicine-app/models"
	"medicine-app/utils"
	"time"

	"github.com/google/uuid"
)

type authService struct {
	AuthRepo         models.AuthRepository
	UserRepo         models.UserRepository
	VerificationRepo models.VerificationRepository
	Secret           string
	Domain           string
	Port             string
	emailSender      *utils.EmailSender
}

func NewAuthService(
	authRepo models.AuthRepository,
	userRepo models.UserRepository,
	verificationRepo models.VerificationRepository,
	secret, domain, port string,
	emailSender *utils.EmailSender,
) models.Authservice {
	if authRepo == nil || userRepo == nil || verificationRepo == nil {
		panic("repo can't be nil")
	}

	return &authService{
		AuthRepo:         authRepo,
		UserRepo:         userRepo,
		VerificationRepo: verificationRepo,
		Secret:           secret,
		Domain:           domain,
		Port:             port,
		emailSender:      emailSender,
	}
}

func (as *authService) SignUpUser(ctx context.Context, user models.SignUpUser) (models.SignUpResponse, error) {
	// check if the user exists with the associated email
	available, _ := as.UserRepo.FindUser(ctx, models.Email, user.Email)
	if available.Exist {
		return wrapSignUpError(errUserExist)
	}

	// * first user will always be admin
	count, err := as.UserRepo.CountAvailableUsers(ctx)
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

	newUser, err := as.UserRepo.SignUp(ctx, person)
	if err != nil {
		return wrapSignUpError(err)
	}

	verifyToken, err := auth.GenerateVerificationToken(newUser.ID, newUser.Role, as.Secret)
	if err != nil {
		return wrapSignUpError(err)
	}

	// TODO: getting slow response for this.

	emailOpts := utils.EmailOptions{
		To:           newUser.Email,
		Verification: true,
		FirstName:    newUser.Name.FirstName,
		Token:        verifyToken,
		Domain:       as.Domain,
		DomainPort:   as.Port,
	}

	if err := as.emailSender.SendEmail(emailOpts); err != nil {
		return wrapSignUpError(err)
	}

	return models.SignUpResponse{
		ID: newUser.ID,
	}, nil
}

func (as *authService) LogInUser(ctx context.Context, login models.LogIn) (models.TokenResponseDTO, error) {
	user, err := as.UserRepo.FindUser(ctx, models.Email, login.Email)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	ok, err := as.VerificationRepo.GetVerification(ctx, user.ID)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	if !ok {
		return wrapTokenResponseError(fmt.Errorf("not verified"))
	}

	if err = auth.CheckPasswordHash(login.Password, user.HashPassword); err != nil {
		return wrapTokenResponseError(err)
	}

	role, err := as.VerificationRepo.GetUserRole(ctx, user.ID)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	accessToken, err := auth.GenerateAccessToken(user.ID, role, as.Secret, time.Minute*15)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		return wrapTokenResponseError(err)
	}

	if err = as.AuthRepo.CreateRefreshToken(ctx, refreshToken, user.ID); err != nil {
		return wrapTokenResponseError(err)
	}

	return models.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (as *authService) LogoutUser(ctx context.Context, id uuid.UUID) error {
	return as.AuthRepo.Logout(ctx, id)
}

func (as *authService) SetVerifiedUser(ctx context.Context, token string) error {
	id, err := auth.ValidateVerificationToken(token, as.Secret)
	if err != nil {
		return err
	}

	if err = as.VerificationRepo.SetVerification(ctx, id); err != nil {
		return err
	}

	return nil
}

func (as *authService) ResetPassEmail(ctx context.Context, email string) error {
	user, err := as.UserRepo.FindUser(ctx, models.Email, email)
	if err != nil {
		return err
	}

	token, err := auth.GenerateVerificationToken(user.ID, user.Role, as.Secret)
	if err != nil {
		return err
	}

	emailOpts := utils.EmailOptions{
		To:            user.Email,
		ResetPassword: true,
		FirstName:     user.Name.FirstName,
		Token:         token,
		Domain:        as.Domain,
		DomainPort:    as.Port,
	}

	if err := as.emailSender.SendEmail(emailOpts); err != nil {
		return err
	}

	return nil
}

func (as *authService) ResetPassword(ctx context.Context, token, password string) error {
	id, err := auth.ValidateVerificationToken(token, as.Secret)
	if err != nil {
		return err
	}

	pass, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	if err := as.UserRepo.UpdatePassword(ctx, pass, id); err != nil {
		return err
	}

	return nil
}