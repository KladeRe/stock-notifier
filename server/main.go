package main

import (
	"net/http"

	"github.com/KladeRe/stock-server/internal/database"
	"github.com/KladeRe/stock-server/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type DelID struct {
	HexID string `json:"message"`
}

func main() {

	db_client := database.EstablishDBClient()

	defer db_client.CloseClient()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from server",
		})
	})
	router.GET("/configs", func(c *gin.Context) {
		documents, err := db_client.GetAllDocuments()

		if err != nil {
			c.IndentedJSON(http.StatusNoContent, err)
			return
		}
		c.IndentedJSON(http.StatusOK, documents)
	})

	router.POST("/configs/add", func(c *gin.Context) {
		var newConfig database.StockConfig

		if err := c.BindJSON(&newConfig); err != nil {
			return
		}
		db_client.AddDocument(newConfig)
		c.IndentedJSON(http.StatusCreated, newConfig)

	})

	router.POST("/configs/del", func(c *gin.Context) {
		var deletion_id DelID

		if err := c.BindJSON(&deletion_id); err != nil {
			return
		}

		deletionPrimitive, err := primitive.ObjectIDFromHex(deletion_id.HexID)

		if err != nil {
			return
		}
		db_client.DeleteDocument(deletionPrimitive)
		c.IndentedJSON(http.StatusOK, deletion_id)
	})
	port, portError := utils.GetEnvVariable("BACKEND_PORT")
	if portError != nil {
		port = "5050"
	}

	router.Run(":" + port)
}
