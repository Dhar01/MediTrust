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
	var newUser models.User

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

func (uc *userController) HandlerDeleteUser(ctx *gin.Context) {
	id, ok := getUserID(ctx)
	if !ok {
		return
	}

	if err := uc.UserService.DeleteUser(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.Status(http.StatusNoContent)
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
