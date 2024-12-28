package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateShortURL generates a random 8-character short URL using base62 encoding
func GenerateShortURL() string {
	const shortURLLength = 8

	var shortURL strings.Builder
	for i := 0; i < shortURLLength; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(base62Chars))))
		shortURL.WriteByte(base62Chars[index.Int64()])
	}

	return shortURL.String()
}
