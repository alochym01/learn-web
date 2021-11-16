package handler

import (
	"net/http"

	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/domain"
	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/service"
	"github.com/gin-gonic/gin"
)

// LoginHandler ...
type LoginHandler struct {
	svcRepo service.LoginServiceRepository
}

// NewLoginHandler ...
func NewLoginHandler(svc service.LoginService) LoginHandler {
	return LoginHandler{
		svcRepo: svc,
	}
}

// Login responds with the list of all albums as JSON.
func (a *LoginHandler) Login(c *gin.Context) {
	var loginRequest domain.LoginRequest

	// Call BindJSON to bind the received JSON to LoginRequest.
	if err := c.BindJSON(&loginRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not Enough Attribute"})
		return
	}

	result, err := a.svcRepo.ByEmail(loginRequest)
	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
	return
}
