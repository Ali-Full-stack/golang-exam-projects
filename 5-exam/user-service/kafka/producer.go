package kafka

import (
	"context"
	"fmt"
	upb "user-service/protos"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/encoding/protojson"
)

func (k *Kafka) ProduceRegistrationEmail(id string, user *upb.UserInfo) error {
	data, err := protojson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal userinfo in KAFKA: %v", err)
	}

	record := kgo.Record{
		Topic: "user-registration",
		Key:   []byte(id),
		Value: data,
	}
	err = k.Client.ProduceSync(context.Background(), &record).FirstErr()
	if err != nil {
		return fmt.Errorf("failed to produce message on Registration: %v", err)
	}
	return nil
}
