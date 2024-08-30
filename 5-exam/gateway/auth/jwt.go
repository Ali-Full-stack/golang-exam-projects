package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func VerifyToken(tokenstring string) (error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret_key")), nil
	})
	if err != nil && !token.Valid {
		return  fmt.Errorf("invalid Token")

	}
	return nil
}