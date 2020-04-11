package detection

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Detect(c *gin.Context) {
	reader, header, err := c.Request.FormFile("upload")
	if err != nil {
		log.Println("ERROR: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer reader.Close()

	result, err := DetectObjects(reader, header.Filename)
	if err != nil {
		log.Println("ERROR: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
