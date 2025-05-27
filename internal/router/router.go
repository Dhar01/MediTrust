package router

import (
	"medicine-app/config"
	"medicine-app/internal/handler"

	"github.com/labstack/echo/v4"
)

func SetUpRouter(conf *config.Configuration, e *echo.Echo) {
	e.GET("/", handler.APIStatus)
}
