package config

import (
	"context"
	"fmt"
	"log"
	"medicine-app/internal/database"
	"medicine-app/utils"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Config struct {
	DB          *database.DB
	DBConn      *pgxpool.Pool
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
		return nil, fmt.Errorf("invalid port number")
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
func connectDB() (*pgxpool.Pool, *database.DB, error) {
	dbURL := mustGetEnv("DB_URL")

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, nil, err
	}

	if err = pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, nil, err
	}

	return pool, database.New(pool), nil
}

// mustGetEnv retrieves an environment variable or logs a fatal error if missing.
func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s required but not set", key)
	}

	return value
}
