package handler

import (
	"fmt"
	"log"
	"net/http"
	"notify-service/kafka"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Websocket struct {
	Kafka *kafka.Kafka
}

func NewKafka(k *kafka.Kafka)*Websocket{
	return &Websocket{Kafka: k}
}

func (web *Websocket) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket is working.....")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade error:", err)
		return
	}
	defer conn.Close()

	go web.Kafka.UserConsumer(conn)
	web.Kafka.BookingConsumer(conn)

}
