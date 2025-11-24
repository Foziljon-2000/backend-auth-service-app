package handler

import (
	"backend-auth-service-app/internal/service"
	responses "backend-auth-service-app/pkg/errors"
	httpResponser "backend-auth-service-app/pkg/http"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func User(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Converter(w)

	switch r.Method {
	case "POST":
		var newUser CreateUserDTO

		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			resp.Message = responses.ErrBadRequest.Error()
			return
		}

		defer r.Body.Close()

		if err := service.CreateUser(newUser.Email, newUser.Password); err != nil {
			resp.Message = err.Error()
			return
		}

	case "GET":
		userIdStr := r.URL.Query().Get("user_id")
		userID, err := uuid.Parse(userIdStr)
		if err != nil {
			resp.Message = responses.ErrBadRequest.Error()
			return
		}

		user, err := service.GetUser(userID)
		if err != nil {
			return
		}

		resp.Payload = user

	case "PUT":
	case "DELETE":
	default:
		resp.Message = responses.ErrMethodNotAllowed.Error()
		return
	}

	resp.Message = responses.ErrSuccess.Error()
}
