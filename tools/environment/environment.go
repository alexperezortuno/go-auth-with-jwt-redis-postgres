package environment

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
)

type ServerValues struct {
	Host            string
	Port            int
	ShutdownTimeout time.Duration
	Context         string
	DbUser          string
	DbPass          string
	DbHost          string
	DbPort          string
	DbName          string
	DbTimeout       time.Duration
	DbTimeZone      string
	EngineSql       string
	TimeZone        string
	RedisHost       string
	RedisPort       string
	RedisPass       string
	RedisDb         int
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
	if err != nil {
		log.Printf("error parsing port")
		port = 8080
	}

	host := getEnv("APP_HOST", "localhost")
	dbHost := getEnv("DB_HOST", "db")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASS", "postgres")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "dbauth")
	dbTimeZone := getEnv("DB_TIME_ZONE", "America/Santiago")
	timeZone := getEnv("APP_TIME_ZONE", "America/Santiago")
	engineSql := getEnv("DB_DRIVER", "postgres")
	context := getEnv("APP_CONTEXT", "api")
	redisHost := getEnv("REDIS_HOST", "redis")
	redisPort := getEnv("REDIS_PORT", "6380")
	redisPass := getEnv("REDIS_PASS", "123")

	redisDb, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
	if err != nil {
		redisDb = 0
	}

	return ServerValues{
		Host:            host,
		Context:         context,
		Port:            port,
		TimeZone:        timeZone,
		DbHost:          dbHost,
		DbPort:          dbPort,
		DbUser:          dbUser,
		DbPass:          dbPass,
		DbName:          dbName,
		DbTimeZone:      dbTimeZone,
		RedisHost:       redisHost,
		RedisPort:       redisPort,
		RedisPass:       redisPass,
		RedisDb:         redisDb,
		ShutdownTimeout: 10 * time.Second,
		EngineSql:       engineSql,
	}
}

func getEnv(envName, valueDefault string) string {
	value := os.Getenv(envName)
	if value == "" {
		return valueDefault
	}
	return value
}
