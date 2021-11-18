package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album represents data about a record Album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

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
		"message": albums,
	})
}
