package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/auth/hash"
	"gateway/auth/jwt"
	"gateway/auth/validate"
	"gateway/internal/file"
	"gateway/internal/models"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func CheckAdminPassword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "authorization info required !!", http.StatusNonAuthoritativeInfo)
			return
		}

		if err := file.CheckAdminfromFile("internal/data/admins.json", id, password); err != nil {
			http.Error(w, "Permission denied !!", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ValidateRequestBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var cl models.ClientInfo
		if err := json.NewDecoder(r.Body).Decode(&cl); err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}
		errors := validate.ValidateInfo(cl.Name, cl.Email, cl.Phone, cl.Card_number, cl.Home_address, float32(cl.Balance))
		if errors != nil {
			log.Println(errors)
			http.Error(w, "Invalid Body Request", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "client", cl)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
func ValidateRequestDriver(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dr models.DriverInfo
		if err := json.NewDecoder(r.Body).Decode(&dr); err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}
		errors :=validate.ValidateInfo(dr.Name, dr.Email, dr.Phone, dr.Card_number, dr.Home_Address, float32(dr.Balance))
		if errors != nil {
			log.Println(errors)
			http.Error(w, "Invalid Body Request", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "driver", dr)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func LoginAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, email, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Invalid  Authorization info!", http.StatusForbidden)
			return
		}
		if err := file.CheckClientFromFile("internal/data/clients.json", models.ClientLogin{Id: id, Email: email}); err != nil {
			fmt.Println(err)
			http.Error(w, "failed login : Invalid client Id", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "email", email)
		ctx2 := context.WithValue(ctx, "id", id)

		next.ServeHTTP(w, r.WithContext(ctx2))
	})
}

func CheckOrderToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("token")
		if err := jwt.VerifyToken(clientToken, os.Getenv("secret_key")); err != nil {
			http.Error(w, "Invalid Token !!", http.StatusBadRequest)
			return
		}
		clientId := jwt.GetIdFromToken(clientToken)
		ctx := context.WithValue(r.Context(), "id", clientId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func IsSuperAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Invalid  Authorization info!", http.StatusForbidden)
			return
		}
		if !hash.ValidateHashPassword(os.Getenv("super_admin_password"), password) && (os.Getenv("super_admin_id")) != id {
			http.Error(w, "Unathorized !", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
