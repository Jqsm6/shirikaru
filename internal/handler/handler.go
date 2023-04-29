package handler

import (
	"github.com/gin-gonic/gin"

	"shirikaru-rest-api/internal/logger"
	"shirikaru-rest-api/internal/service"
)

type Handler struct {
	srv *service.Service
	log *logger.Logger
}

func NewHandler(services *service.Service, log *logger.Logger) *Handler {
	return &Handler{srv: services, log: log}
}

func (uh *Handler) InitRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("upload", uh.upload)
		v1.DELETE("delete/:id", uh.delete)
		v1.GET("get/:id", uh.get)
	}
}
