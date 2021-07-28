package environment

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

type ServerValues struct {
	Host    string
	Port    int
	Context string
}

func EnvVariable(key string) string {
	return os.Getenv(key)
}

func env() {
	env := os.Getenv("APP_ENV")

	if env == "" || env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func Server() ServerValues {
	env()
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	host := os.Getenv("APP_HOST")
	context := os.Getenv("APP_CONTEXT")

	if err != nil {
		log.Printf("error parsing port")
		port = 8080
	}

	if host == "" {
		host = "_APP_NAME_"
	}

	if context == "" {
		context = "api"
	}

	return ServerValues{
		Host:    host,
		Context: context,
		Port:    port,
	}
}
