package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"now":    time.Now(),
	})
}
