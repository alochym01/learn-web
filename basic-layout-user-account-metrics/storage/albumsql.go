package storage

import (
	"database/sql"

	"github.com/alochym01/learn-web/basic-layout-user-account/domain"
	"github.com/alochym01/learn-web/basic-layout-user-account/errs"

	"go.uber.org/zap"
)

type AlbumSQL struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewAlbumSQL(DB *sql.DB, l *zap.Logger) AlbumSQL {
	return AlbumSQL{
		db:     DB,
		logger: l,
	}
}

func (a AlbumSQL) FindAll() ([]domain.Album, *errs.Error) {
	// DB, err := sql.Open("sqlite3", "alochym.db")

	// if err != nil {
	// 	panic(err)
	// }

	// defer DB.Close()

	sqlstmt := "select * from albums"

	rows, err := a.db.Query(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		// fmt.Println("Server Database Error", err.Error())
		a.logger.Error(err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return nil, errs.ServerError("Server Internal Error")
	}

	albums := []domain.Album{}

	for rows.Next() {
		album := domain.Album{}
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		// check err from server DB and Scan function
		if err != nil {
			// if err == sql.ErrNoRows {
			// 	fmt.Println("Record Not Found", err.Error())
			// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
			// 	return
			// }
			// fmt.Println("Server Scanning Rows Error", err.Error())
			a.logger.Error(err.Error())
			// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
			return nil, errs.ServerError("Server Internal Error")
		}
		albums = append(albums, album)
	}
	// c.IndentedJSON(http.StatusOK, albums)
	return albums, nil
}

// FindByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a AlbumSQL) FindByID(id string) (*domain.Album, *errs.Error) {
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
			// fmt.Println("Record Not Found", err.Error())
			a.logger.Error(err.Error())
			// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
			return nil, errs.NotFound("Record Not Found")
		}
		// fmt.Println("Server Scanning Row Error", err.Error())
		a.logger.Error(err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return nil, errs.ServerError("Server Internal Error")
	}

	// c.IndentedJSON(http.StatusOK, album)
	return &album, nil
}

// SQLPostAlbums adds an album from JSON received in the request body.
func (a AlbumSQL) Create(album domain.Album) *errs.Error { // sqlstmt - Avoid SQL Injection Attack
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

	sqlstmt := "INSERT INTO albums(id, title, artist, price) VALUES(?,?,?,?)"

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt, album.ID, album.Title, album.Artist, album.Price)

	// check err from server DB and Query DB
	if err != nil {
		a.logger.Error(err.Error())
		// fmt.Println("Server Database Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return errs.ServerError("Server Internal Error")
	}

	_, err = result.LastInsertId()
	// rowCount, err := result.LastInsertId()
	// rowCount, err := result.RowsAffected()

	// error check for LastInsertId function
	if err != nil {
		a.logger.Error(err.Error())
		// fmt.Println("Server Get RowID Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return errs.ServerError("Server Internal Error")
	}

	// c.IndentedJSON(http.StatusCreated, gin.H{"message": rowCount})
	return nil
}

// Update locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a AlbumSQL) Update(album domain.Album) *errs.Error {
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

	sqlstmt := "UPDATE albums SET title=?, artist=?, price=? where id=?"

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt, album.Title, album.Artist, album.Price, album.ID)

	// check err from server DB and Query DB
	if err != nil {
		a.logger.Error(err.Error())
		// fmt.Println("Server Database Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return errs.ServerError("Server Internal Error")
	}

	rowCount, err := result.RowsAffected()

	// error check for RowsAffected function
	if err != nil {
		a.logger.Error(err.Error())
		// fmt.Println("Server Get RowsAffected Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return errs.ServerError("Server Internal Error")
	}

	// there is no row found
	if rowCount == 0 {
		a.logger.Info("Record Not Found")
		// fmt.Println("Record Not Found")
		// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return errs.NotFound("Record Not Found")
	}

	// c.IndentedJSON(http.StatusAccepted, gin.H{"message": newAlbum})
	return nil
}

// Delete locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a AlbumSQL) Delete(id string) *errs.Error {
	// DB, err1 := sql.Open("sqlite3", "alochym.db")

	// if err1 != nil {
	// 	panic(err1)
	// }

	// defer DB.Close()

	// id := c.Param("id")

	sqlstmt := "DELETE FROM albums where id=?"

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt, id)

	// check err from server DB and Query DB
	if err != nil {
		a.logger.Error(err.Error())
		// fmt.Println("Server Database Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return errs.ServerError("Server Internal Error")
	}

	rowCount, err := result.RowsAffected()

	// error check for RowsAffected function
	if err != nil {
		a.logger.Error(err.Error())
		// fmt.Println("Server Get RowsAffected Error", err.Error())
		// c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
		return errs.ServerError("Server Internal Error")
	}

	// there is no row found
	if rowCount == 0 {
		a.logger.Info("Record Not Found")
		// fmt.Println("Record Not Found")
		// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
		return errs.NotFound("Record Not Found")
	}

	// c.IndentedJSON(http.StatusAccepted, gin.H{"message": rowCount})
	return nil
}
