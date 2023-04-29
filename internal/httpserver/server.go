package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"

	"shirikaru-rest-api/config"
	"shirikaru-rest-api/internal/logger"
)

type Server struct {
	router *gin.Engine
	logger *logger.Logger
}

func NewServer(router *gin.Engine, logger *logger.Logger) *Server {
	return &Server{
		router: router,
		logger: logger,
	}
}

func (s *Server) Run(cfg *config.Config, log *logger.Logger) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      s.router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	log.Fatal().Err(server.Serve(listener))
}
