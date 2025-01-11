package controllers

import (
	"log"
	"medicine-app/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var dev = "dev"

type controller struct {
	DB       *database.Queries
	Platform string
}

func errorMsg(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func NewController(db *database.Queries, platform string) *controller {
	return &controller{
		DB:       db,
		Platform: platform,
	}
}

func (ctrl *controller) HandlerReset(ctx *gin.Context) {
	log.Println(ctrl.Platform)

	if ctrl.Platform != dev {
		ctx.Status(http.StatusForbidden)
		return
	}

	if err := ctrl.DB.Reset(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
