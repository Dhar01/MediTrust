package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	repo "medicine-app/repository"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

const (
	usersBase    = "/users"
	userBaseByID = usersBase + "/:userID"
)

func UserRoutes(router *gin.RouterGroup, cfg *config.Config) {
	userRepo := repo.NewUserRepository(cfg.DB)
	userService := service.NewUserService(userRepo, cfg.SecretKey)
	userCtrl := controllers.NewUserController(userService)

	// GET route for users
	router.GET(userBaseByID, userCtrl.HandlerGetUserByID)

	// PUT route for users
	router.PUT(userBaseByID, userCtrl.HandlerUpdateUser)

	// POST route for users
	router.POST(usersBase, userCtrl.HandlerSignUp)
	router.POST("/signup", userCtrl.HandlerSignUp)
	router.POST("/login", userCtrl.HandlerLogIn)

	// DELETE route for users
	router.DELETE(userBaseByID, userCtrl.HandlerDeleteUser)
}
