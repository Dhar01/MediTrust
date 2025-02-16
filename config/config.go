package config

import (
	"database/sql"
	"fmt"
	"log"
	"medicine-app/internal/database"
	"medicine-app/utils"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	DB          *database.Queries
	DBConn      *sql.DB
	Platform    string
	Domain      string
	Port        string
	SecretKey   string
	EmailSender *utils.EmailSender
}

// LoadConfig initializes and returns the application configuration
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	platform := mustGetEnv("PLATFORM")
	secretKey := mustGetEnv("SECRET_KEY")
	port := mustGetEnv("PORT")
	domain := mustGetEnv("DOMAIN")

	checkPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	if checkPort < 5000 || checkPort > 10000 {
		return nil, fmt.Errorf("Invalid port number")
	}

	dbConn, dbQueries, err := connectDB()
	if err != nil {
		return nil, err
	}

	emailSender, err := initEmailSender()
	if err != nil {
		return nil, err
	}

	return &Config{
		Platform:    platform,
		SecretKey:   secretKey,
		DB:          dbQueries,
		DBConn:      dbConn,
		Domain:      domain,
		Port:        port,
		EmailSender: emailSender,
	}, nil
}

// initEmailSender initializes and returns an EmailSender instance.
func initEmailSender() (*utils.EmailSender, error) {
	smtpHost := mustGetEnv("SMTP_HOST")
	smtpPortStr := mustGetEnv("SMTP_PORT")
	smtpUser := mustGetEnv("SMTP_USER")
	smtpPass := mustGetEnv("SMTP_PASS")
	emailFrom := mustGetEnv("EMAIL_FROM")

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return nil, err
	}

	return utils.NewEmailSender(
		smtpUser,
		smtpPass,
		emailFrom,
		smtpHost,
		smtpPort,
	), nil
}

// connectDB initializes and returns a database connection
func connectDB() (*sql.DB, *database.Queries, error) {
	dbURL := mustGetEnv("DB_URL")

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, nil, err
	}

	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		return nil, nil, err
	}

	return dbConn, database.New(dbConn), nil
}

// mustGetEnv retrieves an environment variable or logs a fatal error if missing.
func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s required but not set", key)
	}

	return value
}
