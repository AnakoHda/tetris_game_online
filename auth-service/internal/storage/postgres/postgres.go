package postgres

import (
	"auth-service/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	queryEmailExists    = "SELECT COUNT(*) FROM users WHERE email = $1"
	queryNicknameExists = "SELECT COUNT(*) FROM users WHERE nickname = $1"
	queryInsertUser     = `
	INSERT INTO users (email, nickname, password, created_time, updated_time)
	VALUES ($1, $2, $3, NOW(), NOW())`
	queryGetUserByEmail = "SELECT * FROM users WHERE email = $1"
	queryUpdateUser     = `
	UPDATE users 
	SET nickname = $1, password = $2, updated_time = NOW() 
	WHERE email = $3`
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) EmailExists(email string) (bool, error) {
	var count int
	err := r.db.Get(&count, queryEmailExists, email)
	if err != nil {
		return count > 0, err
	}
	return count > 0, err
}
func (r *PostgresRepository) NicknameExists(nickname string) (bool, error) {
	var count int
	err := r.db.Get(&count, queryNicknameExists, nickname)
	if err != nil {
		return count > 0, err
	}
	return count > 0, err
}
func (r *PostgresRepository) Save(user models.User) error {
	_, err := r.db.Exec(queryInsertUser, user.Email, user.Nickname, user.Password)
	if err != nil {
		return err
	}
	return err
}

func (r *PostgresRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, queryGetUserByEmail, email)
	if err != nil {
		return user, err
	}
	return user, err
}

func (r *PostgresRepository) Update(user models.User) error {
	_, err := r.db.Exec(queryUpdateUser, user.Nickname, user.Password, user.Email)
	if err != nil {
		return err
	}
	return err
}
