# Gin hello world - Day 10

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
   ├── service
   │   └── album.go
   ├── storage
   │   └── memory
   │       └── album.go
   └── router
       └── router.go
   ```

2. Create a [`storage/sqlite3/album.go`](storage/sqlite3/album.go) file to implement all functions of `AlbumRepository Interface`

   ```go
   // Album ...
   type Album struct {
   	db *sql.DB
   }

   // NewAlbum ...
   func NewAlbum(DB *sql.DB) Album {
   	return Album{
   		db: DB,
   	}
   }

   func (a *Album) FindAll() ([]Album, error) {}
   func (a *Album) FindByID(int) (*Album, error) {}
   func (a *Album) Create(Album) (*int64, error) {} // change func (a *Album) Create(Album) (*int, error) {}
   func (a *Album) Update(Album) error {}
   func (a *Album) Delete(int) error {}
   ```

3. Change `AlbumRepository Interface`

   ```go
   // AlbumRepository ...
   type AlbumRepository interface {
   	FindAll() ([]Album, error)
   	FindByID(int) (*Album, error)
   	Create(Album) (*int64, error) // Create(Album) (*int, error)
   	Update(Album) error
   	Delete(int) error
   }
   ```

4. Putting all together in `router/router.go`

   ```go
	// New Album Storage - SQLite
	storeAlbum := sqlite3.NewAlbum(db)
   ```

5. Update `main.go` with sql object

   ```go
   package main

   import (
   	"database/sql"

   	"github.com/alochym01/web-w-gin/router"

   	_ "github.com/mattn/go-sqlite3"
   )

   func main() {
   	db, err1 := sql.Open("sqlite3", "alochym.db")

   	if err1 != nil {
   		panic(err1)
   	}

   	defer db.Close()

   	r := router.Router(db)
   	r.Run("localhost:8080")
   }
   ```

## Using curl to test

1. curl localhost:8080/albums
2. curl localhost:8080/albums/1
3. curl -XPOST -H "Content-Type: application/json" -d @createAlbum.json http://127.0.0.1:8080/albums
4. curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/albums/1
5. curl -XPUT -H "Content-Type: application/json" -d @updateAlbum.json http://127.0.0.1:8080/albums/3
