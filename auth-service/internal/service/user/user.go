package user

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"auth-service/internal/storage"
	"auth-service/pkg/validator"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
		return fmt.Errorf("failed to check email: %w", err)
	}
	if exists {
		return errors.New("email already exists")
	}
	exists1, err := us.repo.NicknameExists(req.Nickname)
	if err != nil {
		return fmt.Errorf("failed to check nickname: %w", err)
	}
	if exists1 {
		return errors.New("nickname already exists")
	}

	if err := validator.ValidatePassword(req.Password); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := models.User{
		Email:    req.Email,
		Nickname: req.Nickname,
		Password: string(hashedPassword),
	}

	return us.repo.Save(user)
}

func (us *Service) Update(email string, newNickname string, newPassword string) error {
	if newPassword != "" {
		if err := validator.ValidatePassword(newPassword); err != nil {
			return err
		}
	}

	user, err := us.repo.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	user.Nickname = newNickname

	if newPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		user.Password = string(hashedPassword)
	}

	return us.repo.Update(user)
}
