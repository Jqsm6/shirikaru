package anime

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Upload() gin.HandlerFunc
	GetByID() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
