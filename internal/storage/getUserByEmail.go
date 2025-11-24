package storage

import (
	"backend-auth-service-app/internal/entities"
	responses "backend-auth-service-app/pkg/errors"
	"database/sql"
)

func GetUserByEmail(email string) (user entities.User, err error) {
	row := database.QueryRow("select * from users where email = $1", email)

	err = row.Scan(&user.User_ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.ExpiresAT)
	if err == sql.ErrNoRows {
		err = responses.ErrUserDoesNotExist
		return
	} else if err != nil {
		err = responses.ErrInternalServer
		return
	}

	err = nil
	return
}
