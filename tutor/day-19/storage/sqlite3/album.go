package sqlite3

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/alochym01/web-w-gin/domain"
	"go.uber.org/zap"
)

// Album ...
type Album struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewAlbum ...
func NewAlbum(DB *sql.DB, l *zap.Logger) Album {
	return Album{
		db:     DB,
		logger: l,
	}
}

// FindAll ...
func (a Album) FindAll() ([]domain.Album, error) {
	sqlstmt := "select * from albums"

	rows, err := a.db.Query(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		return nil, err
	}

	albums := []domain.Album{}

	for rows.Next() {
		a := domain.Album{}
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		// check err from server DB and Scan function
		if err != nil {
			fmt.Println("Server Scanning Rows Error", err.Error())
			return nil, err
		}
		albums = append(albums, a)
	}
	return albums, nil
}

// FindByID locates the album whose ID value matches the id
func (a Album) FindByID(id int) (*domain.Album, error) {
	sqlstmt := "select * from albums where id=?"

	row := a.db.QueryRow(sqlstmt, id)
	album := domain.Album{}
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			a.logger.Warn(err.Error())
			return nil, err
		}

		a.logger.Error(err.Error())
		return nil, err
	}

	return &album, nil
}

// Create adds an album from JSON received in the request body.
func (a Album) Create(album domain.Album) (*int64, error) { // sqlstmt - Avoid SQL Injection Attack
	sqlstmt := fmt.Sprintf("INSERT INTO albums(id, title, artist, price) VALUES(%d,\"%s\", \"%s\", %f)",
		album.ID,
		album.Title,
		album.Artist,
		album.Price,
	)

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		return nil, err
	}

	rowID, err := result.LastInsertId()
	// rowCount, err := result.LastInsertId()
	// rowCount, err := result.RowsAffected()

	// error check for LastInsertId function
	if err != nil {
		fmt.Println("Server Get Row Error", err.Error())
		return nil, err
	}

	return &rowID, nil
}

// Update locates the album whose ID value matches the id
func (a Album) Update(album domain.Album) error {
	sqlstmt := fmt.Sprintf("UPDATE albums SET title=\"%s\", artist=\"%s\", price=%f where id=%d",
		album.Title,
		album.Artist,
		album.Price,
		album.ID,
	)

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		return err
	}

	rowCount, err := result.RowsAffected()

	// error check for RowsAffected function
	if err != nil {
		fmt.Println("Server Get RowsAffected Error", err.Error())
		return err
	}

	// there is no row found
	if rowCount == 0 {
		fmt.Println("Record Not Found")
		return errors.New("Record Not Found")
	}

	return nil
}

// Delete locates the album whose ID value matches the id
func (a Album) Delete(id int) error {
	sqlstmt := fmt.Sprintf("DELETE FROM albums where id=%d", id)

	// Execute SQL Statements
	result, err := a.db.Exec(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		fmt.Println("Server Database Error", err.Error())
		return err
	}

	rowCount, err := result.RowsAffected()

	// error check for RowsAffected function
	if err != nil {
		fmt.Println("Server Get RowsAffected Error", err.Error())
		return err
	}

	// there is no row found
	if rowCount == 0 {
		fmt.Println("Record Not Found")
		return errors.New("Record Not Found")
	}

	return nil
}
