package handler

import (
	"medicine-app/internal/model"
	"medicine-app/internal/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// medHandler defines the handler structure for medicine route
type medHandler struct {
	validator *validator.Validate
	srv       *service.MedService
}

// NewMedHandler returns a medHandler instance for the medicine route
func NewMedHandler(valid *validator.Validate, srv *service.MedService) *medHandler {
	return &medHandler{
		validator: valid,
		srv:       srv,
	}
}

// CreateMedicine handler creates a new medicine on the database.
// Admin only access.
// POST /api/v1/products/medicines
func (h *medHandler) CreateMedicine(e echo.Context) error {
	var req model.MedicineRequest

	if err := e.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	result, err := h.srv.CreateMedicine(e.Request().Context(), req)
	if err != nil {
		return setErrorResp(err)
	}

	return e.JSON(http.StatusCreated, result.AdminResponse())
}

// UpdateMedicine handler updates medicine info on the database.
// Admin only access.
// PUT /api/v1/products/medicines
func (h *medHandler) UpdateMedicine(e echo.Context) error {
	var req model.MedicineRequest

	if err := e.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	result, err := h.srv.UpdateMedicine(e.Request().Context(), req)
	if err != nil {
		return setErrorResp(err)
	}

	return e.JSON(http.StatusAccepted, result.AdminResponse())
}

// DeleteMedicine handler deletes medicine info on the database
// Admin only access
// DELETE /api/v1/products/medicines/:productID
func (h *medHandler) DeleteMedicine(e echo.Context) error {
	id := e.Param("productID")

	// validate the ID
	productID, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "productID isn't valid")
	}

	if err := h.srv.Delete(e.Request().Context(), productID); err != nil {
		return setErrorResp(err)
	}

	return e.NoContent(http.StatusOK)
}
