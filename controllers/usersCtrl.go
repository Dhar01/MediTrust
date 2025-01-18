package controllers

import (
	"medicine-app/models"
	"net/http"

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

func (uc *userController) HandlerCreateUser(ctx *gin.Context) {
	var newUser models.CreateUserDTO

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	user, err := uc.UserService.CreateUser(ctx, newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (uc *userController) HandlerLogIn(ctx *gin.Context) {
	var login models.LogIn

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	if err := uc.UserService.LogInUser(ctx, login); err != nil {
		ctx.JSON(http.StatusNotFound, errorMsg(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (uc *userController) HandlerGetUserByID(ctx *gin.Context) {
	id, ok := getUserID(ctx)
	if !ok {
		return
	}

	user, err := uc.UserService.FindUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusFound, user)
}

func (uc *userController) HandlerUpdateUser(ctx *gin.Context) {
	id, ok := getUserID(ctx)
	if !ok {
		return
	}

	var updateUser models.UpdateUserDTO

	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	user, err := uc.UserService.UpdateUser(ctx, id, updateUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusAccepted, user)
}

func (uc *userController) HandlerDeleteUser(ctx *gin.Context) {
	id, ok := getUserID(ctx)
	if !ok {
		return
	}

	if err := uc.UserService.DeleteUser(ctx, id); err != nil {
		ctx.JSON(http.StatusNotFound, errorMsg(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (uc *userController) HandlerSignUp(ctx *gin.Context) {
	var newUser models.SignUpUser

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	if err := uc.UserService.SignUpUser(ctx, newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.Status(http.StatusCreated)
}

func getUserID(ctx *gin.Context) (uuid.UUID, bool) {
	userID := ctx.Param("userID")
	id, err := uuid.Parse(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return uuid.Nil, false
	}

	return id, true
}
