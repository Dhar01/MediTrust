package api

import (
	"medicine-app/config"
	"medicine-app/controllers"
	"medicine-app/middleware"
	service "medicine-app/services"

	"github.com/gin-gonic/gin"
)

var usersBase = "/users"

func UserRoutes(router *gin.RouterGroup, cfg *config.Config) {
	userService := service.NewUserProfileService(cfg.DB, cfg.SecretKey)
	authService := service.NewAuthService(cfg.SecretKey, cfg.Domain, cfg.Port, cfg.EmailSender, cfg.DB)
	userCtrl := controllers.NewUserController(userService, authService, cfg.Domain)

	// GET route for users
	router.GET(usersBase+"/:userID", middleware.AdminAuth(cfg.SecretKey), userCtrl.HandlerGetUserByID)
	router.GET("/verify", userCtrl.HandlerVerify)

	// POST route for users
	router.POST(usersBase, userCtrl.HandlerSignUp)
	router.POST("/signup", userCtrl.HandlerSignUp)
	router.POST("/login", userCtrl.HandlerLogIn)
	router.POST(usersBase+"/reset", userCtrl.HandlerRequestPasswordReset)
	router.PUT(usersBase+"/reset", userCtrl.HandlerResetUpdatePass)

	// PUT route for users
	router.PUT(usersBase, middleware.IsLoggedIn(cfg.SecretKey), userCtrl.HandlerUpdateUser)

	// DELETE route for users
	router.DELETE(usersBase, middleware.IsLoggedIn(cfg.SecretKey), userCtrl.HandlerDeleteUser)

	router.POST("/logout", middleware.IsLoggedIn(cfg.SecretKey), userCtrl.HandlerLogout)
}
