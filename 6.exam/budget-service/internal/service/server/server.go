package server

import (
	"budget-service/internal/kafka"
	"budget-service/internal/service/client"
	"budget-service/internal/service/handler"
	"budget-service/proto/budget"
	"budget-service/storage"
	"budget-service/storage/postgres"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func ConnGrpc() {
	lis, err := net.Listen("tcp", os.Getenv("budget_server"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	db, err := postgres.ConnectPostgres(os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	client, err :=client.DialTransactionClient(os.Getenv("transaction_url"))
	if err != nil {
		log.Fatalf("unable to connect transaction service: %v",err)
	}

	kafka, err :=kafka.ConnectKafka(os.Getenv("kafka_url"))
	if err != nil {
		log.Fatal(err)
	}
	defer kafka.Client.Close()

	queries := storage.New(db)
	budgetService :=handler.NewBudgetService(queries, client, kafka)	

	s := grpc.NewServer()
	budget.RegisterBudgetServiceServer(s, budgetService)

	log.Println("Server is listening on port: ", os.Getenv("budget_server"))
	if err = s.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
