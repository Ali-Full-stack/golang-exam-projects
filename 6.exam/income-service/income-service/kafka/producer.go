package kafka

import (
	"context"
	"fmt"
	"transaction-service/proto/income"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/encoding/protojson"
)

func (k *Kafka) ExceedingIncomeLimit(message, user_id string, req *income.ReportResponse) error {
	transaction := &income.ReportResponse{
		Income:     req.Income,
		Expenses:   req.Expenses,
		NetSavings: req.NetSavings,
	}
	data, err := protojson.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to marshal booking info in KAFKA: %v", err)
	}
	record := kgo.Record{
		Topic: "budget-exceed",
		Value: data,
		Headers: []kgo.RecordHeader{{
			Key:   user_id,
			Value: []byte(message),
		},
		},
	}
	err = k.Client.ProduceSync(context.Background(), &record).FirstErr()
	if err != nil {
		return fmt.Errorf("failed to produce message on Budget-Exceed: %v", err)
	}
	return nil
}
