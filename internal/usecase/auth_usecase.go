package usecase

import (
	"shorten-url-be/internal/domain/models"
	"shorten-url-be/internal/repository"
)

// AuthUseCase defines the business logic for managing links
type AuthUseCase struct {
	repo repository.AuthRepository
}

// NewAuthUseCase creates a new instance of AuthUseCase
func NewAuthUseCase(repo repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (uc *AuthUseCase) Login(username string, password string) (*models.User, error) {
	user := models.User{
		Username: username,
		Password: password,
	}

	return uc.repo.Login(&user)
}

func (uc *AuthUseCase) SignUp(username string, password string) (*models.User, error) {
	user := models.User{
		Username: username,
		Password: password,
	}
	return uc.repo.SignUp(&user)
}
