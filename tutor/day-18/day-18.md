# Gin hello world - Day 18

## Organize code by using Handler

1. We create folder structure

   ```bash
   .
   ├── domain
   │   └── album.go
   ├── errs
   │   └── app.go
   ├── go.mod
   ├── go.sum
   ├── handler
   │   └── album.go
   ├── main.go
   ├── service
   │   └── album.go
   ├── storage
   │   └── memory
   │       └── album.go
   └── router
       └── router.go
   ```

2. Create a [`errs/app.go`](errs/app.go) file to handle error

   ```go
   package errs

   // AppErrs ...
   type AppErr struct {
      Code    int
      Message string
   }

   // NotFound ...
   func NotFound() *AppErr {
      return AppErr {
         Code: http.StatusNotFound,
         Message: "Not Found"
      }
   }

   // ServerError ...
   func ServerError() *AppErr {
      return &AppErr{
         Code:    http.StatusInternalServerError,
         Message: "Internal Server Error",
      }
   }
   ```

3. Change `AlbumRepository Interface`

   ```go
   import "github.com/alochym01/web-w-gin/errs"

   // AlbumRepository ...
   type AlbumRepository interface {
       FindAll() ([]Album, error)
       FindByID(int) (*Album, *errs.AppErr)
       Create(Album) (*int64, error) // Create(Album) (*int, error)
       Update(Album) error
       Delete(int) error
   }
   ```

4. Change `AlbumServiceRepository Interface`

   ```go
   import (
      ...
      "github.com/alochym01/web-w-gin/errs"
   )

   // AlbumServiceRepository ...
   type AlbumServiceRepository interface {
      GetAlbums() ([]domain.Album, error)
      ByID(int) (*domain.Album, *errs.AppErr) // using errs.AppErr
      Create(domain.AlbumRequest) (*int64, error)
      Update(int, domain.AlbumRequest) error
      Delete(int) error
   }
   ```

5. Update `main.go` with sql object

   ```go
   // parameter sent by the client, then returns that album as a response.
   func (a *AlbumHandler) GetAlbumsByID(c *gin.Context) {
      id, err := strconv.Atoi(c.Param("id"))

      if err != nil {
         c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
         return
      }

      // Get Album by using service
      album, errs := a.svcRepo.ByID(id)

      if errs != nil {
         c.IndentedJSON(errs.Code, gin.H{"data": errs.Message}) // using errs.AppErr
         return
      }

      c.IndentedJSON(http.StatusOK, gin.H{"data": album})
   }
   ```

## Using curl to test

1. curl localhost:8080/albums
2. curl localhost:8080/albums/1
3. curl -XPOST -H "Content-Type: application/json" -d @createAlbum.json http://127.0.0.1:8080/albums
4. curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/albums/1
5. curl -XPUT -H "Content-Type: application/json" -d @updateAlbum.json http://127.0.0.1:8080/albums/3
