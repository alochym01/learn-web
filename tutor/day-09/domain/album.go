package domain

// Album represents data about a record Album.
type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float64
}

// // AlbumRepository ...
// type AlbumRepository interface {
// 	FindAll() ([]Album, error)
// 	FindByID(id string) (*Album, error)
// 	Create(Album) (*int, error)
// 	Update(Album) error
// 	Delete(int) error
// }

// AlbumRequest represents user request data.
type AlbumRequest struct {
	ID     int
	Title  string
	Artist string
	Price  float64
}
