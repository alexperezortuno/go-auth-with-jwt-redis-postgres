package auth

import (
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req auth.AuthRequest

		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		_, _ = auth.ValidateUser(req)
		ctx.JSON(http.StatusOK, gin.H{
			"token": "...",
		})
	}
}
