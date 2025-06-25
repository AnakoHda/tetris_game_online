package auth

import (
	"auth-service/internal/service"
	"auth-service/internal/storage"
	"auth-service/pkg/tokenManager"
	"auth-service/pkg/validator"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type Service struct {
	repo         storage.Repository
	tokenManager tokenManager.TokenManager
}

func NewAuthService(repo storage.Repository, tokenManager tokenManager.TokenManager) *Service {
	return &Service{repo: repo, tokenManager: tokenManager}
}

func (as *Service) Login(req service.LoginRequest) (string, error) {
	if err := validator.ValidateEmail(req.Email); err != nil {
		slog.Warn("validate email failed", "err", err)
		return "", fmt.Errorf("uncorrect email format")
	}
	user, err := as.repo.GetByEmail(req.Email)
	if err != nil {
		slog.Error("failed to get user by", "email", req.Email, "err", err)
		return "", fmt.Errorf("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		slog.Error("failed compare hash and pass", "err", err)
		return "", fmt.Errorf("invalid email or password")
	}

	token, err := as.tokenManager.GenerateToken(user.ID, user.Email)
	if err != nil {
		slog.Error("failed generate token", "err", err)
		return "", fmt.Errorf("server error")
	}
	return token, nil
}

func (as *Service) ValidateToken(token string) (bool, error) {
	_, err := as.tokenManager.ParseToken(token)
	if err != nil {

		return false, fmt.Errorf("fail validate token: %w", err)
	}
	return true, nil
}
