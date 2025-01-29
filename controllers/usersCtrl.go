package controllers

import (
	"fmt"
	"medicine-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userController struct {
	UserService models.UserService
}

func NewUserController(service models.UserService) *userController {
	return &userController{
		UserService: service,
	}
}

func (uc *userController) HandlerSignUp(ctx *gin.Context) {
	var newUser models.SignUpUser

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	id, err := uc.UserService.SignUpUser(ctx.Request.Context(), newUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, id)
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

func (uc *userController) HandlerLogIn(ctx *gin.Context) {
	var login models.LogIn

	if err := ctx.ShouldBindJSON(&login); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	token, err := uc.UserService.LogInUser(ctx.Request.Context(), login)
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

	if err := uc.UserService.LogoutUser(ctx.Request.Context(), id); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.SetCookie(models.TokenRefresh, models.TokenNull, -1, models.RootPath, models.DomainName, true, true)
	ctx.Status(http.StatusOK)
}

func (uc *userController) HandlerGetUserByID(ctx *gin.Context) {
	id, ok := getUserID(ctx)
	if !ok {
		return
	}

	user, err := uc.UserService.FindUserByID(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusFound, user)
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
