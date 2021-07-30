package server

import (
	"context"
	"fmt"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server/handler/auth"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server/handler/health"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server/handler/user"
	loggingMiddleware "github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server/middleware/logging_middleware"
	recovery "github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server/middleware/recovery_middleware"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/storage/redis_db"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/kit/command"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpAddr        string
	engine          *gin.Engine
	shutdownTimeout time.Duration
	commandBus      command.Bus
}

func New(
	ctx context.Context,
	host string,
	port uint,
	context string,
	shutdownTimeout time.Duration,
	commandBus command.Bus,
	redisdb *redis_db.Database,
) (context.Context, Server) {
	srv := Server{
		engine:          gin.New(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		shutdownTimeout: shutdownTimeout,
		commandBus:      commandBus,
	}

	log.Println(fmt.Sprintf("Check app in %s:%d/%s/%s", host, port, context, "health"))
	srv.registerRoutes(context, redisdb)
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)
	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func (s *Server) registerRoutes(context string, redisDb *redis_db.Database) {
	s.engine.Use(loggingMiddleware.Middleware(), gin.Logger(), recovery.Middleware())

	s.engine.GET(fmt.Sprintf("/%s/%s", context, "/health"), health.CheckHandler())
	s.engine.POST(fmt.Sprintf("/%s/%s", context, "/user"), user.CreateHandler())
	s.engine.GET(fmt.Sprintf("/%s/%s", context, "/user/:id"), user.GetByIdHandler())
	s.engine.DELETE(fmt.Sprintf("/%s/%s", context, "/user/:id"), user.RemoveByIdHandler())

	s.engine.POST(fmt.Sprintf("/%s/%s", context, "/auth"), auth.LoginHandler(redisDb))
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
