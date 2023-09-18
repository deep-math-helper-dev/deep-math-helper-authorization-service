package entity

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(login, password string) string {
	l := md5.New()
	l.Write([]byte(login))

	p := md5.New()
	p.Write([]byte(password))

	userhash := hex.EncodeToString(l.Sum(nil)) + hex.EncodeToString(p.Sum(nil))

	return userhash
}
