package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	OriginalURL string `json:"originalUrl"`
	ShortURL    string `json:"shortUrl"`

	UserID uint
}
