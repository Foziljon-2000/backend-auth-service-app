package main

import (
	"backend-auth-service-app/internal/storage"
	"backend-auth-service-app/internal/transport/http/router"
	"backend-auth-service-app/pkg/databases"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db := databases.DB()
	defer db.Close()

	storage.GetDB(db)

	fmt.Println("listening localhost:2025")
	http.ListenAndServe("localhost:2025", router.NewRouterCompl())

}
