package auth

import (
	"fmt"
	"os"
	"time"
	"user-service/protos"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

func GenerateToken(id , role string) (*protos.UserToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"role": role,
			"id":   id,
			"exp":  time.Now().Add(time.Minute * 30).Unix(),
		})

	accessToken, err := token.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return nil, fmt.Errorf("failed to generate jwt token: %v", err)
	}
	return &protos.UserToken{Token: accessToken, ExpiryInMin: 30}, nil
}

