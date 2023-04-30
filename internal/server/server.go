package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	cl "github.com/xlab/closer"

	"shirikaru/config"
	"shirikaru/pkg/logger"
)

var (
	server *http.Server
)

type Server struct {
	gin *gin.Engine
	cfg *config.Config
	db  *sqlx.DB
	log *logger.Logger
}

func NewServer(gin *gin.Engine, cfg *config.Config, db *sqlx.DB, log *logger.Logger) *Server {
	return &Server{gin: gin, cfg: cfg, db: db, log: log}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port))
	if err != nil {
		return err
	}

	server = &http.Server{
		Handler:      s.gin,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	cl.Bind(close)
	go func() {
		s.log.Info().Msgf("the server is running on port %s", s.cfg.Server.Port)
		s.log.Fatal().Err(server.Serve(listener))
	}()

	s.MapHandlers(s.gin)

	cl.Hold()

	return nil
}

func close() {
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	server.Shutdown(ctx)
}
