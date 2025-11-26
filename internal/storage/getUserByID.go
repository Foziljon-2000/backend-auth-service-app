package storage

import (
	"backend-auth-service-app/internal/entities"
	responses "backend-auth-service-app/pkg/errors"
	"database/sql"

	"github.com/google/uuid"
)

func GetUserByID(user_id uuid.UUID) (user entities.User, err error) {
	row := database.QueryRow("select * from users where user_id = $1", user_id)
	err = row.Scan(&user.User_ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.ExpiresAT)
	if err == sql.ErrNoRows {
		err = responses.ErrAccountDoesNotExist
		return
	} else if err != nil {
		err = responses.ErrInternalServer
		return
	}
	err = nil
	return
}
