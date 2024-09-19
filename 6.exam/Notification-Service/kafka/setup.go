package kafka

import (
  "fmt"

  "github.com/twmb/franz-go/pkg/kgo"
)

type Kafka struct{
	*kgo.Client
}


func ConnectKafka(kurl string, topic string) (*Kafka, error) {
  client, err := kgo.NewClient(
    kgo.SeedBrokers(kurl),
    kgo.ConsumeTopics(topic),
    kgo.ConsumerGroup("my-group"),
  )
  if err != nil {
    return nil, fmt.Errorf("failed consumer client:%v", err)
  }
  return &Kafka{client}, nil
}
