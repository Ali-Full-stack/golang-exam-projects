package api

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func ContentTypeMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// checking the request body  json or not
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Expecting  json format", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsValidId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		_, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid id format", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CheckPermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAdmin := r.Header.Get("admin")
		userPassword := r.Header.Get("password")

		admin := os.Getenv("ADMIN")
		password := os.Getenv("PASSWORD")

		if admin != userAdmin || password != userPassword {
			http.Error(w, "Permission denied ! Invalid Name or Password", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
