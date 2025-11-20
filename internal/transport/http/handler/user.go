package handler

import (
	responses "backend-auth-service-app/pkg/errors"
	httpResponser "backend-auth-service-app/pkg/http"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	var resp httpResponser.Response
	defer resp.Converter(w)

	switch r.Method {
	case "POST":
	case "GET":
	case "PUT":
	case "DELETE":
	default:
		resp.Message = responses.ErrMethodNotAllowed.Error()
		return
	}

	resp.Message = responses.ErrSuccess.Error()
}
