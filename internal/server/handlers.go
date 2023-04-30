package server

import (
	"github.com/gin-gonic/gin"

	animeHandlers "shirikaru/internal/anime/delivery/http"
	animeRepository "shirikaru/internal/anime/repository"
	animeUC "shirikaru/internal/anime/usecase"
)

func (s *Server) MapHandlers(g *gin.Engine) {
	aRepo := animeRepository.NewAnimeRepository(s.db)
	aUC := animeUC.NewAnimeUseCase(aRepo, s.log)
	aHandlers := animeHandlers.NewAnimeHandlers(aUC, s.log)

	animeGroup := g.Group("/anime")
	animeGroup.Use(gin.Logger())
	animeHandlers.MapAnimeRoutes(animeGroup, aHandlers)
}
