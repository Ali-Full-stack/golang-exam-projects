package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println("failed to generate hash password:",err)
	}
	return string(bytes)
}
