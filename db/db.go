package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBanco() *sql.DB {
	connStr := "user=app-web-go password=app-web-go dbname=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}

	return db
}
