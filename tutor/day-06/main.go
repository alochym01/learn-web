package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Album represents data about a record Album.
type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float64
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	// Album section - Start

	// Get All Albums		- method GET, URL /albums
	router.GET("/albums", getAlbums)
	// Get An Albums 		- method GET, URL /albums/:id
	router.GET("/albums/:id", getAlbumsByID)
	// Get Create An Albums - method GET, URL /albums
	router.POST("/albums", postAlbums)
	// Get Update An Albums - method PUT, URL /albums/:id
	router.PUT("/albums/:id", putAlbums)
	// Get Delete An Albums - method DELETE, URL /albums/:id
	router.DELETE("/albums/:id", deleteAlbums)

	// Album section - End

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": albums,
	})
}

// getAlbumsByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumsByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"data": a})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, gin.H{"data": newAlbum.ID})
}

// deleteAlbums locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func deleteAlbums(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range albums {
		if a.ID == id {
			// Create a tempAlbum's element = Album's element - 1
			tempAlbum := make([]Album, len(albums)-1)
			// Copy Album element from index [0 -> i] to tempAlbum
			copy(tempAlbum, albums[:i])
			// Copy Album element from index [i + 1 -> end] to tempAlbum
			copy(tempAlbum, albums[i+1:])
			// Assign albums to tempAlbum
			albums = tempAlbum
			c.IndentedJSON(http.StatusAccepted, gin.H{"data": "albums delete successfull"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
}

// putAlbums locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func putAlbums(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	var updateAlbum Album

	// Call BindJSON to bind the received JSON to updateAlbum.
	if err := c.BindJSON(&updateAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	updateAlbum.ID = id

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range albums {
		if a.ID == id {
			albums[i] = updateAlbum
			c.IndentedJSON(http.StatusOK, gin.H{"data": "album updated"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
}
