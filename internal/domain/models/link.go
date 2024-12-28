package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	OriginalURL string `gorm:"not null" json:"originalUrl"`
	ShortURL    string `gorm:"unique;not null" json:"shortUrl"`

	UserID uint
}
