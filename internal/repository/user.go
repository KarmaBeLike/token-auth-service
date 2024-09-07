package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/KarmaBeLike/token-auth-service/internal/models"
	"github.com/google/uuid"
)

type UserRepo struct {
	DB *sql.DB
}

// Создание нового пользователя
func (repo *UserRepo) CreateUser(email string, passwordHash string) (*models.User, error) {
	// Генерация UUID для пользователя
	userID := uuid.New()

	// Вставка пользователя в базу данных
	user := models.User{
		ID:           userID,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}

	err := repo.DB.QueryRow(
		`INSERT INTO users (id, email, password_hash, created_at)
         VALUES ($1, $2, $3, $4) RETURNING id`,
		user.ID, user.Email, user.PasswordHash, user.CreatedAt,
	).Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}

	return &user, nil
}

// Поиск пользователя по email
func (repo *UserRepo) FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := repo.DB.QueryRow(
		`SELECT id, email, password_hash, created_at
         FROM users WHERE email = $1`, email,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
