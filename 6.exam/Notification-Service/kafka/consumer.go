package kafka

import (
	"context"
	"fmt"
	"log"
	"notification/protos"
	"notification/storage/mongodb"
	"notification/websocket"
	"os"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/encoding/protojson"
)


func (k *Kafka) NotifyConsumer(mongo *mongodb.MongoRepo) {
	fmt.Println("Notification-service: started Consuming messages.....")
	bookingClient, err := ConnectKafka(os.Getenv("kafka_url"), "budget-exceeding")
	if err != nil {
		log.Fatal(err)
	}

	defer bookingClient.Close()
	for {
		fetches := bookingClient.PollFetches(context.Background())
		if errs := fetches.Errors(); len(errs) > 0 {
			log.Fatal(errs)
		}
		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {

			for _, record := range ftp.Records {
				var message string
				var user_id string
				if len(record.Headers) > 0 {
					message = string(record.Headers[0].Value)
					user_id = record.Headers[0].Key
				}
				var report protos.Report
				err := protojson.Unmarshal(record.Value, &report)
				if err != nil {
					log.Fatalf("failed to unmarshal budget information: %v", err)
				}
				mongo.AddNotifyIntoMongo(user_id, message, &report)
				websocket.NotifyClients(fmt.Sprintf("user_id: %s,  message: %s",user_id, message))
			}
		})
	}
}
