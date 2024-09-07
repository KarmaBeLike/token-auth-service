package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KarmaBeLike/token-auth-service/internal/dto"
	"github.com/KarmaBeLike/token-auth-service/internal/service"
)

// Структура для хранения зависимостей
type UserHandler struct {
	UserService *service.UserService // Сервис для работы с пользователями
}

// Обработчик для регистрации пользователя
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest

	// Декодируем JSON из запроса
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Вызываем бизнес-логику для регистрации пользователя
	user, err := h.UserService.RegisterUser(req.Email, req.Password)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Подготавливаем DTO для ответа
	userDTO := &dto.UserDTO{
		ID:        user.ID.String(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"), // Форматирование времени
	}

	// Формируем ответ
	resp := dto.RegisterResponse{
		Message: "User registered successfully",
		User:    userDTO,
	}

	// Отправляем JSON ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Обработчик для логина пользователя
func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	// Декодируем JSON из запроса
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Вызываем бизнес-логику для логина пользователя
	user, err := h.UserService.LoginUser(req.Email, req.Password)
	if err != nil {
		log.Printf("Error logging in user: %v", err)
		http.Error(w, "Failed to login user", http.StatusUnauthorized)
		return
	}

	// Подготавливаем DTO для ответа
	userDTO := &dto.UserDTO{
		ID:        user.ID.String(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	// Формируем ответ
	resp := dto.LoginResponse{
		Message: "User logged in successfully",
		User:    userDTO,
		// JWT токен можно добавить сюда
	}

	// Отправляем JSON ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
