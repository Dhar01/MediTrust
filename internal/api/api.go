package api

import (
	"medicine-app/config"
	med_gen "medicine-app/internal/api/medicines_gen"
	user_gen "medicine-app/internal/api/users_gen"
	"medicine-app/internal/repository"
	"medicine-app/internal/services"

	"github.com/labstack/echo/v4"
)

// type API struct {
// 	services *services.Services
// 	MedAPI   MedicineAPI
// }

// func NewAPI(srv *services.Services) *API {
// 	if srv == nil {
// 		panic("services can't be nil")
// 	}
// 	return &API{
// 		services: srv,
// 		MedAPI:   NewMedicineAPI(srv.MedService),
// 	}
// }

func MedicineRoutes(router *echo.Echo, cfg *config.Config, baseURL string) {
	repo := repository.NewMedicineRepo(&cfg.DB.Medicine)
	srv := services.NewMedicineService(repo)
	api := newMedicineAPI(srv)
	server := med_gen.NewStrictHandler(api, nil)
	med_gen.RegisterHandlersWithBaseURL(router, server, baseURL)
}

func UserRoutes(router *echo.Echo, cfg *config.Config, baseURL string) {
	repo := repository.NewUserRepo(&cfg.DB.User)
	srv := services.NewUserService(repo)
	api := newUserAPI(srv)
	server := user_gen.NewStrictHandler(api, nil)
	user_gen.RegisterHandlersWithBaseURL(router, server, baseURL)
}
