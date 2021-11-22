package domain

import "github.com/alochym01/web-w-gin/errs"

// Album represents data about a record Album.
type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float64
}

// AlbumRepository ...
type AlbumRepository interface {
	FindAll() ([]Album, error)
	FindByID(int) (*Album, *errs.AppErr)
	Create(Album) (*int64, error)
	Update(Album) error
	Delete(int) error
}

// AlbumRequest represents user request data.
type AlbumRequest struct {
	ID     int
	Title  string
	Artist string
	Price  float64
}
