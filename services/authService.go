package service

import (
	"context"
	"fmt"
	"medicine-app/internal/auth"
	"medicine-app/models"
	"medicine-app/models/dto"
	repo "medicine-app/repository"
	"medicine-app/utils"
	"time"

	"github.com/google/uuid"
)

type authService struct {
	authRepo         repo.AuthRepository
	userRepo         repo.UserRepository
	verificationRepo repo.VerificationRepository
	secret           string
	domain           string
	port             string
	emailSender      *utils.EmailSender
}

// AuthService defines the business logic interface for authentication management
// @Description Interface for authentication-authorization related business logic
type AuthService interface {
	SignUpUser(ctx context.Context, user dto.SignUpUserDTO) (dto.SignUpResponseDTO, error) // should act as CreateUser
	LogInUser(ctx context.Context, login dto.LogInDTO) (dto.TokenResponseDTO, error)
	LogoutUser(ctx context.Context, id uuid.UUID) error
	SetVerifiedUser(ctx context.Context, token string) error
	ResetPassEmail(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, password string) error
}

func NewAuthService(
	authRepo repo.AuthRepository,
	userRepo repo.UserRepository,
	verificationRepo repo.VerificationRepository,
	secret, domain, port string,
	emailSender *utils.EmailSender,
) AuthService {
	if authRepo == nil || userRepo == nil || verificationRepo == nil {
		panic("repo can't be nil")
	}

	return &authService{
		authRepo:         authRepo,
		userRepo:         userRepo,
		verificationRepo: verificationRepo,
		secret:           secret,
		domain:           domain,
		port:             port,
		emailSender:      emailSender,
	}
}

func (as *authService) SignUpUser(ctx context.Context, user dto.SignUpUserDTO) (dto.SignUpResponseDTO, error) {
	// check if the user exists with the associated email
	available, _ := as.userRepo.FindUser(ctx, models.Email, user.Email)
	if available.Exist {
		return wrapSignUpError(errUserExist)
	}

	// * first user will always be admin
	count, err := as.userRepo.CountAvailableUsers(ctx)
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

	newUser, err := as.userRepo.SignUp(ctx, person)
	if err != nil {
		return wrapSignUpError(err)
	}

	verifyToken, err := auth.GenerateVerificationToken(newUser.ID, newUser.Role, as.secret)
	if err != nil {
		return wrapSignUpError(err)
	}

	// TODO: getting slow response for this.

	emailOpts := utils.EmailOptions{
		To:           newUser.Email,
		Verification: true,
		FirstName:    newUser.Name.FirstName,
		Token:        verifyToken,
		Domain:       as.domain,
		DomainPort:   as.port,
	}

	if err := as.emailSender.SendEmail(emailOpts); err != nil {
		return wrapSignUpError(err)
	}

	return dto.SignUpResponseDTO{
		ID: newUser.ID,
	}, nil
}

func (as *authService) LogInUser(ctx context.Context, login dto.LogInDTO) (dto.TokenResponseDTO, error) {
	user, err := as.userRepo.FindUser(ctx, models.Email, login.Email)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	ok, err := as.verificationRepo.GetVerification(ctx, user.ID)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	if !ok {
		return wrapTokenResponseError(fmt.Errorf("not verified"))
	}

	if err = auth.CheckPasswordHash(login.Password, user.HashPassword); err != nil {
		return wrapTokenResponseError(err)
	}

	role, err := as.verificationRepo.GetUserRole(ctx, user.ID)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	accessToken, err := auth.GenerateAccessToken(user.ID, role, as.secret, time.Minute*15)
	if err != nil {
		return wrapTokenResponseError(err)
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		return wrapTokenResponseError(err)
	}

	if err = as.authRepo.CreateRefreshToken(ctx, refreshToken, user.ID); err != nil {
		return wrapTokenResponseError(err)
	}

	return dto.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (as *authService) LogoutUser(ctx context.Context, id uuid.UUID) error {
	return as.authRepo.Logout(ctx, id)
}

func (as *authService) SetVerifiedUser(ctx context.Context, token string) error {
	id, err := auth.ValidateVerificationToken(token, as.secret)
	if err != nil {
		return err
	}

	if err = as.verificationRepo.SetVerification(ctx, id); err != nil {
		return err
	}

	return nil
}

func (as *authService) ResetPassEmail(ctx context.Context, email string) error {
	user, err := as.userRepo.FindUser(ctx, models.Email, email)
	if err != nil {
		return err
	}

	token, err := auth.GenerateVerificationToken(user.ID, user.Role, as.secret)
	if err != nil {
		return err
	}

	emailOpts := utils.EmailOptions{
		To:            user.Email,
		ResetPassword: true,
		FirstName:     user.Name.FirstName,
		Token:         token,
		Domain:        as.domain,
		DomainPort:    as.port,
	}

	if err := as.emailSender.SendEmail(emailOpts); err != nil {
		return err
	}

	return nil
}

func (as *authService) ResetPassword(ctx context.Context, token, password string) error {
	id, err := auth.ValidateVerificationToken(token, as.secret)
	if err != nil {
		return err
	}

	pass, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	if err := as.userRepo.UpdatePassword(ctx, pass, id); err != nil {
		return err
	}

	return nil
}
