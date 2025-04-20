package services

import (
	"context"
	"errors"
	"fmt"
	"medicine-app/config"
	"medicine-app/internal/auth"
	"medicine-app/internal/errs"
	"medicine-app/internal/repository"
	"medicine-app/models"
	"medicine-app/utils"
	"time"

	"github.com/google/uuid"
)

type PublicService interface {
	FetchUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	SignUp(ctx context.Context, user models.User) (uuid.UUID, error)
	VerifyUser(ctx context.Context, token string) error
	LogInUser(ctx context.Context, email, password string) (*models.LogInResp, error)
}

type publicService struct {
	publicRepo repository.PublicRepository
	cfg        *config.Config
}

func NewPublicService(
	repo repository.PublicRepository,
	cfg *config.Config,
) PublicService {
	if repo == nil {
		panic("user repository can't be empty/nil")
	}

	return &publicService{
		publicRepo: repo,
		cfg:        cfg,
	}
}

func newUser(user models.User, role string) (*models.User, error) {
	hashedPass, err := auth.HashPassword(user.HashPassword)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Name: models.FullName{
			FirstName: user.Name.FirstName,
			LastName:  user.Name.LastName,
		},
		Age:          user.Age,
		Phone:        user.Phone,
		Email:        user.Email,
		Role:         role,
		IsActive:     false,
		HashPassword: hashedPass,
	}, nil
}

func (srv *publicService) SignUp(ctx context.Context, user models.User) (uuid.UUID, error) {
	exist, err := srv.publicRepo.GetUserByEmail(ctx, user.Email)
	if err == nil && exist.IsActive {
		return uuid.Nil, errs.ErrEmailAlreadyExists
	}

	// * first user will be admin
	count, err := srv.publicRepo.CountUsers(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	var role string
	if count > 0 {
		role = models.Customer
	} else {
		role = models.Admin
	}

	person, err := newUser(user, role)
	if err != nil {
		return uuid.Nil, wrapUserSpecErr(err)
	}

	userInfo, err := srv.publicRepo.CreateUser(ctx, *person)
	if err != nil {
		return uuid.Nil, wrapUserSpecErr(err)
	}

	// ! need to handle error properly
	verificationToken, err := auth.GenerateVerificationToken(userInfo.Id, userInfo.Role, srv.cfg.SecretKey)
	if err != nil {
		return uuid.Nil, wrapUserSpecErr(err)
	}

	// ! need to work on email verification
	emailOpts := utils.EmailOptions{
		To:           userInfo.Email,
		Verification: true,
		FirstName:    userInfo.Name.FirstName,
		Token:        verificationToken,
		Domain:       srv.cfg.Domain,
		DomainPort:   srv.cfg.Port,
	}

	// ! need to handle error properly
	if err := srv.cfg.EmailSender.SendEmail(emailOpts); err != nil {
		return uuid.Nil, wrapUserSpecErr(err)
	}

	return userInfo.Id, nil
}

func (srv *publicService) VerifyUser(ctx context.Context, token string) error {
	id, err := auth.ValidateVerificationToken(token, srv.cfg.SecretKey)
	if err != nil {
		return wrapUserSpecErr(err)
	}

	if err = srv.publicRepo.SetActive(ctx, id); err != nil {
		return wrapUserSpecErr(errs.ErrUserNotExist)
	}

	// user, err := srv.publicRepo.FetchUserByID(ctx, id)
	// if err != nil {
	// 	return wrapUserSpecErr(errs.ErrUserNotAuthorized)
	// }

	// if user.Id != id {
	// 	return wrapUserSpecErr(errs.ErrUserNotAuthorized)
	// }

	return nil
}

func (srv *publicService) LogInUser(ctx context.Context, email, password string) (*models.LogInResp, error) {
	user, err := srv.publicRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, wrapUserSpecErr(errs.ErrUserNotExist)
	}

	if !user.IsActive {
		return nil, wrapUserSpecErr(errs.ErrUserInactive)
	}

	if err = auth.CheckPasswordHash(password, user.HashPassword); err != nil {
		return nil, wrapUserSpecErr(errs.ErrUserNotAuthorized)
	}

	accessToken, err := auth.GenerateAccessToken(user.Id, user.Role, srv.cfg.SecretKey, time.Minute*30)
	if err != nil {
		return nil, wrapUserSpecErr(errs.ErrInternalServer)
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		return nil, wrapUserSpecErr(errs.ErrInternalServer)
	}

	if err := srv.publicRepo.SaveRefreshToken(ctx, user.Id, refreshToken); err != nil {
		return nil, wrapUserSpecErr(errs.ErrInternalServer)
	}

	return &models.LogInResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (srv *publicService) FetchUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := srv.publicRepo.FetchUserByID(ctx, id)
	if err != nil {
		return wrapUserErr(err)
	}

	return user, nil
}

func wrapUserErr(err error) (*models.User, error) {
	return nil, wrapUserSpecErr(err)
}

func wrapUserSpecErr(err error) error {
	fmt.Println("! Error:", err)
	err = errors.Unwrap(err)
	if errors.Is(err, errs.ErrNotFound) {
		return errs.ErrUserNotExist
	}

	return err
}
