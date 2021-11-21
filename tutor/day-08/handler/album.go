package handler

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

// AlbumHandler
type AlbumHandlder struct {
	al []Album
}

// NewAlbumHandler ...
func NewAlbumHandler() AlbumHandlder {
	albums := []Album{
		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return AlbumHandlder{
		al: albums,
	}
}

// GetAlbums responds with the list of all albums as JSON.
func (a *AlbumHandlder) GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": a.al,
	})
}

// GetAlbumsByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumHandlder) GetAlbumsByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range a.al {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"data": album})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
}

// PostAlbums adds an album from JSON received in the request body.
func (a *AlbumHandlder) PostAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	// Add the new album to the slice.
	a.al = append(a.al, newAlbum)
	c.IndentedJSON(http.StatusCreated, gin.H{"data": newAlbum.ID})
}

// DeleteAlbums locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumHandlder) DeleteAlbums(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, album := range a.al {
		if album.ID == id {
			// Create a tempAlbum's element = Album's element - 1
			tempAlbum := make([]Album, len(a.al)-1)
			// Copy Album element from index [0 -> i] to tempAlbum
			copy(tempAlbum, a.al[:i])
			// Copy Album element from index [i + 1 -> end] to tempAlbum
			copy(tempAlbum, a.al[i+1:])
			// Assign albums to tempAlbum
			a.al = tempAlbum
			c.IndentedJSON(http.StatusAccepted, gin.H{"data": "albums delete successfull"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
}

// PutAlbums locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumHandlder) PutAlbums(c *gin.Context) {
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
	for i, album := range a.al {
		if album.ID == id {
			a.al[i] = updateAlbum
			c.IndentedJSON(http.StatusOK, gin.H{"data": "album updated"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
}
