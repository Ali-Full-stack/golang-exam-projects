package token

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

func GenerateToken(id, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name": name,
			"id":   id,
			"exp":  time.Now().Add(time.Minute * 30).Unix(),
		})

	accessToken, err := token.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		log.Println("Failed generating access token:", err)
		return "", err
	}
	return accessToken, nil
}
