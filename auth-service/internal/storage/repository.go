package storage

import "auth-service/internal/models"

type Repository interface {
	ExistsByEmail(email string) (bool, error)
	Save(user models.User) error
	GetByEmail(email string) (models.User, error)
	Update(user models.User) error
}
