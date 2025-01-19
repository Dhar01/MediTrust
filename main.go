package main

import (
	"log"
	"medicine-app/api"
	"medicine-app/config"

	"github.com/gin-gonic/gin"
)

const apiBase = "/api/v1"

// @title           Medicine-Shop Swagger API
// @version         1.0
// @description     Swagger API for Medicine Shop.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Loknath Dhar
// @contact.email  loknathdhar01@yahoo.com

// @license.name  MIT
// @license.url

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	cfg := config.NewConfig()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	// medicines
	api.MedicineRoutes(router.Group(apiBase), cfg)

	// users
	api.UserRoutes(router.Group(apiBase), cfg)

	// // admin
	// api.AdminRoutes(router.Group(apiBase), cfg)

	// general routes
	api.GeneralRoutes(router.Group(apiBase), cfg)

	port := ":8080"

	if err := router.Run(port); err != nil {
		log.Fatalf("can't run in port %s: %v", port, err)
	}
}
