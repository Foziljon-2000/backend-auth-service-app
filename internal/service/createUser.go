package service

import (
	"backend-auth-service-app/internal/entities"
	"backend-auth-service-app/internal/storage"
	responses "backend-auth-service-app/pkg/errors"
	"log"
	"time"

	"github.com/google/uuid"
)

func CreateUser(email string, password string) (err error) {
	// valid
	user_v2, err := storage.GetUserByEmail(email)
	if err != nil && err != responses.ErrUserDoesNotExist {
		log.Println(err)
		return
	}

	if email == "" || password == "" {
		err = responses.ErrBadRequest
		return
	}

	if email == user_v2.Email {
		err = responses.ErrUserLoginAlreadyExists
		return
	}

	// logic
	var newUser entities.User
	now := time.Now()
	expiresAT := time.Date(2999, time.December, 31, 0, 0, 0, 0, time.Local)

	hashPass, err := CreateHashPassword(password)
	if err != nil {
		err = responses.ErrInternalServer
		return
	}
	
	newUser = entities.User{
		User_ID:      uuid.New(),
		Email:        email,
		PasswordHash: hashPass,
		CreatedAt:    now,
		UpdatedAt:    now,
		ExpiresAT:    expiresAT,
	}

	err = storage.CreateUser(newUser)
	if err != nil {
		return
	}

	return nil
}
