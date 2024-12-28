package repository

import "shorten-url-be/internal/domain/models"

type AuthRepository interface {
	SignUp(user *models.User) (*models.User, error)
	Login(user *models.User) (*models.User, error)
}
