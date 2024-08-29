package main

import (
	"log"
	"net/http"
	"notify-service/api/handler"
	"notify-service/kafka"
	"os"

	_ "github.com/joho/godotenv/autoload"
)


func main() {
	mux := http.NewServeMux()

	handler :=handler.NewKafka(&kafka.Kafka{})
	
	mux.HandleFunc("GET /ws", handler.HandleWebSocket)

	log.Println("Notification-service: Server is listening on port:", os.Getenv("notify_url"))
	log.Fatal(http.ListenAndServe(os.Getenv("notify_url"), mux))

}
