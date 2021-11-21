# Gin hello world - Day 06

## Implement Update Album

1. We define the router `/albums/:id` with PUT method

   ```go
   router.PUT("/albums/:id", putAlbums)
   ```

2. Implement a putAlbums which return an ID of album

   ```go
   // putAlbums locates the album whose ID value matches the id
   // parameter sent by the client, then returns that album as a response.
   func putAlbums(c *gin.Context) {
      	id, err := strconv.Atoi(c.Param("id"))

      	if err != nil {
      		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
      		return
		}

        var updateAlbum Album

        // Call BindJSON to bind the received JSON to updateAlbum.
        if err := c.BindJSON(&updateAlbum); err != nil {
        	c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
        	return
		}

		updateAlbum.ID = id

      	// Loop over the list of albums, looking for
      	// an album whose ID value matches the parameter.
      	for i, a := range albums {
      		if a.ID == id {
      			albums[i] = updateAlbum
      			c.IndentedJSON(http.StatusOK, gin.H{"data": "album updated"})
      			return
      		}
      	}
      	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
   }
   ```

   1. Write a putAlbums function that takes a gin.Context parameter

      ```go
      func putAlbums(c *gin.Context) {...}
      ```

      1. `gin.Context` is the most important part of Gin.
      2. `gin.Context` carries request details, validates and serializes JSON, and more
   2. The Request URL `/albums/1`. We are using `c.Param("id")` to get id value in string type
   3. We lookup id value in albums
      1. Found we update an album

         ```go
      	 // Loop over the list of albums, looking for
      	 // an album whose ID value matches the parameter.
      	 for i, a := range albums {
      	 	if a.ID == id {
      	 		albums[i] = updateAlbum
      	 		c.IndentedJSON(http.StatusOK, gin.H{"data": "album updated"})
      	 		return
      	 	}
      	 }
         ```

      2. Not Found we send response to User Request

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
3. curl -XPOST -H "Content-Type: application/json" -d @createAlbum.json http://127.0.0.1:8080/albums
4. curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/albums/1
5. curl -XPUT -H "Content-Type: application/json" -d @updateAlbum.json http://127.0.0.1:8080/albums/3
