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

func (uc *userController) HandlerSignUp(ctx *gin.Context) {
	var newUser models.SignUpUser

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	id, err := uc.UserService.SignUpUser(ctx, newUser);
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusCreated, id)
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

func (uc *userController) HandlerLogIn(ctx *gin.Context) {
	var login models.LogIn

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	token, err := uc.UserService.LogInUser(ctx, login)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, token)
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

func getUserID(ctx *gin.Context) (uuid.UUID, bool) {
	userID := ctx.Param("userID")
	id, err := uuid.Parse(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return uuid.Nil, false
	}

	return id, true
}
