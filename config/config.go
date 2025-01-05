package config

import "medicine-app/internal/database"

type Config struct {
	DB       *database.Queries
	Platform string
}
