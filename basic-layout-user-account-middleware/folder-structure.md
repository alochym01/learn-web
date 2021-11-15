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

## Handler

1. Entry point for User - `albumhandler.go`
    1. Receive user Request
    1. Response to user Request
1. Transform Data Object from user request and then pass to service handlers.

## Service Logic

1. Define all service interface - `service/albumservice.go`
1. Define AlbumService Handler - `service/albumservice.go`
    1. Doing business logic for all albums
    1. Transform Data Object response to Handlers
    1. Implement all service interface methods

## Storage folder

1. Define AlbumMemory struct - use for memory storage - `storage/albummemory.go`
1. Define AlbumSQL struct - use for SQL storage - `storage/albumsql.go`
1. Implement all methods in AlbumRepository interface

## Error Refactor

1. Define custom errors for API - `errs/app.go`
