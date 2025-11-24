package storage

import (
	"backend-auth-service-app/internal/entities"
	"time"
)

func CreateUser(user entities.User) (err error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.ExpiresAT = time.Date(2999, time.December, 31, 0, 0, 0, 0, time.Local)

	_, err = database.Exec("insert into users(user_id, email, password_hash, created_at, updated_at, expires_at) values($1, $2, $3, $4, $5, $6)", user.User_ID, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt, user.ExpiresAT)
	if err != nil {
		return
	}

	return
}
