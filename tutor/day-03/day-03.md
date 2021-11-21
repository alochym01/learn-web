# Gin hello world - Day 03

## Implement Get an Album

1. We define the router `/albums/:id` with GET method

   ```go
   router.GET("/albums/:id", getAlbumsByID)
   ```

2. Implement a getAlbumByID which return an album

   ```go
   // getAlbumsByID locates the album whose ID value matches the id.
   func getAlbumsByID(c *gin.Context) {
   	id := c.Param("id")

   	// Loop over the list of albums, looking for
   	// an album whose ID value matches the parameter.
   	for _, a := range albums {
   		if a.ID == id {
   			c.IndentedJSON(http.StatusOK,  gin.H{"data": a})
   			return
   		}
   	}
   	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
   }
   ```

   1. Write a getAlbum function that takes a gin.Context parameter

      ```go
      func getAlbumByID(c *gin.Context) {...}
      ```

      1. `gin.Context` is the most important part of Gin.
      2. `gin.Context` carries request details, validates and serializes JSON, and more
   2. The Request URL `/albums/1`. We are using `c.Param("id")` to get id value in string type
   3. We lookup id value in albums
      1. Found we send response to User Request

         ```go
         if a.ID == id {
   			c.IndentedJSON(http.StatusOK, a)
   			return
   		}
         ```

      1. Not Found we send response to User Request

         ```go
         c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
         ```

   4. Call `Context.IndentedJSON`
      1. To serialize the struct into JSON
      2. Add it to the response
      3. First argument is the HTTP status code

## Using curl to test

1. curl localhost:8080/albums
2. curl localhost:8080/albums/1