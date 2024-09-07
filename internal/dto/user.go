package dto

// Структура для запроса на регистрацию
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Структура для ответа на регистрацию
type RegisterResponse struct {
	Message string   `json:"message"`
	User    *UserDTO `json:"user"`
}

// Структура для запроса на логин
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Структура для ответа на логин
type LoginResponse struct {
	Message string   `json:"message"`
	User    *UserDTO `json:"user"`
	// Можно добавить JWT токен, если он требуется
	Token string `json:"token,omitempty"`
}

// Общая структура для информации о пользователе
type UserDTO struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
