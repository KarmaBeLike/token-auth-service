package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type RefreshToken struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	TokenHash string    `json:"-"`
	IP        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
