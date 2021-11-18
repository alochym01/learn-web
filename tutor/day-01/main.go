package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Album section - Start
	// Get All Albums		- method GET, URL /albums
	router.GET("/albums", getAlbums)
	// Get An Albums 		- method GET, URL /albums/:id
	// Get Create An Albums - method GET, URL /albums
	// Get Update An Albums - method PUT, URL /albums/:id
	// Get Delete An Albums - method DELETE, URL /albums/:id

	// Album section - End

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Albums",
	})
}
