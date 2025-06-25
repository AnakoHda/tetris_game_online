package authHandler

import (
	"auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
)

type AuthHandler struct {
	authService service.AuthServiceInterface
}

func NewAuthHandler(authService service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn("invalid login request", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := ah.authService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	slog.Info("login success")
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ah *AuthHandler) ValidateToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		slog.Warn("empty token request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing Authorization header"})
		return
	}

	if !strings.HasPrefix(token, "Bearer ") {
		slog.Warn("incorrect request format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorization header must start with 'Bearer '"})
		return
	}
	tokenStr := strings.TrimPrefix(token, "Bearer ")
	valid, err := ah.authService.ValidateToken(tokenStr)
	if !valid || err != nil {
		slog.Warn("invalid token", "err", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}
	slog.Info("validate token success")
	c.JSON(http.StatusOK, gin.H{"message": "token is valid"})
}
