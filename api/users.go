package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	"medicine-app/middleware"
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
	router.GET(userBaseByID, middleware.AdminAuth(cfg.SecretKey), userCtrl.HandlerGetUserByID)

	// POST route for users
	router.POST(usersBase, userCtrl.HandlerSignUp)
	router.POST("/signup", userCtrl.HandlerSignUp)
	router.POST("/login", userCtrl.HandlerLogIn)

	userLoggedIn := router.Group(userBaseByID, middleware.IsLoggedIn(cfg.SecretKey))
	{
		// PUT route for users
		userLoggedIn.PUT(userBaseByID, userCtrl.HandlerUpdateUser)

		// DELETE route for users
		userLoggedIn.DELETE(userBaseByID, userCtrl.HandlerDeleteUser)
	}
}
