package handler

import (
	"errors"
	"medicine-app/internal/errs"
	"net/http"

	"github.com/labstack/echo/v4"
)

// APIStatus returns the status of API
func APIStatus(e echo.Context) error {
	return e.String(http.StatusOK, "live")
}

// helper: setErrorResp checks the error and make the response according to it
func setErrorResp(err error) *echo.HTTPError {
	switch {
	case errors.Is(err, errs.ErrNotFound):
		return echo.NewHTTPError(http.StatusNotFound, "medicine not found")
	case errors.Is(err, errs.ErrConflict):
		return echo.NewHTTPError(http.StatusConflict, "same medicine exist")
	case errors.Is(err, errs.ErrInvalidInput):
		return echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
	}
}
