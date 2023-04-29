package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"shirikaru-rest-api/internal/model"
)

func (uh *Handler) upload(c *gin.Context) {
	var anime model.Anime

	if err := c.BindJSON(&anime); err != nil {
		uh.log.Err(err).Msg("controller level")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	id, err := uh.srv.Upload(c, &anime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload anime"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "id": id})
}

func (uh *Handler) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		uh.log.Err(err).Msg("controller level")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed convert string to int"})
		return
	}

	anime, err := uh.srv.Get(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get anime"})
		return
	}

	c.JSON(http.StatusOK, anime)
}
