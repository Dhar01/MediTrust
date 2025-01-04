package main

import (
	"database/sql"
	"log"
	"medicine-app/config"
	api "medicine-app/handlers"
	"medicine-app/internal/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// medApp := api.MedicineApp{
	// 	DB:     dbQueries,
	// 	Router: router,
	// }

	// // for collection
	// router.POST("/medicines", medApp.CreateMedicine)
	// router.GET("/medicines", medApp.GetMedicine)
	// // for single item
	// router.DELETE("/medicines/:medicineID", medApp.DeleteMedicine)
	// router.PUT("/medicines/:medicineID", medApp.UpdateMedicine)
	// router.GET("/medicines/:medicineID", medApp.GetMedicine)

	// router.Run(":8080")

	// cfg := config.LoadConfig()
	// medApp := app.NewMedicineApp(cfg)

}

func NewMedicineApp(cfg *config.Config) *api.MedicineApp {
	router := gin.Default()

	// database connection
	dbConn, err := sql.Open("postgres", cfg.DBurl)
	if err != nil {
		log.Fatalf("can't connect to database: %v\n", err)
	}

	dbQueries := database.New(dbConn)

	return &api.MedicineApp{
		DB:     dbQueries,
		Router: router,
		Config: cfg,
	}
}
