package handler

import (
	"net/http"

	"github.com/alochym01/learn-web/basic-layout-handler-service-error-handler/domain"
	"github.com/alochym01/learn-web/basic-layout-handler-service-error-handler/service"
	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	svcRepo service.AlbumServiceRepository
}

func NewAlbumHandler(svc service.AlbumService) AlbumHandler {
	return AlbumHandler{
		svcRepo: svc,
	}
}

// GetAlbums responds with the list of all albums as JSON.
func (a *AlbumHandler) GetAlbums(c *gin.Context) {
	result, err := a.svcRepo.GetAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
	return
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	result, err := a.svcRepo.ByID(id)
	if err != nil {
		if err.Error() == "Record Not Found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
	return
}

// PostAlbums adds an album from JSON received in the request body.
func (a *AlbumHandler) PostAlbums(c *gin.Context) {
	var albumRequest domain.AlbumRequest

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&albumRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not Enough Attribute"})
		return
	}

	// Add the new album to the slice.
	err := a.svcRepo.Create(albumRequest)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Record is added"})
}

// UpdateAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumHandler) UpdateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var albumRequest domain.AlbumRequest

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&albumRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not Enough Attribute"})
		return
	}
	err := a.svcRepo.Update(id, albumRequest)
	if err != nil {
		if err.Error() == "Record Not Found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Record is updated"})
}

// DeleteAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumHandler) DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	err := a.svcRepo.Delete(id)
	if err != nil {
		if err.Error() == "Record Not Found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Record is deleted"})
}
