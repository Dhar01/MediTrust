package main

import (
	"log"
	"medicine-app/api"
	"medicine-app/config"

	"github.com/gin-gonic/gin"
)

const apiBase = "/api/v1"

func main() {
	cfg := config.NewConfig()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	// medicines
	api.MedicineRoutes(router.Group(apiBase), cfg)

	// users
	api.UserRoutes(router.Group(apiBase), cfg)

	// general routes
	api.GeneralRoutes(router.Group(apiBase), cfg)

	port := ":8080"

	if err := router.Run(port); err != nil {
		log.Fatalf("cant' run in port %s: %v", port, err)
	}
}
