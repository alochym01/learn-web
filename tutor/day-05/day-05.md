# Gin hello world - Day 05

## Implement Delete Album

1. We define the router `/albums` with DELETE method

   ```go
   router.DELETE("/albums", deleteAlbums)
   ```

2. Implement a deleteAlbums which return an ID of album

   ```go
   // deleteAlbums locates the album whose ID value matches the id
   // parameter sent by the client, then returns that album as a response.
   func deleteAlbums(c *gin.Context) {
   	id, err := strconv.Atoi(c.Param("id"))

   	if err != nil {
   		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": "Bad Request"})
   		return
   	}

   	// Loop over the list of albums, looking for
   	// an album whose ID value matches the parameter.
   	for i, a := range albums {
   		if a.ID == id {
   			// Create a tempAlbum's element = Album's element - 1
   	 		tempAlbum := make([]Album, len(albums)-1)
   	 		// Copy Album element from index [0 -> i] to tempAlbum
   	 		copy(tempAlbum, albums[:i])
   	 		// Copy Album element from index [i + 1 -> end] to tempAlbum
   	 		copy(tempAlbum, albums[i+1:])
   	 		// Assign albums to tempAlbum
   	 		albums = tempAlbum
   	 		c.IndentedJSON(http.StatusAccepted, gin.H{"data": "albums delete  successfull"})
   			return
   		}
   	}
   	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "album not found"})
   }
   ```

   1. Write a postAlbums function that takes a gin.Context parameter

      ```go
      func deleteAlbums(c *gin.Context) {...}
      ```

      1. `gin.Context` is the most important part of Gin.
      2. `gin.Context` carries request details, validates and serializes JSON, and more
   2. The Request URL `/albums/1`. We are using `c.Param("id")` to get id value in string type
   3. We lookup id value in albums
      1. Found we delete an album

         ```go
      	 // Loop over the list of albums, looking for
      	 // an album whose ID value matches the parameter.
      	 for i, a := range albums {
      	 	if a.ID == id {
      	 		// Create a tempAlbum's element = Album's element - 1
      	 		tempAlbum := make([]Album, len(albums)-1)
      	 		// Copy Album element from index [0 -> i] to tempAlbum
      	 		copy(tempAlbum, albums[:i])
      	 		// Copy Album element from index [i + 1 -> end] to tempAlbum
      	 		copy(tempAlbum, albums[i+1:])
      	 		// Assign albums to tempAlbum
      	 		albums = tempAlbum
      	 		c.IndentedJSON(http.StatusAccepted, gin.H{"data": "albums delete  successfull"})
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
