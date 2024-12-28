package usecase

import (
	"shorten-url-be/internal/domain/models"
	"shorten-url-be/internal/repository"
	"shorten-url-be/internal/utils"
)

// LinkUseCase defines the business logic for managing links
type LinkUseCase struct {
	repo repository.LinkRepository
}

// NewLinkUseCase creates a new instance of LinkUseCase
func NewLinkUseCase(repo repository.LinkRepository) *LinkUseCase {
	return &LinkUseCase{repo: repo}
}

func (uc *LinkUseCase) CreateLink(originalURL string, userID uint) (*models.Link, error) {
	shortURL := utils.GenerateShortURL()
	link := &models.Link{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
		UserID:      userID,
	}
	return uc.repo.Create(link)
}

func (uc *LinkUseCase) GetLinkByShortURL(shortURL string) (*models.Link, error) {
	return uc.repo.GetByShortURL(shortURL)
}

func (uc *LinkUseCase) GetAllLinks() ([]models.Link, error) {
	return uc.repo.GetAll()
}

func (uc *LinkUseCase) UpdateLink(id uint, originalURL string) (*models.Link, error) {
	link, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	link.OriginalURL = originalURL
	return uc.repo.Update(link)
}

func (uc *LinkUseCase) DeleteLink(id uint) error {
	return uc.repo.Delete(id)
}
