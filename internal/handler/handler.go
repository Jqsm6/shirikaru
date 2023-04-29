package handler

import (
	"github.com/gin-gonic/gin"
	"shirikaru-rest-api/internal/service"
	"shirikaru-rest-api/pkg/logger"
)

type Handler struct {
	srv *service.Service
	log *logger.Logger
}

func NewHandler(services *service.Service, log *logger.Logger) *Handler {
	return &Handler{srv: services, log: log}
}

func (h *Handler) InitRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("upload", h.upload)
		v1.GET("get/:id", h.get)
	}
}
