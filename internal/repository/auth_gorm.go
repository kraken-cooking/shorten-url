package repository

import (
	"shorten-url-be/internal/domain/models"

	"gorm.io/gorm"
)

type AuthRepositoryGorm struct {
	db *gorm.DB
}

func NewAuthRepositoryGorm(db *gorm.DB) *AuthRepositoryGorm {
	return &AuthRepositoryGorm{db: db}
}

func (r *AuthRepositoryGorm) Login(user *models.User) (*models.User, error) {
	if err := r.db.Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AuthRepositoryGorm) SignUp(user *models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
