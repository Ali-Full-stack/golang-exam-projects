package middleware

import (
	"gateway/auth"
	"log"
	"net/http"
)

func ConfirmToken(next http.Handler)http.Handler{
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){
		token :=r.Header.Get("token")

		if err :=auth.VerifyToken(token); err !=nil {
			log.Println("Invalid Token:",err)
			http.Error(w, "Invalid Token :", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
