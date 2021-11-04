package service

import (
	"github.com/alochym01/learn-web/basic-layout-handler-service/domain"
)

type AlbumServiceRepository interface {
	GetAlbums() ([]domain.Album, error)
	ByID(string) (*domain.Album, error)
	Create(domain.AlbumRequest) error
	Update(string, domain.AlbumRequest) error
	Delete(string) error
}

type AlbumService struct {
	storageRepo domain.AlbumRepository
}

func NewAlbumService(repo domain.AlbumRepository) AlbumService {
	return AlbumService{
		storageRepo: repo,
	}
}

// GetAlbums responds with the list of all albums as JSON.
func (svc AlbumService) GetAlbums() ([]domain.Album, error) {
	return svc.storageRepo.FindAll()
}

// ByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (svc AlbumService) ByID(id string) (*domain.Album, error) {
	return svc.storageRepo.FindByID(id)
}

// PostAlbums adds an album from JSON received in the request body.
func (svc AlbumService) Create(alRequest domain.AlbumRequest) error {
	var newAlbum domain.Album

	newAlbum.ID = "3" // can be assign auto increase using MySQL PRIMARY KEY
	newAlbum.Title = alRequest.Title
	newAlbum.Artist = alRequest.Artist
	newAlbum.Price = alRequest.Price

	return svc.storageRepo.Create(newAlbum)
}

// Update locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (svc AlbumService) Update(id string, alRequest domain.AlbumRequest) error {
	var updateAlbum domain.Album

	updateAlbum.ID = id
	updateAlbum.Title = alRequest.Title
	updateAlbum.Artist = alRequest.Artist
	updateAlbum.Price = alRequest.Price

	return svc.storageRepo.Update(updateAlbum)
}

// Delete locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (svc AlbumService) Delete(id string) error {
	return svc.storageRepo.Delete(id)
}
