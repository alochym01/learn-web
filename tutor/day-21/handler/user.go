package handler

import (
	"net/http"
	"strconv"

	"github.com/alochym01/web-w-gin/domain"
	"github.com/alochym01/web-w-gin/service"
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
		c.IndentedJSON(err.Code, gin.H{"data": err.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
	return
}

// GetUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func (a *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Not Enough Attribute"})
		return
	}

	result, errs := a.svcRepo.ByID(id)
	if errs != nil {
		c.IndentedJSON(errs.Code, gin.H{"data": errs.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": result})
	return
}

// PostUser adds an album from JSON received in the request body.
func (a *UserHandler) PostUser(c *gin.Context) {
	var userRequest domain.UserRequest

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&userRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Not Enough Attribute"})
		return
	}

	// Add the new user to the slice.
	id, errs := a.svcRepo.Create(userRequest)
	if errs != nil {
		c.IndentedJSON(errs.Code, gin.H{"data": errs.Message})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"data": id})
}

// // UpdateUserByID locates the user whose ID value matches the id
// // parameter sent by the client, then returns that user as a response.
// func (a *UserHandler) UpdateUserByID(c *gin.Context) {
// 	id := c.Param("id")

// 	var userRequest domain.UserRequest

// 	// Call BindJSON to bind the received JSON to newAlbum.
// 	if err := c.BindJSON(&userRequest); err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Not Enough Attribute"})
// 		return
// 	}
// 	err := a.svcRepo.Update(id, userRequest)
// 	if err != nil {
// 		c.IndentedJSON(err.Code, gin.H{"data": err.Message})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusAccepted, gin.H{"data": "Record is updated"})
// }

// // DeleteUserByID locates the user whose ID value matches the id
// // parameter sent by the client, then returns that user as a response.
// func (a *UserHandler) DeleteUserByID(c *gin.Context) {
// 	id := c.Param("id")

// 	err := a.svcRepo.Delete(id)
// 	if err != nil {
// 		c.IndentedJSON(err.Code, gin.H{"data": err.Message})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusAccepted, gin.H{"data": "Record is deleted"})
// }
