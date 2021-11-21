# Gin hello world - Day 02

## Implement Get all Album

1. We define the router `/albums` with GET method

   ```go
   router.GET("/albums", getAlbums)
   ```

2. Create a global albums variable

   ```go
   // albums slice to seed record album data.
   var albums = []Album{
   	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
   	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
   	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
   }
   ```

3. Implement a getAlbums which return all albums

   ```go
   // getAlbums responds with the list of all albums as JSON.
   func getAlbums(c *gin.Context) {
   	c.IndentedJSON(http.StatusOK, gin.H{
   		"data": albums,
   	})
   }
   ```

   1. Write a getAlbums function that takes a gin.Context parameter

      ```go
      func getAlbums(c *gin.Context) {...}
      ```

      1. `gin.Context` is the most important part of Gin.
      2. `gin.Context` carries request details, validates and serializes JSON, and more
   2. Call `Context.IndentedJSON`
      1. To serialize the struct into JSON
      2. Add it to the response
      3. First argument is the HTTP status code

## Using curl to test

1. curl localhost:8080/albums