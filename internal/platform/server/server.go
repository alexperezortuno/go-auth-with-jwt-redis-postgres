package server

import (
	"fmt"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server/handler/health"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server/handler/logger"
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/internal/platform/server/middleware/logging_middleware"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint, context string) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	log.Println(fmt.Sprintf("Check app in %s:%d/%s/%s", host, port, context, "health"))
	srv.registerRoutes(context)
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes(context string) {
	s.engine.Use(logging_middleware.Middleware())
	s.engine.Use(gin.Logger())
	s.engine.Use(gin.Recovery())

	s.engine.GET(fmt.Sprintf("/%s/%s", context, "/health"), health.CheckHandler())
}
