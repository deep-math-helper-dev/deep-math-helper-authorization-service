package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// FIXME!
func AddMeta(login string, password string) *sql.Result {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	add, err := db.Exec("insert into Meta (login, password) values ($1, $2)", login, password)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Пользователь зарегестрирован")
	return &add
}
