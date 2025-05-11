package product

import (
	"errors"
	"medicine-app/internal/errs"
	"net/http"

	"github.com/labstack/echo/v4"
)

// func MedicineRoutes(router *echo.Echo, cfg *config.Config, baseURL string) {
// 	repo := NewMedicineRepo(&cfg.DB.Medicine)
// 	srv := NewMedicineService(repo)
// 	api := newMedicineAPI(srv)
// 	// middle := middleware.NewMiddleware(cfg)
// 	// server := NewStrictHandler(api, []StrictMiddlewareFunc{
// 	// 	// Adapt(middle.IsAdmin),
// 	// })
// 	RegisterHandlersWithBaseURL(router, server, baseURL)
// }


type medicineHandler struct {
	service medService
}

var _ ServerInterface = (*medicineHandler)(nil)

func newMedicineHandler(srv medService) *medicineHandler {
	if srv == nil {
		panic("medicine service can't be nil")
	}

	return &medicineHandler{
		service: srv,
	}
}

func (h *medicineHandler) FetchMedicineList(ctx echo.Context) error {
	medicines, err := h.service.FetchMedicineList(ctx.Request().Context())
	if err != nil {
		// ! need to handle different error cases
		return echo.NewHTTPError(http.StatusInternalServerError, "no entry found")
	}

	return ctx.JSON(http.StatusOK, medicines)
}

func (h *medicineHandler) CreateMedicine(ctx echo.Context) error {
	/*
		- learning note!

		We could use `new()` to create a pointer to the request struct for binding, but it would
		allocate on the heap, which requires garbage collection. That's slightly slower, although
		irrelevant for most web apps.

		Instead, we're going to use a stack-allocated value and passing its address to the binder.
		It would be idiomatic and avoids unnecessary heap use.
	*/
	var req medicineRequest

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request")
	}

	result, err := h.service.CreateMedicine(ctx.Request().Context(), fromRequest(req))
	// ! need to handle different types of errors
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "can't create medicine")
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (h *medicineHandler) FetchMedicineByID(ctx echo.Context, medicineID MedicineID) error {
	med, err := h.service.FetchMedicineByID(ctx.Request().Context(), medicineID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return echo.ErrNotFound
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return ctx.JSON(http.StatusOK, toResponse(med))
}

func (h *medicineHandler) UpdateMedicineInfoByID(ctx echo.Context, medicineID MedicineID) error {
	var req medicineRequest

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request")
	}

	med, err := h.service.UpdateMedicine(ctx.Request().Context(), medicineID, fromRequest(req))
	// ! need to handle different types of errors
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "can't update medicine")
	}

	return ctx.JSON(http.StatusAccepted, med)
}

func (h *medicineHandler) DeleteMedicineByID(ctx echo.Context, medicineID MedicineID) error {
	if err := h.service.DeleteMedicine(ctx.Request().Context(), medicineID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "need to work")
	}

	// return ctx.JSON(http.StatusAccepted, map[string]string{
	// 	"message": "medicine deleted",
	// })

	return ctx.NoContent(http.StatusAccepted)
}

func toResponse(m *medicine) MedicineResponse {
	return MedicineResponse{
		ID:           m.ID,
		Name:         m.Name,
		Manufacturer: m.Manufacturer,
		Dosage:       m.Dosage,
		Description:  m.Description,
		Price:        m.Price,
		Stock:        m.Stock,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func fromRequest(req MedicineRequest) medicine {
	return medicine{
		Name:         req.Name,
		Manufacturer: req.Manufacturer,
		Dosage:       req.Dosage,
		Description:  req.Description,
		Price:        req.Price,
		Stock:        req.Stock,
	}
}
