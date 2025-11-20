package entities

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	UserID       uuid.UUID
	RefreshToken string
	ExpiresAt    time.Time
}
