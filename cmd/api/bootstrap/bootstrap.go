package bootstrap

import (
	"context"
	"fmt"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/bus/inmemory"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server"
	db "github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/data_base"
	redisdb "github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/redis_db"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/tools/environment"
	"log"
)

var params = environment.Server()

func Run() error {
	var (
		commandBus = inmemory.NewCommandBus()
	)

	var rStr = fmt.Sprintf("%s:%s", params.RedisHost, params.RedisPort)
	redisDb, err := redisdb.NewDatabase(rStr, params.RedisPass, params.RedisDb)
	if err != nil {
		log.Fatalf("Failed to connect to redis_db: %s", err.Error())
	}

	db.Init(params)
	defer db.CloseConn()
	db.Migrate()

	ctx, srv := server.New(
		context.Background(),
		params.Host,
		uint(params.Port),
		params.Context,
		params.ShutdownTimeout,
		commandBus,
		redisDb,
	)
	return srv.Run(ctx)
}
