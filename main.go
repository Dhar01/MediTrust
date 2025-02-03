package main

import (
	"log"
	"medicine-app/api"
	"medicine-app/config"

	_ "medicine-app/docs"

	"github.com/gin-gonic/gin"
)

const apiBase = "/api/v1"

//	@title			Online Medicine Store API
//	@version		1.0
//	@description	Documentation of api of online medicine store.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	cfg := config.NewConfig()

	// uncomment this line for production
	// gin.SetMode(gin.ReleaseMode)

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
