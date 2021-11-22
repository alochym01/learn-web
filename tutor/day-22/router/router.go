package router

import (
	"database/sql"
	"time"

	"github.com/alochym01/web-w-gin/handler"
	"github.com/alochym01/web-w-gin/logger"
	"github.com/alochym01/web-w-gin/service"
	"github.com/alochym01/web-w-gin/storage/sqlite3"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// Router return a gin.Engine
func Router(db *sql.DB) *gin.Engine {
	// router := gin.Default()
	router := gin.New()

	// Using ZAP Logger
	logger := logger.NewLogger()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(logger, true))

	// Album section - Start

	// New Album Storage - Memory
	// storeAlbum := memory.NewAlbum()

	// New Album Storage - SQLite
	storeAlbum := sqlite3.NewAlbum(db)
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

	// User section - Start

	userRepo := sqlite3.NewUserSQL(db, logger)
	// userRepo := storage.NewAlbumMemory()
	userSVCRepo := service.NewUserService(&userRepo)
	userHandler := handler.NewUserHandler(userSVCRepo)
	// Album routes
	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetUserByID) // GET HTTP method with /Users/id will be routed to getUserByID function
	router.POST("/users", userHandler.PostUser)       // POST HTTP method with /Users will be routed to getUsers function
	// router.PUT("/users/:id", userHandler.UpdateUserByID)    // PUT HTTP method with /Users/:id will be routed to updateUsersByID function
	// router.DELETE("/users/:id", userHandler.DeleteUserByID) // DELETE HTTP method with /albums/:id will be routed to deleteAlbumsByID function

	// User section - End

	return router
}
