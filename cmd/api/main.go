package main

import (
	"shirikaru-rest-api/config"
	"shirikaru-rest-api/internal/db/postgres"
	"shirikaru-rest-api/internal/handler"
	server "shirikaru-rest-api/internal/httpserver"
	"shirikaru-rest-api/internal/logger"
	"shirikaru-rest-api/internal/repository"
	"shirikaru-rest-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	log := logger.GetLogger(cfg)

	router := gin.New()

	db, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	httpserver := server.NewServer(router, log)
	repos := repository.NewRepository(db, log)
	srv := service.NewServices(repos)
	handlers := handler.NewHandler(srv, log)

	handlers.InitRoutes(router)

	httpserver.Run(cfg, log)
}
