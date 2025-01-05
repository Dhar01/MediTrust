package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	ctrl "medicine-app/controllers"
	"medicine-app/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	platform := getEnvVariable("PLATFORM")
	dburl := getEnvVariable("DB_URL")
	dbConn, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatalf("can't connect to database: %v\n", err)
	}

	dbQueries := database.New(dbConn)

	cfg := ctrl.Config{
		Platform: platform,
		DB:       dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/medicines", cfg.CreateMedicineHandler)
	mux.HandleFunc("/medicines/:medID", cfg.DeleteMedicine)

}

func getEnvVariable(env string) string {
	envVar := os.Getenv(env)
	if envVar == "" {
		log.Fatalf("%s must be set", env)
	}

	return envVar
}
