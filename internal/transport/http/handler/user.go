package handler

import (
	"backend-auth-service-app/internal/service"
	responses "backend-auth-service-app/pkg/errors"
	httpResponser "backend-auth-service-app/pkg/http"
	"encoding/json"
	"net/http"
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
	case "PUT":
	case "DELETE":
	default:
		resp.Message = responses.ErrMethodNotAllowed.Error()
		return
	}

	resp.Message = responses.ErrSuccess.Error()
}
