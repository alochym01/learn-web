# Using Go-GIN

1. Design API endpoints.
1. Create a folder for your code.
1. Create the data.
1. Write a handler to return all items.
1. Write a handler to add a new item.
1. Write a handler to return a specific item.

## How to Code

1. Define Album table

    ```go
    type album struct {
        ID     string  `json:"id"`
        Title  string  `json:"title"`
        Artist string  `json:"artist"`
        Price  float64 `json:"price"`
    }
    ```

1. Get all the albums as JSON -- Create function getAlbums

    ```go
    // getAlbums responds with the list of all albums as JSON.
    func getAlbums(c *gin.Context) {
        c.IndentedJSON(http.StatusOK, albums)
    }
    ```

    1. Write a getAlbums function that takes a `gin.Context` parameter.
        1. `gin.Context` is the most important part of Gin
    1. Call `Context.IndentedJSON` to serialize the struct into JSON and add it to the response.
    1. We are passing the `StatusOK` constant from the `net/http` package to indicate 200 OK
1. Add code to add albums data to the list of albums -- Create a new Album

    ```go
    // postAlbums adds an album from JSON received in the request body.
    func postAlbums(c *gin.Context) {
        var newAlbum album

        // Call BindJSON to bind the received JSON to newAlbum.
        if err := c.BindJSON(&newAlbum); err != nil {
            return
        }

        // Add the new album to the slice.
        albums = append(albums, newAlbum)
        c.IndentedJSON(http.StatusCreated, newAlbum)
    }
    ```

1. Add logic to retrieve the requested album.

    ```go
    // getAlbumByID locates the album whose ID value matches the id
    // parameter sent by the client, then returns that album as a response.
    func getAlbumByID(c *gin.Context) {
        id := c.Param("id")

        // Loop over the list of albums, looking for
        // an album whose ID value matches the parameter.
        for _, a := range albums {
            if a.ID == id {
                c.IndentedJSON(http.StatusOK, a)
                return
            }
        }
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
    }
    ```

1. Add logic to update the album.

    ```go
    // updateAlbumByID locates the album whose ID value matches the id
    // parameter sent by the client, then returns that album as a response.
    func updateAlbumByID(c *gin.Context) {
        id := c.Param("id")

        // Loop over the list of albums, looking for
        // an album whose ID value matches the parameter.
        for i, a := range albums {
            if a.ID == id {
                var updateAlbum Album

                // Call BindJSON to bind the received JSON to updateAlbum.
                if err := c.BindJSON(&updateAlbum); err != nil {
                    c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
                    return
                }
                a.Title = updateAlbum.Title
                a.Artist = updateAlbum.Artist
                a.Price = updateAlbum.Price
                albums[i] = a
                c.IndentedJSON(http.StatusOK, a)
                return
            }
        }
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
    }
    ```

1. Add logic to Delete the album.

    ```go
    // deleteAlbumByID locates the album whose ID value matches the id
    // parameter sent by the client, then returns that album as a response.
    func deleteAlbumByID(c *gin.Context) {
        id := c.Param("id")

        // Loop over the list of albums, looking for
        // an album whose ID value matches the parameter.
        for i, a := range albums {
            if a.ID == id {
                tempAlbum := make([]Album, len(albums)-1)
                copy(tempAlbum, albums[:i])
                copy(tempAlbum, albums[i+1:])
                albums = tempAlbum
                c.IndentedJSON(http.StatusOK, albums)
                return
            }
        }
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
    }
    ```

1. Create function main
    1. Create GIN Router - Initialize a Gin router using Default.
    1. Create a router `/albums` with GET method => getAblums function

        ```go
        var albums = []album{
            {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
            {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
            {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
        }

        func main() {
            router := gin.Default() // Initialize a Gin router using Default.
            router.GET("/albums", getAlbums) // GET HTTP method with /albums wil be routed to getAlbums function
            router.GET("/albums/:id", getAlbumByID) // GET HTTP method with /albums/:id wil be routed to getAlbumByID function
            router.POST("/albums", postAlbums) // POST HTTP method with /albums wil be routed to postAlbums function
            router.PUT("/albums/:id", updateAlbumByID) // PUT HTTP method with /albums/:id wil be routed to updateAlbums function
            router.DELETE("/albums/:id", deleteAlbumByID) // PUT HTTP method with /albums/:id wil be routed to deleteAlbumsByID function

            router.Run("localhost:8080") // Run function to attach the router to an http.Server and start the server.
        }
        ```

1. Run the code

    ```go
    go run main.go
    ```

1. Using curl to test `curl http://8080/albums`
