package server

import (
	"github.com/gin-gonic/gin"
	"github.com/seew0/rceIdx/internal/handlers"
	"github.com/seew0/rceIdx/middleware"
)

type Server struct {
	port     string
	engine   *gin.Engine
	handlers handlers.Handler
}

func NewServer(port string, engine *gin.Engine) *Server {
	return &Server{
		port:     port,
		engine:   engine,
	}
}

func (s *Server) Run() {
	s.engine.Use(middleware.CORSmanager)

	s.engine.GET("/health", func(ctx *gin.Context) {
		s.handlers.GetHealth(ctx)
	})

	s.engine.POST("/register", func(ctx *gin.Context) {
		s.handlers.Register(ctx)
	})

	s.engine.Run(s.port)
}
