package client

import (
	"report-service/proto/budget"
	"report-service/proto/transaction"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialBudgetClient(budget_url string) (budget.BudgetServiceClient, error) {

	conn, err := grpc.NewClient(budget_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect budget-service")
		return nil, err
	}
	client := budget.NewBudgetServiceClient(conn)
	return client, nil
}

func DialTransactionClient(transaction_url string) (transaction.TransactionServiceClient, error) {

	conn, err := grpc.NewClient(transaction_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect transaction-service")
		return nil, err
	}
	client := transaction.NewTransactionServiceClient(conn)
	return client, nil
}
