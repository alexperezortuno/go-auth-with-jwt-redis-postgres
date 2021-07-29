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
	if err != nil {
		log.Printf("error parsing port")
		port = 8080
	}

	host := os.Getenv("APP_HOST")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbTimeZone := os.Getenv("DB_TIME_ZONE")
	timeZone := os.Getenv("APP_TIME_ZONE")
	engineSql := os.Getenv("DB_DRIVER")
	context := os.Getenv("APP_CONTEXT")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASS")

	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDb = 0
	}

	if host == "" {
		host = "localhost"
	}

	if context == "" {
		context = "api"
	}

	if dbHost == "" {
		dbHost = "db"
	}

	if dbPort == "" {
		dbPort = "5432"
	}

	if dbUser == "" {
		dbUser = "postgres"
	}

	if dbPass == "" {
		dbPass = "postgres"
	}

	if dbName == "" {
		dbName = "dbauth"
	}

	if engineSql == "" {
		engineSql = "postgres"
	}

	if dbTimeZone == "" {
		dbTimeZone = "America/Santiago"
	}

	if timeZone == "" {
		timeZone = "America/Santiago"
	}

	if redisHost == "" {
		redisHost = "redis"
	}

	if redisPort == "" {
		redisPort = "6379"
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
