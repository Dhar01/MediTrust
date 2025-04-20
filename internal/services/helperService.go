package services

import (
	"context"
	"fmt"
	"medicine-app/internal/repository"
	"medicine-app/models"
)

type HelperService interface {
	ResetAllDB(ctx context.Context) error
}

type helperService struct {
	helperRepo repository.HelperRepository
	platform   string
}

func NewHelperService(repo repository.HelperRepository, platform string) HelperService {
	if repo == nil {
		panic("repo can't be empty/nil")
	}

	return &helperService{
		helperRepo: repo,
		platform:   platform,
	}
}

func (srv *helperService) ResetAllDB(ctx context.Context) error {
	if srv.platform != models.Dev {
		return fmt.Errorf("can't perform on production environment")
	}

	if err := srv.helperRepo.ResetDB(ctx); err != nil {
		return err
	}

	return nil
}

func updateField(newValue, oldValue string) string {
	if newValue == "" {
		return oldValue
	}
	return newValue
}

func updateIntPointerField(newValue, oldValue *int32) *int32 {
	if newValue == nil {
		return oldValue
	}
	return newValue
}

// func GeneralRoutes(router *echo.Echo, cfg *config.Config, baseURL string) {
// 	gen_server := newGeneralServer(cfg.DB, cfg.Platform)
// 	handlers := general_gen.NewStrictHandler(gen_server, []general_gen.StrictMiddlewareFunc{})
// 	general_gen.RegisterHandlersWithBaseURL(router, handlers, baseURL)
// }

// type generalService struct {
// 	DB  *database.Queries
// 	ENV string
// }

// func newGeneralServer(db *database.Queries, env string) *generalService {
// 	if db == nil {
// 		panic("database can't be nil")
// 	}
// 	if env == "" {
// 		panic("environment must be set")
// 	}
// 	return &generalService{
// 		DB:  db,
// 		ENV: env,
// 	}
// }

// func (gs *generalService) WipeOutDatabase(ctx context.Context, request general_gen.WipeOutDatabaseRequestObject) (general_gen.WipeOutDatabaseResponseObject, error) {
// 	if gs.ENV != models.Dev {
// 		return general_gen.UnauthorizedAccessErrorResponse{}, nil
// 	}
// 	if err := gs.DB.ResetMedicines(ctx); err != nil {
// 		return general_gen.InternalServerErrorResponse{}, err
// 	}
// 	if err := gs.DB.ResetUsers(ctx); err != nil {
// 		return general_gen.InternalServerErrorResponse{}, err
// 	}
// 	if err := gs.DB.ResetAddress(ctx); err != nil {
// 		return general_gen.InternalServerErrorResponse{}, err
// 	}
// 	return general_gen.WipeOutDatabase200Response{}, nil
// }
