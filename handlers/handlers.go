package handlers

import (
	"medicine-app/config"
	"medicine-app/internal/database"

	"github.com/gin-gonic/gin"
)

type MedicineApp struct {
	DB     *database.Queries
	Router *gin.Engine
	Config *config.Config
}
