package api

import (
	"deep-translate/api/detection"
	"deep-translate/api/translate"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/translate", translate.Translate)
	router.GET("/detect", detection.Detect)

	return router
}

func Run() {
	r := SetupRouter()

	r.Run()
}
