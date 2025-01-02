package medicines

import (
	"medicine-app/internal/database"

	"github.com/gin-gonic/gin"
)

type MedicineApp struct {
	DB     *database.Queries
	Router *gin.Engine
}

func errorMsg(err error) gin.H {
	return gin.H{"error": err.Error()}
}
