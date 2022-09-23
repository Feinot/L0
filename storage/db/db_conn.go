package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	err error
	db  *sql.DB
)

const (
	host     = "localhost"
	port     = 5432
	login    = "postgres"
	password = "123"
	dbname   = "stan"
)

func Close() {
	db.Close()
}
func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, login, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}

func GetDB() *sql.DB {
	return db
}
