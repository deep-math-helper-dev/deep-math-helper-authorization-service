package entity

import (
	"strings"

	HashPass "github.com/vzglad-smerti/password_hash"
)

// func for format users request data
func TakeData(login, password string) (formattedLogin string, hash string) {
	hash, _ = HashPass.Hash(password)       // hashing password
	formattedLogin = strings.ToLower(login) // format login
	return formattedLogin, hash
}
