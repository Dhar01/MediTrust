package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DocumentationRoute(router *gin.RouterGroup) {
	// Serving Swagger docs

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
