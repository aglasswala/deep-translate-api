package detection

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Detect(c *gin.Context) {
	reader, header, err := c.Request.FormFile("upload")
	if err != nil {
		log.Println(err)
	}

	defer reader.Close()

	// TODO: Validate JSON

	result, err := DetectObjects(reader, header.Filename)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, result)
}
