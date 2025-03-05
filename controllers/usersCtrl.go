package controllers

import (
	"fmt"
	"medicine-app/models"
	"medicine-app/models/dto"
	service "medicine-app/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userController struct {
	userService service.UserProfileService
	authService service.AuthService
	domain      string
}

func NewUserController(userService service.UserProfileService, authService service.AuthService, domain string) *userController {
	return &userController{
		userService: userService,
		authService: authService,
		domain:      domain,
	}
}

// HandlerSignUp will use to sign up a user
//
//	@Summary		Sign up a user
//	@Description	Register a new user with email and password.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.SignUpUserDTO		true	"User SignUp request"
//	@Success		201		{object}	dto.SignUpResponseDTO	"ID: uuid"
//	@Failure		400		{object}	dto.ErrorResponseDTO	"Bad request received"
//	@Failure		500		{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/signup [post]
func (uc *userController) HandlerSignUp(ctx *gin.Context) {
	var newUser dto.SignUpUserDTO

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	signUp, err := uc.authService.SignUpUser(ctx.Request.Context(), newUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, signUp)
}

// HandlerLogIn authenticates a user and returns an access token
//
//	@Summary		User Login
//	@Description	Authenticate a user with email and password to obtain an access token in the response body, while the refresh token is set as a secure HTTP-only cookie.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LogInDTO			true	"User log in request"
//	@Success		200		{object}	dto.ServerResponseDTO	"access_token: token"
//	@Success		200		{string}	string					"Set-Cookie: refresh_token=<token>; HttpOnly; Secure; Path=/; Domain=<your-domain.com>"
//	@Failure		400		{object}	dto.ErrorResponseDTO	"Bad request received"
//	@Failure		401		{object}	dto.ErrorResponseDTO	"Unauthorized - Invalid credentials"
//	@Failure		500		{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/login [post]
func (uc *userController) HandlerLogIn(ctx *gin.Context) {
	var login dto.LogInDTO

	if err := ctx.ShouldBindJSON(&login); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	token, err := uc.authService.LogInUser(ctx.Request.Context(), login)
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	ctx.SetCookie(models.TokenRefresh, token.RefreshToken, int(time.Hour*7*24), models.RootPath, uc.domain, true, true)
	ctx.JSON(http.StatusOK, dto.ServerResponseDTO{
		AccessToken: token.AccessToken,
	})
}

// HandlerLogout logs out a user and revokes the access token.
//
//	@Summary		User Logout
//	@Description	logs out the authenticated user by invalidating the refresh token. The refresh token is cleared by setting an expired cookie.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string					"Set-Cookie: refresh_token=; HttpOnly; Secure; Path=/; Domain=<your-domain.com>; Max-Age=0"
//	@Failure		401	{object}	dto.ErrorResponseDTO	"Unauthorized - Invalid or expired token"
//	@Failure		500	{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/logout [post]
func (uc *userController) HandlerLogout(ctx *gin.Context) {
	id, ok := getUserID(ctx)
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("can't get user-id"))
		return
	}

	if err := uc.authService.LogoutUser(ctx.Request.Context(), id); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.SetCookie(models.TokenRefresh, models.TokenNull, -1, models.RootPath, uc.domain, true, true)
	ctx.Status(http.StatusOK)
}

// HandlerUpdateUser will updates the information of user, takes partial update.
//
//	@Summary		User information update
//	@Description	updates user information for the logged in user, takes partial information update. request comes through isLoggedIn middleware.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UpdateUserDTO		true	"user update information request"
//	@Success		202		{object}	dto.UserResponseDTO		"user update response data"
//	@Failure		400		{object}	dto.ErrorResponseDTO	"Bad request received"
//	@Failure		401		{object}	dto.ErrorResponseDTO	"Unauthorized - Invalid or expired token"
//	@Failure		500		{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/users [put]
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

	var updateUser dto.UpdateUserDTO

	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// log.Println(updateUser)

	user, err := uc.userService.UpdateUser(ctx.Request.Context(), id, updateUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusAccepted, user)
}

// HandlerDeleteUser will delete user data
//
//	@Summary		Delete user data
//	@Description	delete the logged in user, request comes through isLoggedIn middleware.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		204	{string}	string					"status no content"
//	@Failure		401	{object}	dto.ErrorResponseDTO	"Unauthorized - Invalid or expired token"
//	@Failure		404	{object}	dto.ErrorResponseDTO	"not found error"
//	@Router			/users [delete]
func (uc *userController) HandlerDeleteUser(ctx *gin.Context) {
	id, ok := getUserID(ctx)
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, fmt.Errorf("user_id not found"))
		return
	}

	if err := uc.userService.DeleteUser(ctx.Request.Context(), id); err != nil {
		errorResponse(ctx, http.StatusNotFound, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// HandlerGetUserByID will find a user information by the userID [admin only]
//
//	@Summary		Get user data by ID
//	@Description	to handler a user for admin, this handler will be used.
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		302	{object}	dto.UserResponseDTO		"status found"
//	@Failure		400	{object}	dto.ErrorResponseDTO	"bad request status"
//	@Failure		404	{object}	dto.ErrorResponseDTO	"not found error"
//	@Router			/users [get]
func (uc *userController) HandlerGetUserByID(ctx *gin.Context) {
	userID := ctx.Param("userID")
	id, err := uuid.Parse(userID)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := uc.userService.FindUserByID(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusFound, user)
}

// HandlerVerify will verify and set a user on the database
//
//	@Summary		Verify a user on SignUp
//	@Description	Upon SignUp, a autogenerated verify link will be sent to the user's email and this handler will verify that user.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		202	{string}	string					"status accepted"
//	@Failure		400	{object}	dto.ErrorResponseDTO	"bad request status"
//	@Router			/verify [get]
func (us *userController) HandlerVerify(ctx *gin.Context) {
	token := ctx.DefaultQuery("token", "")

	if err := us.authService.SetVerifiedUser(ctx.Request.Context(), token); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.Status(http.StatusAccepted)
}

// HandlerRequestPasswordReset will receive the request to reset password & send password reset link
//
//	@Summary		Request for password reset
//	@Description	if a user forget his/her password, they will request for password reset. A password reset link will be sent to the account email
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		202	{string}	string					"status accepted"
//	@Failure		400	{object}	dto.ErrorResponseDTO	"bad request sent"
//	@Failure		500	{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/users/reset [post]
func (us *userController) HandlerRequestPasswordReset(ctx *gin.Context) {
	var newResetPass dto.RequestResetPassDTO

	if err := ctx.ShouldBindBodyWithJSON(&newResetPass); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	if err := us.authService.ResetPassEmail(ctx.Request.Context(), newResetPass.Email); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusAccepted)
}

// HandlerRequestUpdatePass will receive the request to update password and will update password on the database
//
//	@Summary		Updating user password
//	@Description	user will submit update password and this handler will update the password on the database
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		202	{string}	string					"status accepted"
//	@Failure		400	{object}	dto.ErrorResponseDTO	"bad request received"
//	@Failure		500	{object}	dto.ErrorResponseDTO	"Internal server error"
//	@Router			/users/reset [put]
func (us *userController) HandlerResetUpdatePass(ctx *gin.Context) {
	token := ctx.DefaultQuery("token", "")

	var password dto.ResetPassDTO

	if err := ctx.ShouldBindBodyWithJSON(&password); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	if err := us.authService.ResetPassword(ctx.Request.Context(), token, password.Password); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
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
