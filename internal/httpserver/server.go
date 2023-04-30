package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cl "github.com/xlab/closer"

	"shirikaru-rest-api/config"
	"shirikaru-rest-api/internal/logger"
)

var server *http.Server

type Server struct {
	router *gin.Engine
	logger *logger.Logger
	server *http.Server
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

	server = &http.Server{
		Handler:      s.router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	cl.Bind(close)

	go func() {
		log.Info().Msgf("the server is running on port %s", cfg.Server.Port)
		log.Fatal().Err(server.Serve(listener))
	}()

	cl.Hold()
}

func close() {
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	server.Shutdown(ctx)
}
