package models

import (
	"log"
)

// User model
type User struct {
	name     string
	password string
}

var user User

// AuthenticateCredentials validates whether the user is valid.
func AuthenticateCredentials(username string, password string) bool {
	// byte, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	rows, err := db.Query(`SELECT (1) FROM users WHERE username = $1 AND password = $2 LIMIT 1`, username, password)
	if err != nil {
		log.Fatal(err)
	}
	return rows.Next()
}
