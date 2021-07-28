package logger

import (
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req logger.InputRequest
		if err := context.BindJSON(&req); err != nil {
			log.Printf("error: %s\n", err)
			context.JSON(http.StatusBadRequest, err)
			return
		}

		logger.LogRequest(req)

		context.JSON(http.StatusOK, gin.H{
			"message": "logger created",
		})
	}
}
