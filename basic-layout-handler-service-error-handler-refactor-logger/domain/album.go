package domain

import (
	"github.com/alochym01/learn-web/basic-layout-handler-service-error-handler-refactor-logger/errs"
	_ "github.com/mattn/go-sqlite3"
)

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Album - define Album table with its attribute
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type AlbumRepository interface {
	FindAll() ([]Album, *errs.Error)
	FindByID(id string) (*Album, *errs.Error)
	Create(a Album) *errs.Error
	Update(a Album) *errs.Error
	Delete(string) *errs.Error
}

type AlbumRequest struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// // GetAlbums responds with the list of all albums as JSON.
// func GetAlbums(c *gin.Context) {

// 	c.IndentedJSON(http.StatusOK, albums)
// }

// // PostAlbums adds an album from JSON received in the request body.
// func PostAlbums(c *gin.Context) {
// 	var newAlbum Album

// 	// Call BindJSON to bind the received JSON to newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// // GetAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func GetAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

// // UpdateAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func UpdateAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for i, a := range albums {
// 		if a.ID == id {
// 			var updateAlbum Album

// 			// Call BindJSON to bind the received JSON to updateAlbum.
// 			if err := c.BindJSON(&updateAlbum); err != nil {
// 				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
// 				return
// 			}
// 			a.Title = updateAlbum.Title
// 			a.Artist = updateAlbum.Artist
// 			a.Price = updateAlbum.Price
// 			albums[i] = a
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

// // DeleteAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func DeleteAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for i, a := range albums {
// 		if a.ID == id {
// 			tempAlbum := make([]Album, len(albums)-1)
// 			copy(tempAlbum[:i], albums[:i])
// 			copy(tempAlbum[i:], albums[i+1:])
// 			albums = tempAlbum
// 			c.IndentedJSON(http.StatusOK, albums)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

/*
 SQL
*/
// func SQLGetAlbums(c *gin.Context) {
// 	DB, err := sql.Open("sqlite3", "alochym.db")

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer DB.Close()

// 	sqlstmt := "select * from albums"

// 	rows, err := DB.Query(sqlstmt)

// 	// check err from server DB and Query DB
// 	if err != nil {
// 		fmt.Println("Server Database Error", err.Error())
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
// 		return
// 	}

// 	albums := []Album{}

// 	for rows.Next() {
// 		a := Album{}
// 		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
// 		// check err from server DB and Scan function
// 		if err != nil {
// 			// if err == sql.ErrNoRows {
// 			// 	fmt.Println("Record Not Found", err.Error())
// 			// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
// 			// 	return
// 			// }
// 			fmt.Println("Server Scanning Rows Error", err.Error())
// 			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
// 			return
// 		}
// 		albums = append(albums, a)
// 	}
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// // SQLGetAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func SQLGetAlbumByID(c *gin.Context) {
// 	DB, err1 := sql.Open("sqlite3", "alochym.db")

// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	defer DB.Close()

// 	id := c.Param("id")
// 	sqlstmt := "select * from albums where id=?"

// 	row := DB.QueryRow(sqlstmt, id)
// 	a := Album{}
// 	err := row.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("Record Not Found", err.Error())
// 			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
// 			return
// 		}
// 		fmt.Println("Server Scanning Row Error", err.Error())
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
// 	}

// 	c.IndentedJSON(http.StatusOK, a)
// }

// // SQLPostAlbums adds an album from JSON received in the request body.
// func SQLPostAlbums(c *gin.Context) { // sqlstmt - Avoid SQL Injection Attack
// 	DB, err1 := sql.Open("sqlite3", "alochym.db")

// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	defer DB.Close()

// 	var newAlbum Album

// 	// Call BindJSON to bind the received JSON to newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	sqlstmt := fmt.Sprintf("INSERT INTO albums(id, title, artist, price) VALUES(\"%s\",\"%s\", \"%s\", %f)",
// 		newAlbum.ID,
// 		newAlbum.Title,
// 		newAlbum.Artist,
// 		newAlbum.Price,
// 	)

// 	fmt.Println(sqlstmt)

// 	// Execute SQL Statements
// 	result, err := DB.Exec(sqlstmt)

// 	// check err from server DB and Query DB
// 	if err != nil {
// 		fmt.Println("Server Database Error", err.Error())
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
// 		return
// 	}

// 	rowCount, err := result.LastInsertId()
// 	// rowCount, err := result.RowsAffected()

// 	// error check for LastInsertId function
// 	if err != nil {
// 		fmt.Println("Server Get RowID Error", err.Error())
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusCreated, gin.H{"message": rowCount})
// }

// // SQLUpdateAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func SQLUpdateAlbumByID(c *gin.Context) {
// 	DB, err1 := sql.Open("sqlite3", "alochym.db")

// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	defer DB.Close()

// 	id := c.Param("id")

// 	var newAlbum Album

// 	newAlbum.ID = id
// 	// Call BindJSON to bind the received JSON to newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	sqlstmt := fmt.Sprintf("UPDATE albums SET title=\"%s\", artist=\"%s\", price=%f where id=\"%s\"",
// 		newAlbum.Title,
// 		newAlbum.Artist,
// 		newAlbum.Price,
// 		newAlbum.ID,
// 	)

// 	fmt.Println(sqlstmt)

// 	// Execute SQL Statements
// 	result, err := DB.Exec(sqlstmt)

// 	// check err from server DB and Query DB
// 	if err != nil {
// 		fmt.Println("Server Database Error", err.Error())
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
// 		return
// 	}

// 	rowCount, err := result.RowsAffected()

// 	// error check for RowsAffected function
// 	if err != nil {
// 		fmt.Println("Server Get RowsAffected Error", err.Error())
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
// 		return
// 	}

// 	// there is no row found
// 	if rowCount == 0 {
// 		fmt.Println("Record Not Found")
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusAccepted, gin.H{"message": newAlbum})
// }

// // SQLDeleteAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func SQLDeleteAlbumByID(c *gin.Context) {
// 	DB, err1 := sql.Open("sqlite3", "alochym.db")

// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	defer DB.Close()

// 	id := c.Param("id")

// 	sqlstmt := fmt.Sprintf("DELETE FROM albums where id=%s", id)

// 	fmt.Println(sqlstmt)

// 	// Execute SQL Statements
// 	result, err := DB.Exec(sqlstmt)

// 	// check err from server DB and Query DB
// 	if err != nil {
// 		fmt.Println("Server Database Error", err.Error())
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
// 		return
// 	}

// 	rowCount, err := result.RowsAffected()

// 	// error check for RowsAffected function
// 	if err != nil {
// 		fmt.Println("Server Get RowsAffected Error", err.Error())
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
// 		return
// 	}

// 	// there is no row found
// 	if rowCount == 0 {
// 		fmt.Println("Record Not Found")
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusAccepted, gin.H{"message": rowCount})
// }
