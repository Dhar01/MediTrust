package main

import (
	"log"
	"medicine-app/config"
	"medicine-app/internal/product"
	"net/http"

	_ "medicine-app/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const apiBase string = "/api/v1"

func main() {
	// load the configuration file
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// get the configuration
	cfg := config.GetConfig()

	// defining defaults
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.Secure())
	router.Use(middleware.RequestID())
	router.Use(middleware.CORS())

	// products
	product.ProductRoutes(*router.Group(apiBase), *cfg)

	// medicines
	// product.MedicineRoutes(router, cfg, apiBase)

	// authentication & authorization
	// srv.AuthRoutes(router, cfg, apiBase)

	// users
	// api.AuthUserRoutes(router, cfg, apiBase)

	// public
	// api.PublicRoutes(router, cfg, apiBase)

	// admin
	// api.AdminRoutes(router.Group(apiBase), cfg)

	// general routes
	// api.GeneralRoutes(router, cfg, apiBase)

	// documentation routes
	// handlers.DocumentationRoute(router.Group(apiBase))

	// cart routes
	// handlers.CartRoute(router.Group(apiBase), cfg)

	// starting the server
	if err := router.Start(cfg.Server.ServerHost + ":" + cfg.Server.ServerPort); err != http.ErrServerClosed {
		log.Fatalf("cant run in port %s: %v", cfg.Server.ServerPort, err)
	}
}
