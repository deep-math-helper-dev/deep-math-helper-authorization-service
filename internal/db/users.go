package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func AddMeta(login string, password string) {
	db, err := sql.Open("sqlite", "usersmeta.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("insert into Meta (login, password) values ($1, $2)", login, password)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Succesful register")
}
