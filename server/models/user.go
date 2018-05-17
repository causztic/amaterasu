package models

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	name     string
	password string
}

var user User

func GenerateHashedPassword(password string) string {
	byte, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(byte)
}

// AuthenticateCredentials validates whether the user is valid.
func AuthenticateCredentials(username string, password []byte) bool {
	var hashedPW []byte
	row := db.QueryRow(`SELECT password FROM users WHERE username = $1`, username).Scan(&hashedPW)
	if row == sql.ErrNoRows || row != nil {
		return false
	}
	err := bcrypt.CompareHashAndPassword(hashedPW, password)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
