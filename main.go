package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type fruits struct {
	ID   string `json:"id"`
	Name string `json:"title"`
}

var albums = []fruits{
	{ID: "1", Name: "Apple"},
	{ID: "2", Name: "Banana"},
}

func getFruitsName(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/fruits", getFruitsName)
	router.Run("localhost:8080")
}
