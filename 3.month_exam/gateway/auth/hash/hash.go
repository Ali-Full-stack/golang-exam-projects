package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ValidateHashPassword(hash, password string) bool {
	if err :=bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		log.Println("failed to validate hash password:",err)
		return false
	}
	return true
}