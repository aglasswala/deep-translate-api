package translate

import (
	"github.com/gin-gonic/gin"
)

func Translate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
