package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/alochym01/learn-web/basic-layout-handler-service/domain"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type AlbumSQL struct {
	db *sql.DB
}

func NewAlbumSQL(DB *sql.DB) AlbumSQL {
	return AlbumSQL{
		db: DB,
	}
}

func (a AlbumSQL) SQLGetAlbums(c *gin.Context) {
	// DB, err := sql.Open("sqlite3", "alochym.db")

	// if err != nil {
	// 	panic(err)
	// }

	// defer DB.Close()

	sqlstmt := "select * from albums"

	rows, err := a.db.Query(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}

	albums := []domain.Album{}

	for rows.Next() {
		a := domain.Album{}
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		// check err from server DB and Scan function
		if err != nil {
			// if err == sql.ErrNoRows {
			// 	fmt.Println("Record Not Found", err.Error())
			// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
			// 	return
			// }
			fmt.Println("Server Scanning Rows Error", err.Error())
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
			return
		}
		albums = append(albums, a)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// SQLGetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a AlbumSQL) SQLGetAlbumByID(c *gin.Context) {
	// DB, err1 := sql.Open("sqlite3", "alochym.db")

	// if err1 != nil {
	// 	panic(err1)
	// }

	// defer DB.Close()

	id := c.Param("id")
	sqlstmt := "select * from albums where id=?"

	row := a.db.QueryRow(sqlstmt, id)
	album := domain.Album{}
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Record Not Found", err.Error())
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
			return
		}
		fmt.Println("Server Scanning Row Error", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
	}

	c.IndentedJSON(http.StatusOK, album)
}

// SQLPostAlbums adds an album from JSON received in the request body.
func (a AlbumSQL) SQLPostAlbums(c *gin.Context) { // sqlstmt - Avoid SQL Injection Attack
	// DB, err1 := sql.Open("sqlite3", "alochym.db")

	// if err1 != nil {
	// 	panic(err1)
	// }

	// defer DB.Close()

	var newAlbum domain.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	sqlstmt := fmt.Sprintf("INSERT INTO albums(id, title, artist, price) VALUES(\"%s\",\"%s\", \"%s\", %f)",
		newAlbum.ID,
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Price,
	)

	fmt.Println(sqlstmt)

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}

	rowCount, err := result.LastInsertId()
	// rowCount, err := result.RowsAffected()

	// error check for LastInsertId function
	if err != nil {
		fmt.Println("Server Get RowID Error", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": rowCount})
}

// SQLUpdateAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a AlbumSQL) SQLUpdateAlbumByID(c *gin.Context) {
	// DB, err1 := sql.Open("sqlite3", "alochym.db")

	// if err1 != nil {
	// 	panic(err1)
	// }

	// defer DB.Close()

	id := c.Param("id")

	var newAlbum domain.Album

	newAlbum.ID = id
	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	sqlstmt := fmt.Sprintf("UPDATE albums SET title=\"%s\", artist=\"%s\", price=%f where id=\"%s\"",
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Price,
		newAlbum.ID,
	)

	fmt.Println(sqlstmt)

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}

	rowCount, err := result.RowsAffected()

	// error check for RowsAffected function
	if err != nil {
		fmt.Println("Server Get RowsAffected Error", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}

	// there is no row found
	if rowCount == 0 {
		fmt.Println("Record Not Found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": newAlbum})
}

// SQLDeleteAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a AlbumSQL) SQLDeleteAlbumByID(c *gin.Context) {
	// DB, err1 := sql.Open("sqlite3", "alochym.db")

	// if err1 != nil {
	// 	panic(err1)
	// }

	// defer DB.Close()

	id := c.Param("id")

	sqlstmt := fmt.Sprintf("DELETE FROM albums where id=%s", id)

	fmt.Println(sqlstmt)

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}

	rowCount, err := result.RowsAffected()

	// error check for RowsAffected function
	if err != nil {
		fmt.Println("Server Get RowsAffected Error", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return
	}

	// there is no row found
	if rowCount == 0 {
		fmt.Println("Record Not Found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": rowCount})
}
