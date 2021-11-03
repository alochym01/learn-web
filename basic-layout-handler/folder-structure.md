# Basic Layout

## Entry Point folder

1. `web/api/main.go` - using for web api
1. `web/cmd/main.go` - using for cli command

## Router folder

1. `router/router.go` - define all route for web api application

## Domain folder

1. `domain/album.go`
    1. define all Album table and Album functions
    1. define Album interface methods

        ```go
        type AlbumRepository interface {
            GetAlbums(c *gin.Context)
            GetAlbumByID(c *gin.Context)
            PostAlbums(c *gin.Context)
            UpdateAlbumByID(c *gin.Context)
            DeleteAlbumByID(c *gin.Context)
        }
        ```

## Storage folder

1. Define Memory Storage struct - use for memory storage `storage/albummemory.go`
    1. Implement all Album interface

        ```go
        // GetAlbums responds with the list of all albums as JSON.
        func (a *AlbumMemory) GetAlbums(c *gin.Context) {}

        // GetAlbumByID locates the album whose ID value matches the id
        // parameter sent by the client, then returns that album as a response.
        func (a *AlbumMemory) GetAlbumByID(c *gin.Context) {}

        // PostAlbums adds an album from JSON received in the request body.
        func (a *AlbumMemory) PostAlbums(c *gin.Context) {}

        // UpdateAlbumByID locates the album whose ID value matches the id
        // parameter sent by the client, then returns that album as a response.
        func (a *AlbumMemory) UpdateAlbumByID(c *gin.Context) {}

        // DeleteAlbumByID locates the album whose ID value matches the id
        // parameter sent by the client, then returns that album as a response.
        func (a *AlbumMemory) DeleteAlbumByID(c *gin.Context) {}
        ```

    1. Implement function to create AlbumMemory instance

        ```go
        func NewAlbumMemory() AlbumMemory {
            albums := []domain.Album{
                {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
                {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
                {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
            }

            return AlbumMemory{
                al: albums,
            }
        }
        ```

1. Define SQL Storage struct - use for SQL storage
