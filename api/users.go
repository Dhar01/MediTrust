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
	authRepo := repo.NewAuthRepository(cfg.DB)
	verificationRepo := repo.NewVerificationRepository(cfg.DB)

	userService := service.NewUserProfileService(userRepo, cfg.SecretKey)
	authService := service.NewAuthService(authRepo, userRepo, verificationRepo, cfg.SecretKey)
	userCtrl := controllers.NewUserController(userService, authService)

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
