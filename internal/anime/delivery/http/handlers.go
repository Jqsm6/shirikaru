package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"shirikaru/internal/anime"
	"shirikaru/internal/model"
	"shirikaru/pkg/logger"
)

type animeHandlers struct {
	animeUC anime.UseCase
	log     *logger.Logger
}

func NewAnimeHandlers(animeUC anime.UseCase, log *logger.Logger) anime.Handlers {
	return &animeHandlers{animeUC: animeUC, log: log}
}

func (ah *animeHandlers) Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var anime *model.Anime

		if err := c.BindJSON(&anime); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		id, err := ah.animeUC.Upload(c.Request.Context(), anime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload anime"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func (ah *animeHandlers) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		animeList, err := ah.animeUC.GetAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get anime"})
			return
		}

		c.JSON(http.StatusOK, animeList)
	}
}

func (ah *animeHandlers) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed convert string to int"})
			return
		}

		anime, err := ah.animeUC.GetByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get anime"})
			return
		}

		c.JSON(http.StatusOK, anime)
	}
}

func (ah *animeHandlers) SearchByTitle() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Param("title")

		animeList, err := ah.animeUC.SearchByTitle(c.Request.Context(), title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get anime"})
			return
		} else if animeList == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no anime found for this query"})
			return
		}

		c.JSON(http.StatusOK, animeList)
	}
}

func (ah *animeHandlers) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed convert string to int"})
			return
		}

		err = ah.animeUC.Delete(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete anime"})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{})
	}
}
