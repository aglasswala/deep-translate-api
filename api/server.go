package api

import (
	"log"

	"deep-translate/api/detection"
	"deep-translate/api/translate"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR:", err)
	}

	router.POST("/translate", translate.Translate)
	router.POST("/detect", detection.Detect)

	return router
}

func Run() {
	r := SetupRouter()

	r.Run()
}
