// Tutorial: https://go.dev/doc/tutorial/web-service-gin

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

var albums = []album{
	{ID: "1", Title: "My Neck", Artist: "MyBack", Price: 3.99},
	{ID: "2", Title: "Lick'm I. Puss E.", Artist: "Andmai Krak", Price: 4.99},
}
