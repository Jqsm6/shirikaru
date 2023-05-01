package anime

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Upload() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetByID() gin.HandlerFunc
	SearchByTitle() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
