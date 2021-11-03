package handler

import (
	"net/http"

	"github.com/alochym01/learn-web/basic-layout-handler/domain"
	"github.com/gin-gonic/gin"
)

type AlbumMemory struct {
	al []domain.Album
}

func NewAlbumMemory() AlbumMemory {
	albums := []domain.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return AlbumMemory{
		al: albums,
	}
}

// GetAlbums responds with the list of all albums as JSON.
func (a *AlbumMemory) GetAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, a.al)
}

// PostAlbums adds an album from JSON received in the request body.
func (a *AlbumMemory) PostAlbums(c *gin.Context) {
	var newAlbum domain.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	a.al = append(a.al, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumMemory) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range a.al {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// UpdateAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumMemory) UpdateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, album := range a.al {
		if album.ID == id {
			var updateAlbum domain.Album

			// Call BindJSON to bind the received JSON to updateAlbum.
			if err := c.BindJSON(&updateAlbum); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
				return
			}
			album.Title = updateAlbum.Title
			album.Artist = updateAlbum.Artist
			album.Price = updateAlbum.Price
			a.al[i] = album
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// DeleteAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumMemory) DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, album := range a.al {
		if album.ID == id {
			tempAlbum := make([]domain.Album, len(a.al)-1)
			copy(tempAlbum[:i], a.al[:i])
			copy(tempAlbum[i:], a.al[i+1:])
			a.al = tempAlbum
			c.IndentedJSON(http.StatusOK, a.al)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
