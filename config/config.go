package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBurl string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	dbURL := getEnvVar("DB_URL")

	return &Config{
		DBurl: dbURL,
	}
}

func getEnvVar(env string) string {
	envVar := os.Getenv(env)
	if envVar == "" {
		log.Fatalf("%s must be set", env)
	}

	return envVar
}