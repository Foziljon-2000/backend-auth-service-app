package databases

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	st := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASS") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	db, err := sql.Open("postgres", st)
	if err != nil {
		log.Fatalf("Ошибка открытия соединения: %v", err)
	}

	return db
}
