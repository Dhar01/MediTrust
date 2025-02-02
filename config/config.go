package config

import (
	"database/sql"
	"log"
	"medicine-app/internal/database"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	DB        *database.Queries
	Platform  string
	SecretKey string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	platform := getEnvVariable("PLATFORM")
	secretKey := getEnvVariable("SECRET_KEY")

	// domainName := getEnvVariable("DOMAIN_NAME")
	// companyName := getEnvVariable("COMPANY_NAME")
	// backendEmail := getEnvVariable("EMAIL")

	return &Config{
		Platform:  platform,
		DB:        connectDB(),
		SecretKey: secretKey,
	}
}

func connectDB() *database.Queries {
	dbURL := getEnvVariable("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("can't connect to database: %v", err)
	}
	// defer dbConn.Close()

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Database not reachable: %v", err)
	}

	return database.New(dbConn)
}

func getEnvVariable(env string) string {
	envVar := os.Getenv(env)
	if envVar == "" {
		log.Fatalf("%s must be set", env)
	}

	return envVar
}
