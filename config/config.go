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
	Platform string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	platform := getEnvVar("PLATFORM")
	dbURL := getEnvVar("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("can't connect to database: %v\n", err)
	}

	dbQuries := database.New(dbConn)

	return &Config{
		DB:       dbQuries,
		Platform: platform,
	}
}

func getEnvVar(env string) string {
	envVar := os.Getenv(env)
	if envVar == "" {
		log.Fatalf("%s must be set", env)
	}

	return envVar
}
