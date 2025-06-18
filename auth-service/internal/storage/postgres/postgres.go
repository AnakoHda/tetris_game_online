package postgres

import (
	"auth-service/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) ExistsByEmail(email string) (bool, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM users WHERE email = $1", email)
	return count > 0, err
}

func (r *PostgresRepository) Save(user models.User) error {
	_, err := r.db.Exec(`
		INSERT INTO users (email, nickname, password, created_time, updated_time)
		VALUES ($1, $2, $3, NOW(), NOW())
	`, user.Email, user.Nickname, user.Password)
	return err
}

func (r *PostgresRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	return user, err
}

func (r *PostgresRepository) Update(user models.User) error {
	_, err := r.db.Exec(`
		UPDATE users SET nickname = $1, password = $2, updated_time = NOW() WHERE email = $3
	`, user.Nickname, user.Password, user.Email)
	return err
}
