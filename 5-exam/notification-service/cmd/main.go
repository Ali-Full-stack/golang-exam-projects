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

	userClient, err :=kafka.ConnectKafka(os.Getenv("kafka_url"), "user-registration")
	if err != nil {
		log.Fatal(err)
	}
	defer userClient.Close()
	bookingClient, err :=kafka.ConnectKafka(os.Getenv("kafka_url"), "booking-confirmation")
	if err != nil {
		log.Fatal(err)
	}
	defer bookingClient.Close()

	kafka.NewNotify()

	mux.HandleFunc("GET /ws", handler.HandleWebSocket)

	log.Println("Notification-service: Server is listening on port:", os.Getenv("notify_url"))
	log.Fatal(http.ListenAndServe(os.Getenv("notify_url"), mux))

}
