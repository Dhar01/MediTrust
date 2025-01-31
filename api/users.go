package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	"medicine-app/middleware"
	repo "medicine-app/repository"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

var usersBase = "/users"

func UserRoutes(router *gin.RouterGroup, cfg *config.Config) {
	userRepo := repo.NewUserRepository(cfg.DB)
	userService := service.NewUserService(userRepo, cfg.SecretKey)
	userCtrl := controllers.NewUserController(userService)

	// GET route for users
	router.GET(usersBase+"/:userID", middleware.AdminAuth(cfg.SecretKey), userCtrl.HandlerGetUserByID)
	router.GET("/verify", userCtrl.HandlerVerify)

	// POST route for users
	router.POST(usersBase, userCtrl.HandlerSignUp)
	router.POST("/signup", userCtrl.HandlerSignUp)
	router.POST("/login", userCtrl.HandlerLogIn)

	// PUT route for users
	router.PUT(usersBase, middleware.IsLoggedIn(cfg.SecretKey), userCtrl.HandlerUpdateUser)

	// DELETE route for users
	router.DELETE(usersBase, middleware.IsLoggedIn(cfg.SecretKey), userCtrl.HandlerDeleteUser)

	router.POST("/logout", middleware.IsLoggedIn(cfg.SecretKey), userCtrl.HandlerLogout)
}
