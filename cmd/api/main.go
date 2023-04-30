package main

import (
	"github.com/gin-gonic/gin"

	"shirikaru/config"
	"shirikaru/internal/db/postgres"
	"shirikaru/internal/server"
	"shirikaru/pkg/logger"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	log := logger.GetLogger(cfg)

	db, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	g := gin.New()
	s := server.NewServer(g, cfg, db, log)

	s.Run()
}
