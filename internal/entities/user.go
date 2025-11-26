package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	User_ID      uuid.UUID `json:"user_id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"update_at"`
	ExpiresAT    time.Time `json:"expires_at"`
}
