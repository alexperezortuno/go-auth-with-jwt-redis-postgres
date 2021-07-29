package bootstrap

import (
	"context"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/bus/inmemory"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server"
	db "github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/data_base"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/tools/environment"
)

var params = environment.Server()

func Run() error {
	var (
		commandBus = inmemory.NewCommandBus()
	)

	db.Init(params)
	defer db.CloseConn()
	db.Migrate()

	ctx, srv := server.New(context.Background(), params.Host, uint(params.Port), params.Context, params.ShutdownTimeout, commandBus)
	return srv.Run(ctx)
}
