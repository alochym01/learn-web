package router

import (
	"github.com/alochym01/learn-web/basic-layout/domain"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	// Album routes
	router.GET("/albums", domain.GetAlbums)
	router.GET("/albums/:id", domain.GetAlbumByID)       // GET HTTP method with /albums/id will be routed to getAlbumByID function
	router.POST("/albums", domain.PostAlbums)            // POST HTTP method with /albums will be routed to getAlbums function
	router.PUT("/albums/:id", domain.UpdateAlbumByID)    // PUT HTTP method with /albums/:id will be routed to updateAlbumsByID function
	router.DELETE("/albums/:id", domain.DeleteAlbumByID) // DELETE HTTP method with /albums/:id will be routed to deleteAlbumsByID function

	return router
}
