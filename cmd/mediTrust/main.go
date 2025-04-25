package main

import (
	"log"
	"medicine-app/config"
	"medicine-app/internal/api"
	"medicine-app/models"
	"net/http"

	_ "medicine-app/docs"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const apiBase = "/api/v1"

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	defer cfg.DBConn.Close()

	if cfg.Platform != models.Dev {
		gin.SetMode(gin.ReleaseMode)
	}

	router := echo.New()
	router.Use(middleware.Recover())
	router.Use(middleware.Logger())
	router.Use(middleware.Secure())
	router.Use(middleware.RequestID())
	router.Use(middleware.CORS())

	// medicines
	api.MedicineRoutes(router, cfg, apiBase)

	// authentication & authorization
	// srv.AuthRoutes(router, cfg, apiBase)

	// users
	api.AuthUserRoutes(router, cfg, apiBase)

	// public
	api.PublicRoutes(router, cfg, apiBase)

	// admin
	// api.AdminRoutes(router.Group(apiBase), cfg)

	// general routes
	api.GeneralRoutes(router, cfg, apiBase)

	// documentation routes
	// handlers.DocumentationRoute(router.Group(apiBase))

	// cart routes
	// handlers.CartRoute(router.Group(apiBase), cfg)

	port := ":" + cfg.Port

	if err := router.Start(port); err != http.ErrServerClosed {
		log.Fatalf("cant run in port %s: %v", cfg.Port, err)
	}
}
