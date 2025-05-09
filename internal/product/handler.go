package product

import (
	"errors"
	"medicine-app/internal/errs"
	"net/http"

	"github.com/labstack/echo/v4"
)

type medicineHandler struct {
	service medService
}

// type medicineAPI struct {
// 	medService MedService
// }

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

// var _ StrictServerInterface = (*medicineAPI)(nil)

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
	return nil
}

func (h *medicineHandler) CreateMedicine(ctx echo.Context) error {
	return nil
}

func (h *medicineHandler) FetchMedicineByID(ctx echo.Context, medicineID MedicineID) error {
	med, err := h.service.FetchMedicineByID(ctx.Request().Context(), medicineID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return echo.ErrNotFound
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	resp := toResponse(med)

	return ctx.JSON(http.StatusOK, resp)
}

func (h *medicineHandler) UpdateMedicineInfoByID(ctx echo.Context, medicineID MedicineID) error {
	return nil
}

func (h *medicineHandler) DeleteMedicineByID(ctx echo.Context, medicineID MedicineID) error {
	return nil
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

// func (api *medicineAPI) FetchMedicineList(ctx context.Context, request FetchMedicineListRequestObject) (FetchMedicineListResponseObject, error) {
// 	return FetchMedicineList200JSONResponse{}, nil
// }

// func (api *medicineAPI) CreateNewMedicine(ctx context.Context, request CreateNewMedicineRequestObject) (CreateNewMedicineResponseObject, error) {
// 	medicine, err := api.medService.CreateMedicine(ctx, CreateMedicineDTO{
// 		Name:         request.Body.Name,
// 		Dosage:       request.Body.Dosage,
// 		Manufacturer: request.Body.Manufacturer,
// 		Description:  request.Body.Description,
// 		Price:        request.Body.Price,
// 		Stock:        request.Body.Stock,
// 	})
// 	if err != nil {
// 		return InternalServerErrorResponse{}, err
// 	}
// 	return CreateNewMedicine201JSONResponse(toResponseDomain(medicine)), nil
// }

// func (api *medicineAPI) DeleteMedicineByID(ctx context.Context, request DeleteMedicineByIDRequestObject) (DeleteMedicineByIDResponseObject, error) {
// 	if err := api.medService.DeleteMedicine(ctx, request.MedicineID); err != nil {
// 		return BadRequestErrorResponse{}, err
// 	}
// 	return DeleteMedicineByID204Response{}, nil
// }

// func (api *medicineAPI) UpdateMedicineInfoByID(ctx context.Context, request UpdateMedicineInfoByIDRequestObject) (UpdateMedicineInfoByIDResponseObject, error) {
// 	med, err := api.medService.UpdateMedicine(ctx, request.MedicineID, medicine{
// 		Name:         *request.Body.Name,
// 		Dosage:       *request.Body.Dosage,
// 		Description:  *request.Body.Description,
// 		Manufacturer: *request.Body.Manufacturer,
// 		Price:        *request.Body.Price,
// 		Stock:        *request.Body.Stock,
// 	})
// 	if errors.Is(err, errs.ErrMedicineNotExist) {
// 		return NotFoundErrorResponse{}, errs.ErrMedicineNotExist
// 	}
// 	if err != nil {
// 		return InternalServerErrorResponse{}, err
// 	}
// 	return UpdateMedicineInfoByID202JSONResponse(toResponseDomain(med)), nil
// }

// func (mc *medicineService) FetchMedicineList(ctx context.Context, request med_gen.FetchMedicineListRequestObject) (med_gen.FetchMedicineListResponseObject, error) {
// 	medicines, err := mc.DB.GetMedicines(ctx)
// 	if err != nil {
// 		return med_gen.InternalServerErrorResponse{}, err
// 	}
// 	var medList []med_gen.Medicine
// 	for _, medicine := range medicines {
// 		medList = append(medList, toResponseDomain(medicine))
// 	}
// 	return med_gen.FetchMedicineList200JSONResponse(medList), nil
// }
