package entities

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Sesseon_id   uuid.UUID `json:"session_id"`
	UserID       uuid.UUID `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}
