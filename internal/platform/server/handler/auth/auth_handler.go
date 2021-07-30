package auth

import (
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/auth"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/redis_db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginHandler(redisDb *redis_db.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req auth.AuthRequest

		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := auth.ValidateUser(req, redisDb)
		if err != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func VerifyHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, true)
	}
}
