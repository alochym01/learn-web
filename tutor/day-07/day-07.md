# Gin hello world - Day 07

## Organize code by using Router

1. We create folder structure
   ```bash
   .
   ├── go.mod
   ├── go.sum
   ├── main.go
   └── router
       └── router.go
   ```
2. Move all function to `router.go`

   ```go
   // Album represents data about a record Album.
   type Album struct {...}

   // albums slice to seed record album data.
   var albums = []Album{...}

   // Router return a gin.Engine
   func Router() *gin.Engine {
   	router := gin.Default()

   	// Album section - Start

   	// Get All Albums		- method GET, URL /albums
   	router.GET("/albums", getAlbums)
   	// Get An Albums 		- method GET, URL /albums/:id
   	router.GET("/albums/:id", getAlbumsByID)
   	// Get Create An Albums - method GET, URL /albums
   	router.POST("/albums", postAlbums)
   	// Get Update An Albums - method PUT, URL /albums/:id
   	router.PUT("/albums/:id", putAlbums)
   	// Get Delete An Albums - method DELETE, URL /albums/:id
   	router.DELETE("/albums/:id", deleteAlbums)

   	// Album section - End

   	return router
   }

   // getAlbums responds with the list of all albums as JSON.
   func getAlbums(c *gin.Context) {...}

   // getAlbumsByID locates the album whose ID value matches the id
   func getAlbumsByID(c *gin.Context) {...}

   // postAlbums adds an album from JSON received in the request body.
   func postAlbums(c *gin.Context) {...}

   // deleteAlbums locates the album whose ID value matches the id
   func deleteAlbums(c *gin.Context) {...}

   // putAlbums locates the album whose ID value matches the id
   func putAlbums(c *gin.Context) {...}
   ```

3. Import `Router` function into `main.go` file

   ```go
   package main

   import "github.com/alochym01/web-w-gin/router"

   func main() {
   	r := router.Router()
   	r.Run("localhost:8080")
   }
   ```

## Using curl to test

1. curl localhost:8080/albums
2. curl localhost:8080/albums/1
3. curl -XPOST -H "Content-Type: application/json" -d @createAlbum.json http://127.0.0.1:8080/albums
4. curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/albums/1
5. curl -XPUT -H "Content-Type: application/json" -d @updateAlbum.json http://127.0.0.1:8080/albums/3
