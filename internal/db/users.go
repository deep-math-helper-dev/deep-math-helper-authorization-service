package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var db, _ = sql.Open("sqlite", "storage.db")

type User struct {
	UserHash string
}

func (u User) Add() { // func to add user into db
	db.Exec("insert into Meta (userHash) values ($1)", u.UserHash)
	defer db.Close()
}

func (u User) Auth() {
	db.Query("select * from Meta (userHash) where values ($1)", u.UserHash)
	defer db.Close()
}
