package router

import (
	"log"
	"medicine-app/config"
	"medicine-app/internal/database"
	"medicine-app/internal/handler"
	"medicine-app/internal/service"
	"medicine-app/internal/store"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func SetUpRouter(conf *config.Config, e *echo.Echo) {
	dsn := database.GetDSN(conf.Database)

	db, err := database.GetDB(conf.Database)
	if err != nil {
		log.Fatalf("can't get to DB instance: %v", err)
	}

	dbPool, err := database.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("can't connect to DB: %v", err)
	}

	store := store.NewStore(db)

	v := validator.New()

	e.GET("/", handler.APIStatus)

	// Product route - medicine
	medSrv := service.NewMedService(dbPool, store)
	medHandler := handler.NewMedHandler(v, medSrv)

	e.POST("/products/medicines", medHandler.CreateMedicine)
	e.PUT("/products/medicines/:productID", medHandler.UpdateMedicine)
	e.DELETE("/products/medicines/:productID", medHandler.DeleteMedicine)
}
