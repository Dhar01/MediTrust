package config

import (
	"database/sql"
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
	Platform    string
	SecretKey   string
	EmailSender *utils.EmailSender
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	platform := getEnvVariable("PLATFORM")
	secretKey := getEnvVariable("SECRET_KEY")

	return &Config{
		Platform:    platform,
		DB:          connectDB(),
		SecretKey:   secretKey,
		EmailSender: initEmailSender(),
	}
}

func initEmailSender() *utils.EmailSender {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("can't get SMTP PORT")
	}

	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	emailFrom := os.Getenv("EMAIL_FROM")
	domain := os.Getenv("DOMAIN")

	return utils.NewEmailSender(
		smtpUser,
		smtpPass,
		emailFrom,
		domain,
		smtpHost,
		smtpPort,
	)
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
