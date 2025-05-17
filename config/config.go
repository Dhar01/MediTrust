package config

import (
	"log"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

const Activated string = "yes"

type Configuration struct {
	Database DatabaseConfig
	Server   ServerConfig
}

var configAll *Configuration

// GetConfig - return all config variables
func GetConfig() *Configuration {
	return configAll
}

// LoadConfig - load all available configurations
func LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	var config Configuration

	config.Server = server()

	rDbms, err := databaseRDbms()
	if err != nil {
		return err
	}
	config.Database.RDbms = rDbms

	redis, err := databaseRedis()
	if err != nil {
		return err
	}
	config.Database.Redis = redis

	configAll = &config

	return nil
}

// helper function
func mustGetEnv(key string) string {
	value := strings.TrimSpace(os.Getenv(key))

	if value == "" {
		log.Fatalf("ENV variable %s is required", key)
	}

	return value
}
