package service

import (
	"backend-auth-service-app/internal/storage"
	responses "backend-auth-service-app/pkg/errors"

	"github.com/google/uuid"
)

func GetUser(user_id uuid.UUID) (result map[string]interface{}, err error) {
	// valid
	if user_id == uuid.Nil {
		err = responses.ErrBadRequest
		return
	}

	// logic
	user, err := storage.GetUserByID(user_id)
	if err != nil {
		return
	}

	result = make(map[string]interface{})

	result["user_id"] = user.User_ID
	result["email"] = user.Email

	return
}
