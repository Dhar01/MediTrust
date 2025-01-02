package main

import (
	"database/sql"
	"log"
	med "medicine-app/handlers"
	"medicine-app/internal/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	dbURL := getEnvVariable("DB_URL")

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("can't connect to database: %v\n", err)
	}
	dbQueries := database.New(dbConn)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	medApp := med.MedicineApp{
		DB:     dbQueries,
		Router: router,
	}

	// for collection
	router.POST("/medicines", medApp.CreateMedicine)
	router.GET("/medicines", medApp.GetMedicine)

	// for single item
	router.DELETE("/medicines/:medicineID", medApp.DeleteMedicine)
	router.PUT("/medicines/:medicineID", medApp.UpdateMedicine)
	router.GET("/medicines/:medicineID", medApp.GetMedicine)

	router.Run(":8080")

}

func getEnvVariable(env string) string {
	envVar := os.Getenv(env)
	if envVar == "" {
		log.Fatalf("%s must be set", env)
	}

	return envVar
}
