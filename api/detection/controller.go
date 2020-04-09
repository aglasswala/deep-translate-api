package detection

import (
	"github.com/gin-gonic/gin"
)

func Detect(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "detection",
	})
}
