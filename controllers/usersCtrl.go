package controllers

import (
	"fmt"
	"log"
	"medicine-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userController struct {
	UserService models.UserProfileService
	AuthService models.Authservice
}

func NewUserController(userService models.UserProfileService, authService models.Authservice) *userController {
	return &userController{
		UserService: userService,
		AuthService: authService,
	}
}

// HandlerSignUp will use to sign up a user
//
//	@Summary		Sign up a user
//	@Description	Register a new user with email and password.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.SignUpUser		true	"User signup request"
//	@Success		201		{object}	models.SignUpResponse	"ID: uuid"
//	@Failure		400		{object}	models.ErrorResponse	"Bad request received"
//	@Failure		500		{object}	models.ErrorResponse	"Internal server error"
//	@Router			/signup [post]
func (uc *userController) HandlerSignUp(ctx *gin.Context) {
	var newUser models.SignUpUser

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	signUp, err := uc.AuthService.SignUpUser(ctx.Request.Context(), newUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, signUp)
}

func (uc *userController) HandlerLogIn(ctx *gin.Context) {
	var login models.LogIn

	if err := ctx.ShouldBindJSON(&login); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	token, err := uc.AuthService.LogInUser(ctx.Request.Context(), login)
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, err)
		return
	}

	ctx.SetCookie(models.TokenRefresh, token.RefreshToken, int(time.Hour*7*24), models.RootPath, models.DomainName, true, true)
	ctx.JSON(http.StatusOK, models.ServerResponse{
		AccessToken: token.AccessToken,
	})
}

func (uc *userController) HandlerLogout(ctx *gin.Context) {
	id, ok := getUserID(ctx)
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("can't get user-id"))
		return
	}

	if err := uc.AuthService.LogoutUser(ctx.Request.Context(), id); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.SetCookie(models.TokenRefresh, models.TokenNull, -1, models.RootPath, models.DomainName, true, true)
	ctx.Status(http.StatusOK)
}

func (uc *userController) HandlerUpdateUser(ctx *gin.Context) {
	id, ok := getUserID(ctx)
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("user_id not found"))
		return
	}

	role, ok := getRole(ctx)
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("role not found"))
		return
	}

	if role != models.Admin && role != models.Customer {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("role not applicable"))
		return
	}

	var updateUser models.UpdateUserDTO

	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// log.Println(updateUser)

	user, err := uc.UserService.UpdateUser(ctx.Request.Context(), id, updateUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusAccepted, user)
}

func (uc *userController) HandlerDeleteUser(ctx *gin.Context) {
	id, ok := getUserID(ctx)
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("user_id not found"))
		return
	}

	if err := uc.UserService.DeleteUser(ctx.Request.Context(), id); err != nil {
		errorResponse(ctx, http.StatusNotFound, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (uc *userController) HandlerGetUserByID(ctx *gin.Context) {

	// id, ok := getUserID(ctx)
	// if !ok {
	// 	return
	// }

	userID := ctx.Param("userID")
	id, err := uuid.Parse(userID)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	log.Println(userID)

	user, err := uc.UserService.FindUserByID(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusFound, user)
}

func (us *userController) HandlerVerify(ctx *gin.Context) {
	token := ctx.DefaultQuery("token", "")

	if err := us.AuthService.SetVerifiedUser(ctx.Request.Context(), token); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (us *userController) HandlerRequestPasswordReset(ctx *gin.Context) {
	var newResetPass models.RequestResetPass

	if err := ctx.ShouldBindBodyWithJSON(&newResetPass); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	if err := us.AuthService.ResetPassEmail(ctx.Request.Context(), newResetPass.Email); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (us *userController) HandlerResetUpdatePass(ctx *gin.Context) {
	token := ctx.DefaultQuery("token", "")

	var password models.ResetPass

	if err := ctx.ShouldBindBodyWithJSON(&password); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	if err := us.AuthService.ResetPassword(ctx.Request.Context(), token, password.Password); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.Status(http.StatusAccepted)
}

func getUserID(ctx *gin.Context) (uuid.UUID, bool) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return uuid.Nil, false
	}

	return userID.(uuid.UUID), true
}

func getRole(ctx *gin.Context) (string, bool) {
	role, exists := ctx.Get("role")
	if !exists {
		return "", false
	}

	return role.(string), true
}
