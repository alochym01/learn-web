package handler

import (
	"net/http"

	"github.com/alochym01/learn-web/basic-layout-user-account/domain"
	"github.com/alochym01/learn-web/basic-layout-user-account/service"
	"github.com/gin-gonic/gin"
)

// UserHandler ...
type UserHandler struct {
	svcRepo service.UserServiceRepository
}

// NewUserHandler ...
func NewUserHandler(svc service.UserService) UserHandler {
	return UserHandler{
		svcRepo: svc,
	}
}

// GetUser responds with the list of all albums as JSON.
func (a *UserHandler) GetUsers(c *gin.Context) {
	result, err := a.svcRepo.GetUsers()
	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
	return
}

// GetUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func (a *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	result, err := a.svcRepo.ByID(id)
	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
	return
}

// // GetUserByEmail locates the user whose ID value matches the id
// // parameter sent by the client, then returns that user as a response.
// func (a *UserHandler) GetUserByEmail(c *gin.Context) {
// 	email := c.Param("email")

// 	result, err := a.svcRepo.ByEmail(email)
// 	if err != nil {
// 		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, result)
// 	return
// }

// PostUser adds an album from JSON received in the request body.
func (a *UserHandler) PostUser(c *gin.Context) {
	var userRequest domain.UserRequest

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&userRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not Enough Attribute"})
		return
	}

	// Add the new user to the slice.
	err := a.svcRepo.Create(userRequest)
	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Record is added"})
}

// UpdateUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func (a *UserHandler) UpdateUserByID(c *gin.Context) {
	id := c.Param("id")

	var userRequest domain.UserRequest

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&userRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not Enough Attribute"})
		return
	}
	err := a.svcRepo.Update(id, userRequest)
	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Record is updated"})
}

// DeleteUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func (a *UserHandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")

	err := a.svcRepo.Delete(id)
	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Record is deleted"})
}
