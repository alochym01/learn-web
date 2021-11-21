# Gin hello world - Day 09

## Organize code by using Handler

1. We create folder structure

   ```bash
   .
   ├── domain
   │   └── album.go
   ├── go.mod
   ├── go.sum
   ├── handler
   │   └── album.go
   ├── main.go
   └── router
       └── router.go
   ```

2. Move all function to `domain/album.go`

   ```go
   package domain

   // Album represents data about a record Album.
   type Album struct {
   	ID     int
   	Title  string
   	Artist string
   	Price  float64
   }

   // AlbumRequest represents user request data.
   type AlbumRequest struct {
   	ID     int
   	Title  string
   	Artist string
   	Price  float64
   }
   ```

3. Import `domain` function into `handler/album.go` file

   ```go
   package router

   import (
   	"github.com/alochym01/web-w-gin/domain"
      "github.com/gin-gonic/gin"
      ...
   )

   // AlbumHandler ...
   type AlbumHandler struct {
   	al []domain.Album // using domain.Album
   }
   ```

## Using curl to test

1. curl localhost:8080/albums
2. curl localhost:8080/albums/1
3. curl -XPOST -H "Content-Type: application/json" -d @createAlbum.json http://127.0.0.1:8080/albums
4. curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/albums/1
5. curl -XPUT -H "Content-Type: application/json" -d @updateAlbum.json http://127.0.0.1:8080/albums/3
