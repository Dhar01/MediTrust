package product

import (
	"log"
	"medicine-app/config"
	"medicine-app/internal/database"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(router echo.Group, cfg config.Configuration) {
	db, err := database.GetDB(cfg.Database.RDbms)
	if err != nil {
		log.Fatalf("Can't initialize database, %v", err)
	}

	repo := newMedicineRepo(db)
	svc := newMedicineService(repo)
	api := newMedicineHandler(svc)

	RegisterHandlersWithBaseURL(&router, api, config.BaseURL)
}
