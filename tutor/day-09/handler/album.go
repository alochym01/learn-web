package handler

import (
	"net/http"
	"strconv"

	"github.com/alochym01/web-w-gin/domain"
	"github.com/gin-gonic/gin"
)

// AlbumHandler ...
type AlbumHandler struct {
	al []domain.Album
}

// NewAlbumHandler ...
func NewAlbumHandler() AlbumHandler {
	albums := []domain.Album{
		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return AlbumHandler{
		al: albums,
	}
}

// GetAlbums responds with the list of all albums as JSON.
func (a *AlbumHandler) GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": a.al,
	})
}

// GetAlbumsByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumHandler) GetAlbumsByID(c *gin.Context) {
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
func (a *AlbumHandler) PostAlbums(c *gin.Context) {
	var requestAlbum domain.AlbumRequest
	var newAlbum domain.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	// Data Transfer Object - DTO
	if err := c.BindJSON(&requestAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	// assign request album data to a temp Album
	newAlbum.ID = requestAlbum.ID
	newAlbum.Title = requestAlbum.Title
	newAlbum.Artist = requestAlbum.Artist
	newAlbum.Price = requestAlbum.Price

	// Add the new album to the slice.
	a.al = append(a.al, newAlbum)
	c.IndentedJSON(http.StatusCreated, gin.H{"data": newAlbum.ID})
}

// DeleteAlbums locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumHandler) DeleteAlbums(c *gin.Context) {
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
			tempAlbum := make([]domain.Album, len(a.al)-1)
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
func (a *AlbumHandler) PutAlbums(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	var requestAlbum domain.AlbumRequest

	// Call BindJSON to bind the received JSON to updateAlbum.
	// Data Transfer Object - DTO
	if err := c.BindJSON(&requestAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	requestAlbum.ID = id

	// assign request album data to a temp Album
	var updateAlbum domain.Album

	updateAlbum.ID = requestAlbum.ID
	updateAlbum.Title = requestAlbum.Title
	updateAlbum.Artist = requestAlbum.Artist
	updateAlbum.Price = requestAlbum.Price

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
