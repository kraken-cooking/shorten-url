package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key used to sign JWT
var jwtSecret = []byte("supper_secret")

// Claims is a struct that represents the JWT claims
type AuthClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token with a custom claim
func GenerateJWT(userID uint) (string, error) {
	// Set custom claims
	claims := &AuthClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign the token: %w", err)
	}

	return tokenString, nil
}

// ParseJWT parses and validates the JWT token and returns the claims
func ParseJWT(tokenString string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the key used for signing the token
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Check if the token is valid and the claims are what we expect
	claims, ok := token.Claims.(*AuthClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("Token is not valid: %w", err)
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("Token has expired")
	}

	return claims, nil
}
