package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Add(login string, password string) { // func to add user into db
	db, _ := sql.Open("sqlite", "usersmeta.db")                                    //open db
	db.Exec("insert into Meta (login, password) values ($1, $2)", login, password) // insert
	defer db.Close()
}
