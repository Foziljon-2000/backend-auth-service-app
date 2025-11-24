package handler

import (
	"backend-auth-service-app/internal/entities"
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
		var user entities.User
		jsonBody := json.NewDecoder(r.Body)

		err := jsonBody.Decode(&user)
		if err != nil {
			resp.Message = responses.ErrBadRequest.Error()
			return
		}

		err = service.CreateUser(user)
		if err != nil {
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
