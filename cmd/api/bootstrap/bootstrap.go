package bootstrap

import (
    "github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/tools/environment"
)

var params = environment.Server()

func Run() error {
	srv := server.New(params.Host, uint(params.Port), params.Context)
	return srv.Run()
}
