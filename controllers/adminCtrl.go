package controllers

import "medicine-app/models"

type adminController struct {
	AdminService models.AdminService
}

func NewAdminController(service models.AdminService) *adminController {
	return &adminController{
		AdminService: service,
	}
}