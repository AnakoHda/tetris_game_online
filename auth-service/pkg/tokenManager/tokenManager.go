package tokenManager

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	GenerateToken(userID int, email string) (string, error)
	ParseToken(tokenStr string) (jwt.MapClaims, error)
}
