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

func VerifyToken(tokenstring string) (string, error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret_key")), nil
	})
	if err != nil && !token.Valid {
		return "", fmt.Errorf("invalid Token")

	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		role, ok := claims["role"].(string)
		if !ok {
			return "", fmt.Errorf("no role found in the JWT")
		}
		return role, nil
	}
	return "", fmt.Errorf("invalid token")
}
