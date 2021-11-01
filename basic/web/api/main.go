package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Album - define Album table with its attribute
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// updateAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range albums {
		if a.ID == id {
			var updateAlbum Album

			// Call BindJSON to bind the received JSON to updateAlbum.
			if err := c.BindJSON(&updateAlbum); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
				return
			}
			a.Title = updateAlbum.Title
			a.Artist = updateAlbum.Artist
			a.Price = updateAlbum.Price
			albums[i] = a
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// deleteAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range albums {
		if a.ID == id {
			tempAlbum := make([]Album, len(albums)-1)
			copy(tempAlbum, albums[:i])
			copy(tempAlbum, albums[i+1:])
			albums = tempAlbum
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	fmt.Println("Server is starting with port 8080")
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)       // GET HTTP method with /albums/id will be routed to getAlbumByID function
	router.POST("/albums", postAlbums)            // POST HTTP method with /albums will be routed to getAlbums function
	router.PUT("/albums/:id", updateAlbumByID)    // PUT HTTP method with /albums/:id will be routed to updateAlbumsByID function
	router.DELETE("/albums/:id", deleteAlbumByID) // DELETE HTTP method with /albums/:id will be routed to deleteAlbumsByID function

	router.Run("localhost:8080")
}
