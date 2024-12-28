package repository

import (
	"shorten-url-be/internal/domain/models"

	"gorm.io/gorm"
)

// LinkRepositoryGorm is a GORM implementation of the LinkRepository
type LinkRepositoryGorm struct {
	db *gorm.DB
}

// NewLinkRepositoryGorm creates a new instance of LinkRepositoryGorm
func NewLinkRepositoryGorm(db *gorm.DB) *LinkRepositoryGorm {
	return &LinkRepositoryGorm{db: db}
}

func (r *LinkRepositoryGorm) Create(link *models.Link) (*models.Link, error) {
	if err := r.db.Create(link).Error; err != nil {
		return nil, err
	}
	return link, nil
}

func (r *LinkRepositoryGorm) GetByID(id uint) (*models.Link, error) {
	var link models.Link
	if err := r.db.First(&link, id).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *LinkRepositoryGorm) GetByShortURL(shortURL string) (*models.Link, error) {
	var link models.Link
	if err := r.db.Where("short_url = ?", shortURL).First(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *LinkRepositoryGorm) GetAll() ([]models.Link, error) {
	var links []models.Link
	if err := r.db.Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

func (r *LinkRepositoryGorm) Update(link *models.Link) (*models.Link, error) {
	if err := r.db.Save(link).Error; err != nil {
		return nil, err
	}
	return link, nil
}

func (r *LinkRepositoryGorm) Delete(id uint) error {
	if err := r.db.Delete(&models.Link{}, id).Error; err != nil {
		return err
	}
	return nil
}
