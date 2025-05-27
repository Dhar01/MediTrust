package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// APIStatus returns the status of API
func APIStatus(e echo.Context) error {
	return e.String(http.StatusOK, "live")
}
