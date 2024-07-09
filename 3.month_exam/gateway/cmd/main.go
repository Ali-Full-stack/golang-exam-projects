package main

import (
	"gateway/api/routes"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mux := routes.Routes()

	log.Println("Gateway: Started listening on port", os.Getenv("gateway_url"))
	log.Fatal(http.ListenAndServe(os.Getenv("gateway_url"), mux))

}
