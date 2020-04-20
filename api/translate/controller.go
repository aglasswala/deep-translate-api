package translate

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type TranslateJSON struct {
	Words          []string `json:"words"`
	TargetLanguage string   `json:"targetlanguage"`
}

func Translate(c *gin.Context) {
	var json TranslateJSON
	c.BindJSON(&json)

	log.Println(os.Getenv("AZURE_SUBSCRIPTION_KEY"))

	if "" == os.Getenv("AZURE_SUBSCRIPTION_KEY") {
		log.Fatal("Please set/export the environment variable AZURE_SUBSCRIPTION_KEY.")
	}
	subscriptionKey := os.Getenv("AZURE_SUBSCRIPTION_KEY")
	if "" == os.Getenv("AZURE_TRANSLATE_ENDPOINT") {
		log.Fatal("Please set/export the environment variable AZURE_TRANSLATE_ENDPOINT.")
	}
	endpoint := os.Getenv("AZURE_TRANSLATE_ENDPOINT")
	uri := endpoint + "/translate?api-version=3.0"

	words, err := TranslateObjects(json.Words, json.TargetLanguage, uri, subscriptionKey)
	if err != nil {
		log.Println("ERROR: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"words":    string(words),
		"language": json.TargetLanguage,
	})
}
