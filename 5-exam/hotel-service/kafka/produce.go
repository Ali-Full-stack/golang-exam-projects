package kafka

import (
	"context"
	"fmt"
	"hotel-service/protos"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/encoding/protojson"
)

func (k *Kafka) ProduceRoomAvailability(room *protos.RoomCount) error {
	data, err := protojson.Marshal(room)
	if err != nil {
		return fmt.Errorf("failed to marshal room details in KAFKA: %v", err)
	}
	record := kgo.Record{
		Topic: "room-availability",
		Value: data,
	}
	err = k.Client.ProduceSync(context.Background(), &record).FirstErr()
	if err != nil {
		return fmt.Errorf("failed to produce message on Room Availability: %v", err)
	}
	return nil
}
