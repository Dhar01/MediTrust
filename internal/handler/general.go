package handler

import (
	"context"
	general_gen "medicine-app/internal/handler/gen_gen"
	"medicine-app/internal/services"
)

type helperAPI struct {
	helperService services.HelperService
}

var _ general_gen.StrictServerInterface = (*helperAPI)(nil)

func newHelperAPI(srv services.HelperService) *helperAPI {
	if srv == nil {
		panic("helper service can't be empty/nil")
	}

	return &helperAPI{
		helperService: srv,
	}
}

func (api *helperAPI) WipeOutDatabase(ctx context.Context, request general_gen.WipeOutDatabaseRequestObject) (general_gen.WipeOutDatabaseResponseObject, error) {
	if err := api.helperService.ResetAllDB(ctx); err != nil {
		return general_gen.UnauthorizedAccessErrorResponse{}, err
	}

	return general_gen.WipeOutDatabase200Response{}, nil
}