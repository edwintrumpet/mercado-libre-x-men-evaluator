package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbURL := os.Getenv("DB_URL")
	dbSSL := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf("postgres://%s:%s@%s/mutants?sslmode=%s", dbUser, dbPassword, dbURL, dbSSL)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
