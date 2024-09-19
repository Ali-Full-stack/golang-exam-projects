package client

import (
	"budget-service/proto/transaction"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialTransactionClient(transaction_url string) (transaction.ReportServiceClient, error) {
	conn, err := grpc.NewClient(transaction_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect transactions-service")
		return nil, err
	}
	client := transaction.NewReportServiceClient(conn)
	return client, nil
}
