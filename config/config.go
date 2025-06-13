package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

type Config struct {
	App      AppConfig
	Database DBConfig
}

type AppConfig struct {
	Env  string
	Host string
	Port string
}

type DBConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DbName  string
	SslMode string
}

var cfgAll *Config

// GetConfig - return all config variables
func GetConfig() *Config {
	return cfgAll
}

// LoadConfig will load the configuration from .env file
func LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	var cfg Config

	server, err := server()
	if err != nil {
		return err
	}

	rdbms, err := databaseRDbms()
	if err != nil {
		return err
	}

	cfg.App = server
	cfg.Database = rdbms

	cfgAll = &cfg
	return nil
}

// helper: this function will only retrieve the key in string format
func getEnvOrErr(key string) (string, error) {
	value := strings.TrimSpace(os.Getenv(key))

	if value == "" {
		return "", &MissingEnvError{
			Key: key,
		}
	}

	return value, nil
}

// helper: this function will retrieve the key and convert to integer format
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
