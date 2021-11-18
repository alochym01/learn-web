# Web with Go

## Tools
1. Golang version `go1.17.3`
2. VS Code
3. Gin framework
4. Auto reload go app - `github.com/cosmtrek/air`

## Day 01
1. Create basic REST server using go-gin - hello world
2. Create Album model
3. Storage backend is Memory

## Day 02
1. Implement a GET all albums
2. URL - `/albums`

## Day 03
1. Implement a GET an album
2. URL - `/albums/:id`

## Day 04
1. Implement a POST an album - Create an album
2. URL - `/albums`

## Day 05
1. Implement a DELETE an album - Delete an album
2. URL - `/albums/:id`

## Day 06
1. Implement a PUT an albums - Update an album
2. URL - `/albums/:id`

## Day 07
1. Organize code by using Router

## Day 08
1. Organize code by using album handler - `type AlbumHandler struct`

## Day 09
1. Organize code by using Storage Interface - `type AlbumRepository interface`
2. Implement all methods in `AlbumRepository interface`
3. Link `AlbumHandler` to `AlbumRepository interface` storage

## Day 10
1. Mapping Album request data into AlbumRequest Object - Data Transfer Object(DTO)
   1. Using `AlbumHandler`

## Day 11
1. Clean Architechture concept - Hexagonal Architechture
2. Organize code by using business layer - `type AlbumServiceRepository interface`
3. Implement all methods in business layer - `AlbumServiceRepository interface`
4. Link `AlbumHandler` to `AlbumServiceRepository interface` bunisness layer
5. Link Business Layer to `AlbumRepository interface` storage layer

## Day 12
1. Mapping Storage response to AlbumResponse Object - Data Transfer Object(DTO)
   1. Using `Service Layer`

## Day 13
1. Switch storage backend by using Database - SQLITE3
2. Create SQL Connection

## Day 14
1. Implement a GET all albums by using SQL
2. URL - `/albums`

## Day 15
1. Implement a GET an album by using SQL
2. URL - `/albums/:id`

## Day 16
1. Implement a POST an album - Create an album by using SQL
2. URL - `/albums`

## Day 17
1. Implement a DELETE an album - Delete an album by using SQL
2. URL - `/albums/:id`

## Day 18
1. Implement a PUT an albums - Update an album by using SQL
2. URL - `/albums/:id`

## Day 19
1. SQL Error handling
   1. Album NOT FOUND
   2. SQL Server
      1. Connection error
      2. Result error

## Day 20
1. Logger

## Day 21
1. User Model - `/users`
   1. Get all users
   1. Get a user
   2. Create a user

## Day 22
1. Simple Authentication Login

## Day 22
1. Response User login request with JWT token

## Day 23
1. Simple Authorization with JWT token - using middleware

## Day 24
1. Golang expvar Metrics

## Day 25
1. Validation Request Data