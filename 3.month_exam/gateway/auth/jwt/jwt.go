package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(sub, email, id, key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"email": email,
			"role":  "client",
			"exp":   time.Now().Add(time.Minute * 30).Unix(),
		})

	accessToken, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println("Filed generating access token:", err)
	}
	return accessToken
}

func VerifyToken(token, key string) error {

	tokens, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return err
	}
	if !tokens.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func GetIdFromToken(token string) string {
	parts := strings.Split(token, ".")
	payload, _ := base64.RawURLEncoding.DecodeString(parts[1])
	var data map[string]interface{}
	json.Unmarshal(payload, &data)
	return data["id"].(string)
}
