package api

import (
	"context"
	med_gen "medicine-app/internal/api/medicines_gen"
	"medicine-app/internal/services"
)

type MedicineAPI struct {
	services *services.MedService
}

var _ med_gen.StrictServerInterface = (*MedicineAPI)(nil)

func NewMedicineAPI(srv *services.MedService) MedicineAPI {
	return MedicineAPI{
		services: srv,
	}
}

func (api MedicineAPI) FetchMedicineList(ctx context.Context, request med_gen.FetchMedicineListRequestObject) (med_gen.FetchMedicineListResponseObject, error) {
	return med_gen.FetchMedicineList200JSONResponse{}, nil
}

func (api MedicineAPI) CreateNewMedicine(ctx context.Context, request med_gen.CreateNewMedicineRequestObject) (med_gen.CreateNewMedicineResponseObject, error) {
	return med_gen.CreateNewMedicine201JSONResponse{}, nil
}

func (api MedicineAPI) DeleteMedicineByID(ctx context.Context, request med_gen.DeleteMedicineByIDRequestObject) (med_gen.DeleteMedicineByIDResponseObject, error) {
	return med_gen.DeleteMedicineByID204Response{}, nil
}

func (api MedicineAPI) FetchMedicineByID(ctx context.Context, request med_gen.FetchMedicineByIDRequestObject) (med_gen.FetchMedicineByIDResponseObject, error) {
	return med_gen.FetchMedicineByID200JSONResponse{}, nil
}

func (api MedicineAPI) UpdateMedicineInfoByID(ctx context.Context, request med_gen.UpdateMedicineInfoByIDRequestObject) (med_gen.UpdateMedicineInfoByIDResponseObject, error) {
	return med_gen.UpdateMedicineInfoByID202JSONResponse{}, nil
}

// func MedicineRoutes(router *echo.Echo, cfg *config.Config, baseURL string) {
// 	medServer := newMedicineServer(cfg.DB)
// 	handlers := med_gen.NewStrictHandler(medServer, []med_gen.StrictMiddlewareFunc{})
// 	med_gen.RegisterHandlersWithBaseURL(router, handlers, baseURL)
// }

// type medicineService struct {
// 	DB *database.Queries
// }

// var _ med_gen.StrictServerInterface = (*medicineService)(nil)

// func newMedicineServer(db *database.Queries) *medicineService {
// 	if db == nil {
// 		panic("database can't be nil")
// 	}
// 	return &medicineService{
// 		DB: db,
// 	}
// }

// func (mc *medicineService) CreateNewMedicine(ctx context.Context, request med_gen.CreateNewMedicineRequestObject) (med_gen.CreateNewMedicineResponseObject, error) {
// 	if request.Body == nil {
// 		return med_gen.BadRequestErrorResponse{}, nil
// 	}
// 	medicine, err := mc.DB.CreateMedicine(ctx, database.CreateMedicineParams{
// 		Name:         request.Body.Name,
// 		Dosage:       request.Body.Dosage,
// 		Description:  request.Body.Description,
// 		Manufacturer: request.Body.Manufacturer,
// 		Price:        request.Body.Price,
// 		Stock:        request.Body.Stock,
// 	})
// 	if err != nil {
// 		return med_gen.InternalServerErrorResponse{}, err
// 	}
// 	return med_gen.CreateNewMedicine201JSONResponse(toMedicineDomain(medicine)), nil
// }

// func (mc *medicineService) FetchMedicineList(ctx context.Context, request med_gen.FetchMedicineListRequestObject) (med_gen.FetchMedicineListResponseObject, error) {
// 	medicines, err := mc.DB.GetMedicines(ctx)
// 	if err != nil {
// 		return med_gen.InternalServerErrorResponse{}, err
// 	}
// 	var medList []med_gen.Medicine
// 	for _, medicine := range medicines {
// 		medList = append(medList, toMedicineDomain(medicine))
// 	}
// 	return med_gen.FetchMedicineList200JSONResponse(medList), nil
// }

// func (mc *medicineService) FetchMedicineByID(ctx context.Context, request med_gen.FetchMedicineByIDRequestObject) (med_gen.FetchMedicineByIDResponseObject, error) {
// 	if request.MedicineID == uuid.Nil {
// 		return med_gen.BadRequestErrorResponse{}, nil
// 	}
// 	medicine, err := mc.DB.GetMedicine(ctx, request.MedicineID)
// 	if err != nil {
// 		return med_gen.InternalServerErrorResponse{}, err
// 	}
// 	return med_gen.FetchMedicineByID200JSONResponse(toMedicineDomain(medicine)), nil
// }

// func (mc *medicineService) UpdateMedicineInfoByID(ctx context.Context, request med_gen.UpdateMedicineInfoByIDRequestObject) (med_gen.UpdateMedicineInfoByIDResponseObject, error) {
// 	if request.Body == nil {
// 		return med_gen.BadRequestErrorResponse{}, nil
// 	}
// 	oldInfo, err := mc.DB.GetMedicine(ctx, request.MedicineID)
// 	if err != nil {
// 		return med_gen.NotFoundErrorResponse{}, err
// 	}
// 	updateInfo, err := mc.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
// 		ID:           request.MedicineID,
// 		Name:         updateField(*request.Body.Name, oldInfo.Name),
// 		Dosage:       updateField(*request.Body.Dosage, oldInfo.Dosage),
// 		Description:  updateField(*request.Body.Description, oldInfo.Description),
// 		Manufacturer: updateField(*request.Body.Manufacturer, oldInfo.Manufacturer),
// 		Price:        *updateIntPointerField(request.Body.Price, &oldInfo.Price),
// 		Stock:        *updateIntPointerField(request.Body.Stock, &oldInfo.Stock),
// 	})
// 	if err != nil {
// 		return med_gen.InternalServerErrorResponse{}, err
// 	}
// 	return med_gen.UpdateMedicineInfoByID202JSONResponse(toMedicineDomain(updateInfo)), nil
// }

// func (mc *medicineService) DeleteMedicineByID(ctx context.Context, request med_gen.DeleteMedicineByIDRequestObject) (med_gen.DeleteMedicineByIDResponseObject, error) {
// 	if request.MedicineID == uuid.Nil {
// 		return med_gen.BadRequestErrorResponse{}, nil
// 	}
// 	if err := mc.DB.DeleteMedicine(ctx, request.MedicineID); err != nil {
// 		return med_gen.InternalServerErrorResponse{}, err
// 	}
// 	return med_gen.DeleteMedicineByID204Response{}, nil
// }

// func toMedicineDomain(dbMed database.Medicine) med_gen.Medicine {
// 	return med_gen.Medicine{
// 		Id:           (*uuid.UUID)(&dbMed.ID),
// 		Name:         &dbMed.Name,
// 		Dosage:       &dbMed.Dosage,
// 		Description:  &dbMed.Description,
// 		Manufacturer: &dbMed.Manufacturer,
// 		Price:        &dbMed.Price,
// 		Stock:        &dbMed.Stock,
// 	}
// }
