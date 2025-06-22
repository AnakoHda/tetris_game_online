package postgres

import (
	"auth-service/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) EmailExists(email string) (bool, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM users WHERE email = $1", email)
	if err != nil {
		slog.Error("EmailExists: failed to execute query", "error", err, "email", email)
		return count > 0, err
	}

	slog.Info("ExistsByEmail: result", "email", email, "exists", count > 0)
	return count > 0, err
}
func (r *PostgresRepository) NicknameExists(nickname string) (bool, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM users WHERE nickname = $1", nickname)
	if err != nil {
		slog.Error("NicknameExists: failed to execute query", "nickname", nickname, "error", err)
		return count > 0, err
	}
	slog.Info("NicknameExists: result", "nickname", nickname, "exists", count > 0)
	return count > 0, err
}
func (r *PostgresRepository) Save(user models.User) error {
	_, err := r.db.Exec(`
		INSERT INTO users (email, nickname, password, created_time, updated_time)
		VALUES ($1, $2, $3, NOW(), NOW())
	`, user.Email, user.Nickname, user.Password)
	if err != nil {
		slog.Error("Save: failed to save user with", "email", user.Email, "error", err)
		return err
	}
	slog.Info("Save: user saved successfully", "email", user.Email)
	return err
}

func (r *PostgresRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		slog.Error("GetByEmail: failed to get user by", "email", email, "err", err)
		return user, err
	}
	slog.Info("GetByEmail: user found", "email", user.Email)
	return user, err
}

func (r *PostgresRepository) Update(user models.User) error {
	_, err := r.db.Exec(`
		UPDATE users SET nickname = $1, password = $2, updated_time = NOW() WHERE email = $3
	`, user.Nickname, user.Password, user.Email)
	if err != nil {
		slog.Error("Update: failed to update user with", "email", user.Email, "err", err)
		return err
	}
	slog.Info("Update: user updated successfully", "email", user.Email)
	return err
}
