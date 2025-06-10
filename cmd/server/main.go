package main

import (
	"log"
	"medicine-app/config"
	r "medicine-app/internal/router"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// load the configuration file
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("failed to load configuration: %v", err)
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

	r.SetUpRouter(cfg, router)

	// starting the server
	if err := router.Start(cfg.Server.ServerHost + ":" + cfg.Server.ServerPort); err != http.ErrServerClosed {
		log.Fatalf("cant run in port %s: %v", cfg.Server.ServerPort, err)
	}
}
