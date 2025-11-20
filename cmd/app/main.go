package main

import (
	"backend-auth-service-app/pkg/databases"

	_ "github.com/lib/pq"
)

func main() {
	db := databases.DB()
	defer db.Close()

}
