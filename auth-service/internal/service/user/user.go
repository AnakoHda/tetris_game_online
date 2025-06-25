package user

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"auth-service/internal/storage"
	"auth-service/pkg/validator"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type Service struct {
	repo storage.Repository
}

func NewUserService(repo storage.Repository) *Service {
	return &Service{repo: repo}
}
func (us *Service) Register(req service.RegisterRequest) error {
	if err := validator.ValidateEmail(req.Email); err != nil {
		return err
	}

	exists, err := us.repo.EmailExists(req.Email)
	if err != nil {
		slog.Error("failed check email exists", "err", err)
		return fmt.Errorf("server error")
	}
	if exists {
		slog.Warn("nickname already exists")
		return errors.New("email already exists")
	}
	exists1, err := us.repo.NicknameExists(req.Nickname)
	if err != nil {
		slog.Error("failed check nickname exists", "err", err)
		return fmt.Errorf("server error")
	}
	if exists1 {
		slog.Warn("nickname already exists")
		return fmt.Errorf("nickname already exists")
	}

	if err := validator.ValidatePassword(req.Password); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed generate from pass", "err", err)
		return fmt.Errorf("server error")
	}

	user := models.User{
		Email:    req.Email,
		Nickname: req.Nickname,
		Password: string(hashedPassword),
	}
	err = us.repo.Save(user)
	if err != nil {
		slog.Error("failed save", "user", user.Email, "err", err)
		return fmt.Errorf("server error")
	}
	return nil
}

func (us *Service) Update(email string, newNickname string, newPassword string) error {
	exists, err := us.repo.NicknameExists(newNickname)
	if err != nil {
		slog.Error("failed check nickname exists", "err", err)
		return fmt.Errorf("server error")
	}
	if exists {
		slog.Warn("nickname already exists")
		return fmt.Errorf("new nickname already exists")
	}

	if err := validator.ValidatePassword(newPassword); err != nil {
		slog.Warn("failed validate password", "err", err)
		return fmt.Errorf("new %w", err)
	}

	user, err := us.repo.GetByEmail(email)
	if err != nil {
		slog.Error("failed to get user by ", "err", err)
		return fmt.Errorf("server error")
	}

	user.Nickname = newNickname

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed generate from pass", "err", err)
		return fmt.Errorf("server error")
	}
	user.Password = string(hashedPassword)

	err = us.repo.Update(user)
	if err != nil {
		slog.Error("failed update db", "user", user.Email, "err", err)
		return fmt.Errorf("server error")
	}
	return nil
}
