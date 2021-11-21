package memory

import (
	"errors"

	"github.com/alochym01/web-w-gin/domain"
)

// Album ...
type Album struct {
	al []domain.Album
}

// NewAlbum ...
func NewAlbum() Album {
	albums := []domain.Album{
		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return Album{
		al: albums,
	}
}

// FindAll responds with the list of all albums as JSON.
func (a *Album) FindAll() ([]domain.Album, error) {
	return a.al, nil
}

// FindByID locates the album whose ID value matches the id
func (a *Album) FindByID(id int) (*domain.Album, error) {
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range a.al {
		if album.ID == id {
			return &album, nil
		}
	}
	return nil, errors.New("Record Not Found")
}

// Create adds an album from JSON received in the request body.
func (a *Album) Create(album domain.Album) (*int64, error) {
	// Add the new album to the slice.
	a.al = append(a.al, album)
	return &album.ID, nil
}

// Update locates the album whose ID value matches the id
func (a *Album) Update(temp domain.Album) error {
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, album := range a.al {
		if album.ID == temp.ID {
			a.al[i] = temp
			return nil
		}
	}
	return errors.New("Record Not Found")
}

// Delete locates the album whose ID value matches the id
func (a *Album) Delete(id int) error {
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, album := range a.al {
		if album.ID == id {
			tempAlbum := make([]domain.Album, len(a.al)-1)
			copy(tempAlbum[:i], a.al[:i])
			copy(tempAlbum[i:], a.al[i+1:])
			a.al = tempAlbum
			return nil
		}
	}
	return errors.New("Record Not Found")
}
