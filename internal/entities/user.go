package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	User_ID      uuid.UUID `json:"user_id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	UpdateAT     time.Time `json:"update_at"`
	ExpiresAT    time.Time `json:"expires_at"`
}
