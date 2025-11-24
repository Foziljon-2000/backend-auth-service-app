package storage

import (
	"backend-auth-service-app/internal/entities"
)

func CreateUser(user entities.User) (err error) {
	_, err = database.Exec("insert into users(user_id, email, password_hash, created_at, updated_at, expires_at) values($1, $2, $3, $4, $5, $6)", user.User_ID, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt, user.ExpiresAT)
	if err != nil {
		return
	}

	return
}
