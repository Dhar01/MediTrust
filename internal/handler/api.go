package handler

import (
	"log"
	"medicine-app/config"
	auth_gen "medicine-app/internal/handler/authUser_gen"
	general_gen "medicine-app/internal/handler/gen_gen"
	med_gen "medicine-app/internal/handler/medicines_gen"
	"medicine-app/internal/handler/public_gen"
	"medicine-app/internal/middleware"
	"medicine-app/internal/repository"
	"medicine-app/internal/services"

	"github.com/labstack/echo/v4"
)

func MedicineRoutes(router *echo.Echo, cfg *config.Config, baseURL string) {
	repo := repository.NewMedicineRepo(&cfg.DB.Medicine)
	srv := services.NewMedicineService(repo)
	api := newMedicineAPI(srv)

	middle := middleware.NewMiddleware(cfg)

	server := med_gen.NewStrictHandler(api, []med_gen.StrictMiddlewareFunc{
		middle.IsAdmin,
	})

	med_gen.RegisterHandlersWithBaseURL(router, server, baseURL)
}

func PublicRoutes(router *echo.Echo, cfg *config.Config, baseURL string) {
	repo := repository.NewPublicRepo(&cfg.DB.User)
	srv := services.NewPublicService(repo, cfg)
	api := newPublicAPI(srv)
	server := public_gen.NewStrictHandler(api, nil)
	public_gen.RegisterHandlersWithBaseURL(router, server, baseURL)
}

func AuthUserRoutes(router *echo.Echo, cfg *config.Config, baseURL string) {
	authRepo := repository.NewAuthUserRepo(&cfg.DB.User)
	publicRepo := repository.NewPublicRepo(&cfg.DB.User)
	authSrv := services.NewAuthUserService(authRepo, publicRepo, cfg)
	pubSrv := services.NewPublicService(publicRepo, cfg)
	api := newAuthUserAPI(authSrv, pubSrv)

	middle := middleware.NewMiddleware(cfg)

	server := auth_gen.NewStrictHandler(api, []med_gen.StrictMiddlewareFunc{
		middle.IsUser,
	})

	auth_gen.RegisterHandlersWithBaseURL(router, server, baseURL)
}

func GeneralRoutes(router *echo.Echo, cfg *config.Config, baseURL string) {
	helpRepo := repository.NewHelperRepo(&cfg.DB.Helper)
	log.Println(cfg.Platform)
	helpSrv := services.NewHelperService(helpRepo, cfg.Platform)
	api := newHelperAPI(helpSrv)
	server := general_gen.NewStrictHandler(api, nil)
	general_gen.RegisterHandlersWithBaseURL(router, server, baseURL)
}
