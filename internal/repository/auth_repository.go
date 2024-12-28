package repository

import "shorten-url-be/internal/domain/models"

// LinkRepository defines the CRUD operations for Link entity
type AuthRepository interface {
	SignUp(user *models.User) (*models.User, error)
	Login(user *models.User) (*models.User, error)
}
