package storage

import (
	"github.com/alochym01/learn-web/basic-layout-handler-service-error-handler-refactor/domain"
	"github.com/alochym01/learn-web/basic-layout-handler-service-error-handler-refactor/errs"
)

type AlbumMemory struct {
	al []domain.Album
}

func NewAlbumMemory() AlbumMemory {
	albums := []domain.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return AlbumMemory{
		al: albums,
	}
}

// FindAll responds with the list of all albums as JSON.
func (a *AlbumMemory) FindAll() ([]domain.Album, *errs.Error) {
	return a.al, nil
}

// FindByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumMemory) FindByID(id string) (*domain.Album, *errs.Error) {
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range a.al {
		if album.ID == id {
			return &album, nil
		}
	}
	return nil, errs.NotFound("Record Not Found")
}

// Create(Album) error
// Create adds an album from JSON received in the request body.
func (a *AlbumMemory) Create(album domain.Album) *errs.Error {

	// Add the new album to the slice.
	a.al = append(a.al, album)
	return nil
}

// Update locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumMemory) Update(temp domain.Album) *errs.Error {
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, album := range a.al {
		if album.ID == temp.ID {
			a.al[i] = temp
			return nil
		}
	}
	return errs.NotFound("Record Not Found")
}

// Delete locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (a *AlbumMemory) Delete(id string) *errs.Error {
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
	return errs.NotFound("Record Not Found")
}
