package main

import (
	"log"
	"medicine-app/api"
	"medicine-app/config"
	"medicine-app/models"

	_ "medicine-app/docs"

	"github.com/gin-gonic/gin"
)

const apiBase = "/api/v1"

//	@title			MediTrust Backend API
//	@version		1.0
//	@description	Documentation of api of online medicine pharmacy - MediTrust

//	@contact.name	API Support
//	@contact.url	http://t.me/Dhar01
//	@contact.email	loknathdhar66@gmail.com

//	@license.name	GPL v3
//	@license.url	https://www.gnu.org/licenses/gpl-3.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							cookie
//	@name						refresh-token

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
	api.MedicineRoutes(router.Group(apiBase), cfg)

	// users
	api.UserRoutes(router.Group(apiBase), cfg)

	// admin
	api.AdminRoutes(router.Group(apiBase), cfg)

	// general routes
	api.GeneralRoutes(router.Group(apiBase), cfg)

	// documentation routes
	api.DocumentationRoute(router.Group(apiBase))

	port := ":8080"

	if err := router.Run(port); err != nil {
		log.Fatalf("can't run in port %s: %v", port, err)
	}
}
