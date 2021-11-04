package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/alochym01/learn-web/basic-layout-handler-service/domain"
)

type AlbumSQL struct {
	db *sql.DB
}

func NewAlbumSQL(DB *sql.DB) AlbumSQL {
	return AlbumSQL{
		db: DB,
	}
}

func (a AlbumSQL) FindAll() ([]domain.Album, error) {
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
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return nil, err
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
			// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
			return nil, err
		}
		albums = append(albums, a)
	}
	// c.IndentedJSON(http.StatusOK, albums)
	return albums, nil
}

// SQLGetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a AlbumSQL) FindByID(id string) (*domain.Album, error) {
	// DB, err1 := sql.Open("sqlite3", "alochym.db")

	// if err1 != nil {
	// 	panic(err1)
	// }

	// defer DB.Close()

	// id := c.Param("id")
	sqlstmt := "select * from albums where id=?"

	row := a.db.QueryRow(sqlstmt, id)
	album := domain.Album{}
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Record Not Found", err.Error())
			// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
			return nil, err
		}
		fmt.Println("Server Scanning Row Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return nil, err
	}

	// c.IndentedJSON(http.StatusOK, album)
	return &album, nil
}

// SQLPostAlbums adds an album from JSON received in the request body.
func (a AlbumSQL) Create(album domain.Album) error { // sqlstmt - Avoid SQL Injection Attack
	// DB, err1 := sql.Open("sqlite3", "alochym.db")

	// if err1 != nil {
	// 	panic(err1)
	// }

	// defer DB.Close()

	// var newAlbum domain.Album

	// // Call BindJSON to bind the received JSON to newAlbum.
	// if err := c.BindJSON(&newAlbum); err != nil {
	// 	return
	// }

	sqlstmt := fmt.Sprintf("INSERT INTO albums(id, title, artist, price) VALUES(\"%s\",\"%s\", \"%s\", %f)",
		album.ID,
		album.Title,
		album.Artist,
		album.Price,
	)

	fmt.Println(sqlstmt)

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return err
	}

	_, err = result.LastInsertId()
	// rowCount, err := result.LastInsertId()
	// rowCount, err := result.RowsAffected()

	// error check for LastInsertId function
	if err != nil {
		fmt.Println("Server Get RowID Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return err
	}

	// c.IndentedJSON(http.StatusCreated, gin.H{"message": rowCount})
	return nil
}

// Update locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a AlbumSQL) Update(album domain.Album) error {
	// DB, err1 := sql.Open("sqlite3", "alochym.db")

	// if err1 != nil {
	// 	panic(err1)
	// }

	// defer DB.Close()

	// id := c.Param("id")

	// var newAlbum domain.Album

	// newAlbum.ID = id
	// // Call BindJSON to bind the received JSON to newAlbum.
	// if err := c.BindJSON(&newAlbum); err != nil {
	// 	return
	// }

	sqlstmt := fmt.Sprintf("UPDATE albums SET title=\"%s\", artist=\"%s\", price=%f where id=\"%s\"",
		album.Title,
		album.Artist,
		album.Price,
		album.ID,
	)

	fmt.Println(sqlstmt)

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return err
	}

	rowCount, err := result.RowsAffected()

	// error check for RowsAffected function
	if err != nil {
		fmt.Println("Server Get RowsAffected Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return err
	}

	// there is no row found
	if rowCount == 0 {
		fmt.Println("Record Not Found")
		// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return errors.New("Record Not Found")
	}

	// c.IndentedJSON(http.StatusAccepted, gin.H{"message": newAlbum})
	return nil
}

// Delete locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a AlbumSQL) Delete(id string) error {
	// DB, err1 := sql.Open("sqlite3", "alochym.db")

	// if err1 != nil {
	// 	panic(err1)
	// }

	// defer DB.Close()

	// id := c.Param("id")

	sqlstmt := fmt.Sprintf("DELETE FROM albums where id=%s", id)

	fmt.Println(sqlstmt)

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return err
	}

	rowCount, err := result.RowsAffected()

	// error check for RowsAffected function
	if err != nil {
		fmt.Println("Server Get RowsAffected Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return err
	}

	// there is no row found
	if rowCount == 0 {
		fmt.Println("Record Not Found")
		// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return errors.New("Record Not Found")
	}

	// c.IndentedJSON(http.StatusAccepted, gin.H{"message": rowCount})
	return nil
}
