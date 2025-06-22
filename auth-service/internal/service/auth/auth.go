package auth

import (
	"auth-service/internal/service"
	"auth-service/internal/storage"
	"auth-service/pkg/tokenManager"
	"auth-service/pkg/validator"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
		return "", err
	}

	user, err := as.repo.GetByEmail(req.Email)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := as.tokenManager.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func (as *Service) ValidateToken(token string) (bool, error) {

	_, err := as.tokenManager.ParseToken(token)
	if err != nil {
		return false, fmt.Errorf("invalid token: %w", err)
	}

	return true, nil
}
