package main

import (
	"log"
	"medicine-app/api"
	"medicine-app/config"

	"github.com/gin-gonic/gin"
	"github.com/go-swagno/swagno"
	"github.com/go-swagno/swagno/components/endpoint"
)

const apiBase = "/api/v1"

func main() {
	cfg := config.NewConfig()

	// uncomment this line for production
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Serving Swagger docs
	sw := swagno.New(swagno.Config{
		Title:       "Online Medicine Store API",
		Version:     "v1.0.0",
		Description: "API for managing medicines and orders.",
	})

	endpoints := []*endpoint.EndPoint{
		endpoint.New(
			endpoint.GET,
			"/medicines",
			endpoint.WithTags("medicine"),
		),
	}

	sw.AddEndpoints(endpoints)

	// medicines
	api.MedicineRoutes(router.Group(apiBase), cfg)

	// users
	api.UserRoutes(router.Group(apiBase), cfg)

	// admin
	api.AdminRoutes(router.Group(apiBase), cfg)

	// general routes
	api.GeneralRoutes(router.Group(apiBase), cfg)

	port := ":8080"

	if err := router.Run(port); err != nil {
		log.Fatalf("can't run in port %s: %v", port, err)
	}
}
