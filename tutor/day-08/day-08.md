# Gin hello world - Day 08

## Organize code by using Handler

1. We create folder structure

   ```bash
   .
   ├── go.mod
   ├── go.sum
   ├── handler
   │   └── album.go
   ├── main.go
   └── router
       └── router.go
   ```

2. Move all function to `album.go`

   ```go
   package handler

   // Album represents data about a record Album.
   type Album struct {...}

   // AlbumHandler ...
   type AlbumHandlder struct {
   	al []Album
   }

   // NewAlbumHandler ...
   func NewAlbumHandler() AlbumHandlder {
   	albums := []Album{
   		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
   		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
   		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
   	}

   	return AlbumHandlder{
   		al: albums,
   	}
   }

   // GetAlbums responds with the list of all albums as JSON.
   func (a *AlbumHandlder) GetAlbums(c *gin.Context) {...}

   // GetAlbumsByID locates the album whose ID value matches the id
   func (a *AlbumHandlder) GetAlbumsByID(c *gin.Context) {...}

   // PostAlbums adds an album from JSON received in the request body.
   func  (a *AlbumHandlder) PostAlbums(c *gin.Context) {...}

   // DeleteAlbums locates the album whose ID value matches the id
   func (a *AlbumHandlder) DeleteAlbums(c *gin.Context) {...}

   // PutAlbums locates the album whose ID value matches the id
   func (a *AlbumHandlder) PutAlbums(c *gin.Context) {...}
   ```

3. Import `handler` function into `router/router.go` file

   ```go
   package router

   import (
   	"github.com/alochym01/web-w-gin/handler"
   	"github.com/gin-gonic/gin"
   )

   // Router return a gin.Engine
   func Router() *gin.Engine {
   	router := gin.Default()

   	// Album section - Start

   	// New Album Handler
   	aHandler := handler.NewAlbumHandler()

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
   ```

## Using curl to test

1. curl localhost:8080/albums
2. curl localhost:8080/albums/1
3. curl -XPOST -H "Content-Type: application/json" -d @createAlbum.json http://127.0.0.1:8080/albums
4. curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/albums/1
5. curl -XPUT -H "Content-Type: application/json" -d @updateAlbum.json http://127.0.0.1:8080/albums/3
