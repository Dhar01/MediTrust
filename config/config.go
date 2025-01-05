package config

import (
	"database/sql"
	"log"
	"medicine-app/internal/database"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB       *database.Queries
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	_ = getEnvVariable("PLATFORM")

	connectDB()
}

func connectDB() *Config {
	dbURL := getEnvVariable("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("can't connect to database: %v", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Database not reachable: %v", err)
	}

	db := database.New(dbConn)

	return &Config{
		DB: db,
	}
}

func getEnvVariable(env string) string {
	envVar := os.Getenv(env)
	if envVar == "" {
		log.Fatalf("%s must be set", env)
	}

	return envVar
}
