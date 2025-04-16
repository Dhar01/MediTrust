package main

import (
	"log"
	"medicine-app/internal/handlers"
	"medicine-app/config"
	"medicine-app/models"
	srv "medicine-app/internal/services"

	_ "medicine-app/docs"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	router.SetTrustedProxies(nil)

	// medicines
	srv.MedicineRoutes(router.Group(apiBase), cfg)

	// users
	handlers.UserRoutes(router.Group(apiBase), cfg)

	// admin
	// api.AdminRoutes(router.Group(apiBase), cfg)

	// general routes
	handlers.GeneralRoutes(router.Group(apiBase), cfg)

	// documentation routes
	handlers.DocumentationRoute(router.Group(apiBase))

	// cart routes
	handlers.CartRoute(router.Group(apiBase), cfg)

	port := ":" + cfg.Port

	if err := router.Run(port); err != nil {
		log.Fatalf("can't run in port %s: %v", cfg.Port, err)
	}
}
