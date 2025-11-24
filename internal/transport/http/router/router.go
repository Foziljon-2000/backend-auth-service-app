package router

import (
	"backend-auth-service-app/internal/transport/http/handler"

	"github.com/gorilla/mux"
)

func NewRouterCompl() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user", handler.User)

	return router
}
