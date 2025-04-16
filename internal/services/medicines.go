package services

import (
	"context"

	"medicine-app/config"
	med "medicine-app/internal/api/medicines_gen"
	"medicine-app/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MedicineRoutes(router *gin.RouterGroup, cfg *config.Config) {
	medServer := newMedicineServer(cfg.DB)
	handlers := med.NewStrictHandler(medServer, []med.StrictMiddlewareFunc{})
	med.RegisterHandlers(router, handlers)
}

type medicineService struct {
	DB *database.Queries
}

var _ med.StrictServerInterface = (*medicineService)(nil)

func newMedicineServer(db *database.Queries) *medicineService {
	if db == nil {
		panic("database can't be nil")
	}

	return &medicineService{
		DB: db,
	}
}

func (mc *medicineService) CreateNewMedicine(ctx context.Context, request med.CreateNewMedicineRequestObject) (med.CreateNewMedicineResponseObject, error) {
	if request.Body == nil {
		return med.CreateNewMedicine400Response{}, nil
	}

	medicine, err := mc.DB.CreateMedicine(ctx, database.CreateMedicineParams{
		Name:         request.Body.Name,
		Dosage:       request.Body.Dosage,
		Description:  request.Body.Description,
		Manufacturer: request.Body.Manufacturer,
		Price:        request.Body.Price,
		Stock:        request.Body.Stock,
	})

	if err != nil {
		return med.CreateNewMedicine500Response{}, err
	}

	return med.CreateNewMedicine201JSONResponse(toMedicineDomain(medicine)), nil
}

func (mc *medicineService) FetchMedicineList(ctx context.Context, request med.FetchMedicineListRequestObject) (med.FetchMedicineListResponseObject, error) {
	medicines, err := mc.DB.GetMedicines(ctx)
	if err != nil {
		return med.FetchMedicineList500Response{}, err
	}

	var medList []med.Medicine

	for _, medicine := range medicines {
		medList = append(medList, toMedicineDomain(medicine))
	}

	return med.FetchMedicineList200JSONResponse(medList), nil
}

func (mc *medicineService) FetchMedicineByID(ctx context.Context, request med.FetchMedicineByIDRequestObject) (med.FetchMedicineByIDResponseObject, error) {
	if request.MedicineID == uuid.Nil {
		return med.FetchMedicineByID400Response{}, nil
	}

	medicine, err := mc.DB.GetMedicine(ctx, request.MedicineID)
	if err != nil {
		return med.FetchMedicineByID500Response{}, err
	}

	return med.FetchMedicineByID200JSONResponse(toMedicineDomain(medicine)), nil
}

func (mc *medicineService) UpdateMedicineInfoByID(ctx context.Context, request med.UpdateMedicineInfoByIDRequestObject) (med.UpdateMedicineInfoByIDResponseObject, error) {
	if request.Body == nil {
		return med.UpdateMedicineInfoByID400Response{}, nil
	}

	oldInfo, err := mc.DB.GetMedicine(ctx, request.MedicineID)
	if err != nil {
		return med.UpdateMedicineInfoByID404Response{}, err
	}

	updateInfo, err := mc.DB.UpdateMedicine(ctx, database.UpdateMedicineParams{
		ID:           request.MedicineID,
		Name:         updateField(*request.Body.Name, oldInfo.Name),
		Dosage:       updateField(*request.Body.Dosage, oldInfo.Dosage),
		Description:  updateField(*request.Body.Description, oldInfo.Description),
		Manufacturer: updateField(*request.Body.Manufacturer, oldInfo.Manufacturer),
		Price:        *updateIntPointerField(request.Body.Price, &oldInfo.Price),
		Stock:        *updateIntPointerField(request.Body.Stock, &oldInfo.Stock),
	})
	if err != nil {
		return med.UpdateMedicineInfoByID500Response{}, err
	}

	return med.UpdateMedicineInfoByID202JSONResponse(toMedicineDomain(updateInfo)), nil
}

func (mc *medicineService) DeleteMedicineByID(ctx context.Context, request med.DeleteMedicineByIDRequestObject) (med.DeleteMedicineByIDResponseObject, error) {
	if request.MedicineID == uuid.Nil {
		return med.DeleteMedicineByID400Response{}, nil
	}

	if err := mc.DB.DeleteMedicine(ctx, request.MedicineID); err != nil {
		return med.DeleteMedicineByID500Response{}, err
	}

	return med.DeleteMedicineByID204Response{}, nil
}

func toMedicineDomain(dbMed database.Medicine) med.Medicine {
	return med.Medicine{
		Id:           (*uuid.UUID)(&dbMed.ID),
		Name:         &dbMed.Name,
		Dosage:       &dbMed.Dosage,
		Description:  &dbMed.Description,
		Manufacturer: &dbMed.Manufacturer,
		Price:        &dbMed.Price,
		Stock:        &dbMed.Stock,
	}
}
