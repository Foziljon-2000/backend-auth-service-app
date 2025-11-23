package main

import (
	"backend-auth-service-app/internal/storage"
	"backend-auth-service-app/pkg/databases"

	_ "github.com/lib/pq"
)

func main() {
	db := databases.DB()
	defer db.Close()

	storage.GetDB(db)

}
