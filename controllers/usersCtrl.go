package controllers

import "medicine-app/models"

type userController struct {
	UserService models.UserService
}

func NewUserService(service models.UserService) *userController {
	return &userController{
		UserService: service,
	}
}
