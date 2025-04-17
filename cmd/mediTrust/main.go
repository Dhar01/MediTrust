package main

import (
	"log"
	"medicine-app/config"
	srv "medicine-app/internal/services"
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

	// router := gin.Default()
	router := echo.New()
	router.Logger.SetLevel(0)
	router.Use(middleware.Logger())
	// router.SetTrustedProxies(nil)

	// medicines
	// srv.MedicineRoutes(router.Group(apiBase), cfg)
	srv.MedicineRoutes(router, cfg, apiBase)

	// authentication & authorization
	srv.AuthRoutes(router, cfg, apiBase)

	// users
	srv.UserRoutes(router, cfg, apiBase)
	// handlers.UserRoutes(router.Group(apiBase), cfg)

	// admin
	// api.AdminRoutes(router.Group(apiBase), cfg)

	// general routes
	// handlers.GeneralRoutes(router.Group(apiBase), cfg)
	srv.GeneralRoutes(router, cfg, apiBase)

	// documentation routes
	// handlers.DocumentationRoute(router.Group(apiBase))

	// cart routes
	// handlers.CartRoute(router.Group(apiBase), cfg)

	port := ":" + cfg.Port

	// if err := router.Run(port); err != nil {
	// 	log.Fatalf("can't run in port %s: %v", cfg.Port, err)
	// }

	if err := router.Start(port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
