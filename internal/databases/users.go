package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// Also FIXME!!
type UserMeta struct {
	DB *sql.DB
}

func (u *UserMeta) AddMeta(login string, password string) *sql.Result {
	db, err := sql.Open("sqlite", "usersmeta.db")
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
