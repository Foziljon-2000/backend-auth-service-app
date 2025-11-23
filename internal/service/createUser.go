package service

import (
	"backend-auth-service-app/internal/entities"
	"backend-auth-service-app/internal/storage"
	responses "backend-auth-service-app/pkg/errors"
)

func CreateUser(user entities.User) (err error) {
	// valid

	user_v2, err := storage.GetUserByEmail(user.Email)
	if err != responses.ErrUserLoginAlreadyExists || err != nil {
		return
	}

	if user.Email == user_v2.Email {
		err = responses.ErrUserLoginAlreadyExists
		return
	}

	// logic

	user.PasswordHash, err = CreateHashPassword(user.PasswordHash)
	if err != nil {
		err = responses.ErrInternalServer
		return
	}

	err = storage.CreateUser(user)
	if err != nil {
		return
	}

	return nil
}
