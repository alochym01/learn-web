package handler

import (
	"net/http"
	"strconv"

	"github.com/alochym01/web-w-gin/domain"
	"github.com/alochym01/web-w-gin/service"
	"github.com/gin-gonic/gin"
)

// AlbumHandler ...
type AlbumHandler struct {
	svcRepo service.AlbumServiceRepository
}

// NewAlbumHandler ...
func NewAlbumHandler(svc service.AlbumService) AlbumHandler {
	return AlbumHandler{
		svcRepo: svc,
	}
}

// GetAlbums responds with the list of all albums as JSON.
func (a *AlbumHandler) GetAlbums(c *gin.Context) {
	albums, err := a.svcRepo.GetAlbums()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"data": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		// Get all Albums by using service
		"data": albums,
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

	// Get Album by using service
	album, errs := a.svcRepo.ByID(id)

	if errs != nil {
		c.IndentedJSON(errs.Code, gin.H{"data": errs.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": album})
}

// PostAlbums adds an album from JSON received in the request body.
func (a *AlbumHandler) PostAlbums(c *gin.Context) {
	var requestAlbum domain.AlbumRequest

	// Call BindJSON to bind the received JSON to newAlbum.
	// Data Transfer Object - DTO
	if err := c.BindJSON(&requestAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	// Create an Album by using service
	temp, err := a.svcRepo.Create(requestAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"data": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": *temp})
}

// DeleteAlbums locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumHandler) DeleteAlbums(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
		return
	}

	// Delete an Album by using service
	err = a.svcRepo.Delete(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"data": "albums delete successfull"})
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

	// Update an Album by using service
	err = a.svcRepo.Update(id, requestAlbum)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": "album updated"})
}
