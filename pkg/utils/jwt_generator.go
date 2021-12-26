package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateNewAccessToken func to generate a new Access token.
func GenerateNewAccessToken() (string, error) {
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")

	// Set expiration minutes for secret key from .env file.
	expirationMinutes, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRATION_MINUTES"))

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expirationMinutes)).Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error with empty JWT token if generation failed.
		return "", err
	}

	return t, nil
}
