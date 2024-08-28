package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/config", getConfig)
	router.Run("localhost:8080")
}

func getConfig(c *gin.Context) {
	config_data := getFileData("./config.json")
	c.IndentedJSON(http.StatusOK, config_data)
}
