package api

import (
	"context"
	"errors"
	med_gen "medicine-app/internal/api/medicines_gen"
	"medicine-app/internal/errs"
	"medicine-app/internal/services"
	"medicine-app/models"

	"github.com/google/uuid"
)

type medicineAPI struct {
	medService services.MedService
}

var _ med_gen.StrictServerInterface = (*medicineAPI)(nil)

func newMedicineAPI(srv services.MedService) *medicineAPI {
	if srv == nil {
		panic("medicine service can't be empty/nil")
	}

	return &medicineAPI{
		medService: srv,
	}
}

func (api *medicineAPI) FetchMedicineList(ctx context.Context, request med_gen.FetchMedicineListRequestObject) (med_gen.FetchMedicineListResponseObject, error) {

	return med_gen.FetchMedicineList200JSONResponse{}, nil
}

func (api *medicineAPI) CreateNewMedicine(ctx context.Context, request med_gen.CreateNewMedicineRequestObject) (med_gen.CreateNewMedicineResponseObject, error) {
	medicine, err := api.medService.CreateMedicine(ctx, models.CreateMedicineDTO{
		Name:         request.Body.Name,
		Dosage:       request.Body.Dosage,
		Manufacturer: request.Body.Manufacturer,
		Description:  request.Body.Description,
		Price:        request.Body.Price,
		Stock:        request.Body.Stock,
	})

	if err != nil {
		return med_gen.InternalServerErrorResponse{}, err
	}

	return med_gen.CreateNewMedicine201JSONResponse(toMedicineDomain(medicine)), nil
}

func (api *medicineAPI) DeleteMedicineByID(ctx context.Context, request med_gen.DeleteMedicineByIDRequestObject) (med_gen.DeleteMedicineByIDResponseObject, error) {
	if err := api.medService.DeleteMedicine(ctx, request.MedicineID); err != nil {
		return med_gen.BadRequestErrorResponse{}, err
	}

	return med_gen.DeleteMedicineByID204Response{}, nil
}

func (api *medicineAPI) FetchMedicineByID(ctx context.Context, request med_gen.FetchMedicineByIDRequestObject) (med_gen.FetchMedicineByIDResponseObject, error) {
	medicine, err := api.medService.FetchMedicineByID(ctx, request.MedicineID)
	if err != nil {
		return med_gen.NotFoundErrorResponse{}, err
	}

	return med_gen.FetchMedicineByID200JSONResponse(toMedicineDomain(medicine)), nil
}

func (api *medicineAPI) UpdateMedicineInfoByID(ctx context.Context, request med_gen.UpdateMedicineInfoByIDRequestObject) (med_gen.UpdateMedicineInfoByIDResponseObject, error) {
	medicine, err := api.medService.UpdateMedicine(ctx, request.MedicineID, models.UpdateMedicineDTO{
		Name:         *request.Body.Name,
		Dosage:       *request.Body.Dosage,
		Description:  *request.Body.Description,
		Manufacturer: *request.Body.Manufacturer,
		Price:        *request.Body.Price,
		Stock:        *request.Body.Stock,
	})

	if errors.Is(err, errs.ErrMedicineNotExist) {
		return med_gen.NotFoundErrorResponse{}, errs.ErrMedicineNotExist
	}

	if err != nil {
		return med_gen.InternalServerErrorResponse{}, err
	}

	return med_gen.UpdateMedicineInfoByID202JSONResponse(toMedicineDomain(medicine)), nil
}

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

func toMedicineDomain(medicine *models.Medicine) med_gen.Medicine {
	return med_gen.Medicine{
		Id:           (*uuid.UUID)(&medicine.Id),
		Name:         &medicine.Name,
		Dosage:       &medicine.Dosage,
		Description:  &medicine.Description,
		Manufacturer: &medicine.Manufacturer,
		Price:        &medicine.Price,
		Stock:        &medicine.Stock,
	}
}
