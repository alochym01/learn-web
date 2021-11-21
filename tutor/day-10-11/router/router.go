package router

import (
	"github.com/alochym01/web-w-gin/handler"
	"github.com/alochym01/web-w-gin/service"
	"github.com/alochym01/web-w-gin/storage/memory"
	"github.com/gin-gonic/gin"
)

// Router return a gin.Engine
func Router() *gin.Engine {
	router := gin.Default()

	// Album section - Start

	// New Album Storage
	storeAlbum := memory.NewAlbum()
	// New Album Service
	svcService := service.NewAlbumService(&storeAlbum)
	// New Album Handler
	aHandler := handler.NewAlbumHandler(svcService)

	// Get All Albums		- method GET, URL /albums
	router.GET("/albums", aHandler.GetAlbums)
	// Get An Albums 		- method GET, URL /albums/:id
	router.GET("/albums/:id", aHandler.GetAlbumsByID)
	// Get Create An Albums - method GET, URL /albums
	router.POST("/albums", aHandler.PostAlbums)
	// Get Update An Albums - method PUT, URL /albums/:id
	router.PUT("/albums/:id", aHandler.PutAlbums)
	// Get Delete An Albums - method DELETE, URL /albums/:id
	router.DELETE("/albums/:id", aHandler.DeleteAlbums)

	// Album section - End

	return router
}
