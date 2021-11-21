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

2. Create an `AlbumRepository Interface` in `domain/album.go`

   ```go
   package domain

   // AlbumRepository ...
   type AlbumRepository interface {
   	FindAll() ([]Album, error)
   	FindByID(int) (*Album, error)
   	Create(Album) (*int, error)
   	Update(Album) error
   	Delete(int) error
   }
   ```

3. Create `AlbumStorage Object` in [`storage/memory/album.go`](storage/memory/album.go) to implement all functions of `AlbumRepository Interface`

   ```go
   // Album ...
   type Album struct {
   	al []domain.Album
   }

   // NewAlbum ...
   func NewAlbum() Album {
   	albums := []domain.Album{
   		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
   		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
   		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
   	}

   	return Album{
   		al: albums,
   	}
   }

   func (a *Album) FindAll() ([]Album, error) {}
   func (a *Album) FindByID(int) (*Album, error) {}
   func (a *Album) Create(Album) (*int, error) {}
   func (a *Album) Update(Album) error {}
   func (a *Album) Delete(int) error {}
   ```

4. Create a `AlbumServiceRepository Interface` in `service/album.go`

   ```go
   // AlbumServiceRepository ...
   type AlbumServiceRepository interface {
   	GetAlbums() ([]domain.Album, error)
   	ByID(int) (*domain.Album, error)
   	Create(domain.AlbumRequest) (*int, error)
   	Update(int, domain.AlbumRequest) error
   	Delete(int) error
   }
   ```

5. Create AlbumService Object in `service/album.go` to implement all function of `AlbumServiceRepository Interface`

   ```go
   // AlbumService ...
   type AlbumService struct {
   	storageRepo domain.AlbumRepository
   }

   // NewAlbumService ...
   func NewAlbumService(repo domain.AlbumRepository) AlbumService {
   	return AlbumService{
   		storageRepo: repo,
   	}
   }

   func (svc AlbumService) GetAlbums() ([]domain.Album, error) {}
   func (svc AlbumService) ByID(int) (*domain.Album, error) {}
   func (svc AlbumService) Create(domain.AlbumRequest) (*int, error) {}
   func (svc AlbumService) Update(int, domain.AlbumRequest) error {}
   func (svc AlbumService) Delete(int) error {}
   ```

6. Refactor AlbumHandler using svc attribute to bind a an AlbumServiceRepository

   ```go
   // AlbumHandler ...
   type AlbumHandler struct {
   	svcRepo service.AlbumServiceRepository
   }

   // NewAlbumHandler ...
   func NewAlbumHandler(svc service.AlbumService) AlbumHandler {
   	return AlbumHandler{
   		svcRepo: svc,
   	}
   }
   ```

7. Putting all together in `router/router.go`

   ```go
   // New Album Storage
	storeAlbum := memory.NewAlbum()
	// New Album Service
	svcService := service.NewAlbumService(&storeAlbum) // because of using pointer **receiver
	// New Album Handler
	aHandler := handler.NewAlbumHandler(svcService)
   ```

## Using curl to test

1. curl localhost:8080/albums
2. curl localhost:8080/albums/1
3. curl -XPOST -H "Content-Type: application/json" -d @createAlbum.json http://127.0.0.1:8080/albums
4. curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/albums/1
5. curl -XPUT -H "Content-Type: application/json" -d @updateAlbum.json http://127.0.0.1:8080/albums/3
