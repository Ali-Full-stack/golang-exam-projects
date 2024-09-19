package client

import (
	"gateway/proto/budget"
	"gateway/proto/users"
	"gateway/proto/income"
	"gateway/proto/notification"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialUsersClient(users_url string) (users.UserServiceClient) {
	conn, err := grpc.NewClient(users_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect users-service: ", err)
	}
	client := users.NewUserServiceClient(conn)
	return client
}

func DialBudgetClient(budget_url string) (budget.BudgetServiceClient) {
	conn, err := grpc.NewClient(budget_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect budget-service: ", err)
	}
	client := budget.NewBudgetServiceClient(conn)
	return client
}

func DialIncomeClient(income_url string) (income.TransactionServiceClient){
	conn, err := grpc.NewClient(income_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect income-service: ", err)
	}
	client := income.NewTransactionServiceClient(conn)
	return client
}


func DialNotifyClient(notification_url string) (notification.NotificationServiceClient){
	conn, err := grpc.NewClient(notification_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect notification-service", err)
	}
	client := notification.NewNotificationServiceClient(conn)
	return client
}

