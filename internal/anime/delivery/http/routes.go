package http

import (
	"github.com/gin-gonic/gin"

	"shirikaru/internal/anime"
)

func MapAnimeRoutes(animeGroup *gin.RouterGroup, h anime.Handlers) {
	animeGroup.POST(upload, h.Upload())
	animeGroup.GET(getByID, h.GetByID())
	animeGroup.GET(getByTitle, h.GetByTitle())
	animeGroup.DELETE(delete, h.Delete())
}
