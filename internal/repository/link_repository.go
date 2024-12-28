package repository

import "shorten-url-be/internal/domain/models"

// LinkRepository defines the CRUD operations for Link entity
type LinkRepository interface {
	Create(link *models.Link) (*models.Link, error)
	GetByID(id uint) (*models.Link, error)
	GetByShortURL(shortURL string) (*models.Link, error)
	GetAll() ([]models.Link, error)
	Update(link *models.Link) (*models.Link, error)
	Delete(id uint) error
}
