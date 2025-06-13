package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(username string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(jwtKey)
}
func ValidateJWT(tokenStr string) (string, error) {
	tokenParsed, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := tokenParsed.Claims.(*jwt.RegisteredClaims); ok && tokenParsed.Valid {
		return claims.Subject, nil
	}

	return "", jwt.ErrTokenInvalidClaims
}
