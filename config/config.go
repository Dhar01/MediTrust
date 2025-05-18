package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

const (
	Activated string = "yes"
	BaseURL   string = "/api/v1"
)

type MissingEnvError struct {
	Key string
}

func (e *MissingEnvError) Error() string {
	return fmt.Sprintf("environment variable %s is required", e.Key)
}

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

	server, err := server()
	if err != nil {
		return err
	}

	rDbms, err := databaseRDbms()
	if err != nil {
		return err
	}

	redis, err := databaseRedis()
	if err != nil {
		return err
	}

	config.Server = server
	config.Database.RDbms = rDbms
	config.Database.Redis = redis

	configAll = &config

	return nil
}

// helper functions
func getEnvOrErr(key string) (string, error) {
	value := strings.TrimSpace(os.Getenv(key))

	if value == "" {
		return "", &MissingEnvError{
			Key: key,
		}
	}

	return value, nil
}

func getEnvNumber(key string) (int, error) {
	value, err := getEnvOrErr(key)
	if err != nil {
		return 0, err
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, &MissingEnvError{
			Key: key,
		}
	}

	return number, nil
}
