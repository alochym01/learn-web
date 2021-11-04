package router

import (
	"database/sql"

	"github.com/alochym01/learn-web/basic-layout-handler-service-error-handler/handler"
	"github.com/alochym01/learn-web/basic-layout-handler-service-error-handler/service"
	"github.com/alochym01/learn-web/basic-layout-handler-service-error-handler/storage"
	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// mem := handler.NewAlbumMemory()
	// // Album routes
	// router.GET("/albums", mem.GetAlbums)
	// router.GET("/albums/:id", mem.GetAlbumByID)       // GET HTTP method with /albums/id will be routed to getAlbumByID function
	// router.POST("/albums", mem.PostAlbums)            // POST HTTP method with /albums will be routed to getAlbums function
	// router.PUT("/albums/:id", mem.UpdateAlbumByID)    // PUT HTTP method with /albums/:id will be routed to updateAlbumsByID function
	// router.DELETE("/albums/:id", mem.DeleteAlbumByID) // DELETE HTTP method with /albums/:id will be routed to deleteAlbumsByID function

	// sql := handler.NewAlbumSQL(db)
	// router.GET("/sqlalbums", sql.SQLGetAlbums)
	// router.GET("/sqlalbums/:id", sql.SQLGetAlbumByID)       // GET HTTP method with /sqlalbums/id will be routed to SQLGetAlbumByID function
	// router.POST("/sqlalbums", sql.SQLPostAlbums)            // POST HTTP method with /sqlalbums will be routed to SQLGetAlbums function
	// router.PUT("/sqlalbums/:id", sql.SQLUpdateAlbumByID)    // PUT HTTP method with /sqlalbums/:id will be routed to SQLUpdateAlbumsByID function
	// router.DELETE("/sqlalbums/:id", sql.SQLDeleteAlbumByID) // DELETE HTTP method with /sqlalbums/:id will be routed to SQLDeleteAlbumsByID functio

	albumRepo := storage.NewAlbumSQL(db)
	// albumRepo := storage.NewAlbumMemory()
	albumSVCRepo := service.NewAlbumService(&albumRepo)
	albumHandler := handler.NewAlbumHandler(albumSVCRepo)
	// Album routes
	router.GET("/albums", albumHandler.GetAlbums)
	router.GET("/albums/:id", albumHandler.GetAlbumByID)       // GET HTTP method with /albums/id will be routed to getAlbumByID function
	router.POST("/albums", albumHandler.PostAlbums)            // POST HTTP method with /albums will be routed to getAlbums function
	router.PUT("/albums/:id", albumHandler.UpdateAlbumByID)    // PUT HTTP method with /albums/:id will be routed to updateAlbumsByID function
	router.DELETE("/albums/:id", albumHandler.DeleteAlbumByID) // DELETE HTTP method with /albums/:id will be routed to deleteAlbumsByID function

	return router
}
