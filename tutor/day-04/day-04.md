# Gin hello world - Day 04

## Implement Create Album

1. We define the router `/albums` with POST method

   ```go
   router.POST("/albums", postAlbums)
   ```

2. Implement a postAlbums which return an ID of album

   ```go
   // postAlbums adds an album from JSON received in the request body.
   func postAlbums(c *gin.Context) {
   	var newAlbum Album

   	// Call BindJSON to bind the received JSON to newAlbum.
   	if err := c.BindJSON(&newAlbum); err != nil {
         c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
   		return
   	}

   	// Add the new album to the slice.
   	albums = append(albums, newAlbum)
   	c.IndentedJSON(http.StatusCreated, gin.H{"data": newAlbum.ID})
   }
   ```

   1. Write a postAlbums function that takes a gin.Context parameter

      ```go
      func postAlbums(c *gin.Context) {...}
      ```

      1. `gin.Context` is the most important part of Gin.
      2. `gin.Context` carries request details, validates and serializes JSON, and more
   2. Declare a `newAlbum` and convert request data into `newAlbum` variable
   3. Add the new album to the slice.

         ```go
      	albums = append(albums, newAlbum)
      	c.IndentedJSON(http.StatusCreated, gin.H{"data": newAlbum.ID})
         ```

   4. Call `Context.IndentedJSON`
      1. To serialize the struct into JSON
      2. Add it to the response
      3. First argument is the HTTP status code

## Using curl to test

1. curl localhost:8080/albums
2. curl localhost:8080/albums/1
3. curl -XPOST -H "Content-Type: application/json" -d @createAlbum.json http://127.0.0.1:8080/albums
