package storage

import "auth-service/internal/models"

type Repository interface {
	EmailExists(email string) (bool, error)
	NicknameExists(nickname string) (bool, error)
	Save(user models.User) error
	GetByEmail(email string) (models.User, error)
	Update(user models.User) error
}
