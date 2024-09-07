package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/KarmaBeLike/token-auth-service/internal/models"
	"github.com/KarmaBeLike/token-auth-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// Service для работы с пользователями
type UserService struct {
	Repository repository.UserRepo // Интерфейс для работы с репозиторием
}

// Интерфейс репозитория (для мокирования или использования разных реализаций)
type UserRepository interface {
	CreateUser(email string, password string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}

// Создание нового пользователя
func (s *UserService) RegisterUser(email string, password string) (*models.User, error) {
	// Проверка, существует ли пользователь с таким email
	existingUser, err := s.Repository.FindUserByEmail(email)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error checking if user exists: %v", err)
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Хеширование пароля (в сервисе можно продублировать)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Создание пользователя через репозиторий
	user, err := s.Repository.CreateUser(email, string(hashedPassword))
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Здесь можно добавить отправку приветственного письма или другие действия

	return user, nil
}

// Логин пользователя (пример другой бизнес-логики)
func (s *UserService) LoginUser(email string, password string) (*models.User, error) {
	// Получение пользователя по email через репозиторий
	user, err := s.Repository.FindUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		log.Printf("Failed to find user by email: %v", err)
		return nil, err
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Можно также сгенерировать JWT токены и вернуть их
	return user, nil
}
