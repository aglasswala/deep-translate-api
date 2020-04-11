package translate

import (
	"github.com/gin-gonic/gin"
)

type TranslateJSON struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func Translate(c *gin.Context) {
	var json TranslateJSON
	c.BindJSON(&json)

	c.JSON(200, gin.H{
		"name":    json.Name,
		"message": json.Message,
	})
}
