package router

import (
	"github.com/alochym01/learn-web/basic-layout-repository/domain"
	"github.com/alochym01/learn-web/basic-layout-repository/storage"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	mem := storage.NewAlbumMemory()
	// Album routes
	router.GET("/albums", mem.GetAlbums)
	router.GET("/albums/:id", mem.GetAlbumByID)       // GET HTTP method with /albums/id will be routed to getAlbumByID function
	router.POST("/albums", mem.PostAlbums)            // POST HTTP method with /albums will be routed to getAlbums function
	router.PUT("/albums/:id", mem.UpdateAlbumByID)    // PUT HTTP method with /albums/:id will be routed to updateAlbumsByID function
	router.DELETE("/albums/:id", mem.DeleteAlbumByID) // DELETE HTTP method with /albums/:id will be routed to deleteAlbumsByID function

	router.GET("/sqlalbums", domain.SQLGetAlbums)
	router.GET("/sqlalbums/:id", domain.SQLGetAlbumByID)       // GET HTTP method with /sqlalbums/id will be routed to SQLGetAlbumByID function
	router.POST("/sqlalbums", domain.SQLPostAlbums)            // POST HTTP method with /sqlalbums will be routed to SQLGetAlbums function
	router.PUT("/sqlalbums/:id", domain.SQLUpdateAlbumByID)    // PUT HTTP method with /sqlalbums/:id will be routed to SQLUpdateAlbumsByID function
	router.DELETE("/sqlalbums/:id", domain.SQLDeleteAlbumByID) // DELETE HTTP method with /sqlalbums/:id will be routed to SQLDeleteAlbumsByID functio

	return router
}
