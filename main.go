package main

import (
	"database/sql"
	"log"
	"medicine-app/handlers"
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

	r := gin.Default()

	medApp := handlers.MedicineApp{
		DB:     dbQueries,
		Router: r,
	}

	r.POST("/medicines", medApp.CreateMedicine)
	r.GET("/medicines", medApp.GetMedicine)

	r.Run(":8080")

}

func getEnvVariable(env string) string {
	envVar := os.Getenv(env)
	if envVar == "" {
		log.Fatalf("%s must be set", env)
	}

	return envVar
}
