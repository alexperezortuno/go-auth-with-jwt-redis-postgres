package user

import (
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req user.InputRequest

		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		userReq, e := user.CreateNewUser(req)
		if e != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": e,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "user created",
			"data":    userReq,
		})
	}
}

func GetByIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "" {
			log.Printf("[ERROR] %s", "you must set a valid id")
			ctx.JSON(http.StatusBadRequest, "you must set a valid id")
			return
		}

		s, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, "you must set a valid id")
			return
		}

		userReq, e := user.GetById(s)
		if e != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": e,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    userReq,
		})
	}
}

func RemoveByIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "" {
			log.Printf("[ERROR] %s", "you must set a valid id")
			ctx.JSON(http.StatusBadRequest, "you must set a valid id")
			return
		}

		s, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, "you must set a valid id")
			return
		}

		_, e := user.RemoveById(s)
		if e != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": e,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
