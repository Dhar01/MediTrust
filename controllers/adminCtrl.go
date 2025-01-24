package controllers

import (
	"medicine-app/models"

	"github.com/gin-gonic/gin"
)

type adminController struct {
	AdminService models.AdminService
}

func NewAdminController(service models.AdminService) *adminController {
	return &adminController{
		AdminService: service,
	}
}

func (ac *adminController) GetCarts(ctx *gin.Context) {
}