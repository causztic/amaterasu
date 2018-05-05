package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Name     string
	email    string
	password string
}

var user User

// AuthenticateCredentials validates whether the user is valid.
func AuthenticateCredentials(email string, password string) bool {
	byte, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	rows, err := db.Query(`SELECT (1) FROM users WHERE email = $1 AND password = $2 LIMIT 1`, email, string(byte))
	if err != nil {
		log.Fatal(err)
	}
	return rows.Next()
}
