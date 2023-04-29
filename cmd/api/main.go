package main

import (
	"shirikaru-rest-api/internal/handler"
	server "shirikaru-rest-api/internal/httpserver"
	"shirikaru-rest-api/internal/repository"
	"shirikaru-rest-api/internal/service"
	"shirikaru-rest-api/pkg/db/postgres"
	"shirikaru-rest-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := logger.GetLogger()

	router := gin.New()

	db, err := postgres.NewPsqlDB()
	if err != nil {
		panic(err)
	}

	server := server.NewServer(router, logger)
	repos := repository.NewRepository(db, logger)
	srv := service.NewServices(repos)
	handlers := handler.NewHandler(srv, logger)

	handlers.InitRoutes(router)

	server.Run()
}
