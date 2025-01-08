package repository

import (
	"fmt"
	"shorten-url-be/internal/domain/models"
	"strings"

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

		if strings.Contains(err.Error(), "ERROR: duplicate key value") {
			return nil, fmt.Errorf("username already taken")
		} else {
			return nil, fmt.Errorf("Internal error")
		}

	}

	return user, nil
}
