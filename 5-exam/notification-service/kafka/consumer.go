package kafka

import (
	"context"
	"fmt"
	"log"
	"notify-service/pkg"
	"notify-service/protos"

	"github.com/gorilla/websocket"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/encoding/protojson"
)
type Notify struct{
	Kafka *kgo.Client
}

func NewNotify(k *kgo.Client)*Notify{
	return &Notify{Kafka: k}
}

func (n *Notify) UserConsumer(conn *websocket.Conn) {
	fmt.Println("User-Registration: started Consuming messages.....")
	for {
		fetches := n.Kafka.PollFetches(context.Background())
		if errs := fetches.Errors(); len(errs) > 0 {
			log.Fatal(errs)
		}
		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			for _, record := range ftp.Records {
				id := record.Key
				var user protos.UserInfo
				err := protojson.Unmarshal(record.Value, &user)
				if err != nil {
					log.Fatalf("failed to unmarshal user information: %v", err)
				}
				pkg.RegistrationEmail(string(id), &user)
				conn.WriteMessage(websocket.TextMessage, record.Value)
			}
		})
	}
}
func(n *Notify)  BookingConsumer(conn *websocket.Conn){
	fmt.Println("Booking-Confirmation: started Consuming messages.....")
	for {
		fetches := n.Kafka.PollFetches(context.Background())
		if errs := fetches.Errors(); len(errs) > 0 {
			log.Fatal(errs)
		}
		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			for _, record := range ftp.Records {
				var booking protos.BookingEmail
				err := protojson.Unmarshal(record.Value, &booking)
				if err != nil {
					log.Fatalf("failed to unmarshal booking information: %v", err)
				}
				pkg.BookingConfirmationEmail(&booking)
				conn.WriteMessage(websocket.TextMessage , record.Value)
			}
		})
	}
}