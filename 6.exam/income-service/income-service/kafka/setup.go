package kafka

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Kafka struct {
	*kgo.Client
}

func ConnectKafka(kurl string) (*Kafka, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(kurl),
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed kafka Client error: %v", err)
	}

	err = client.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed kafka connection: %v", err)
	}
	return &Kafka{client}, nil
}