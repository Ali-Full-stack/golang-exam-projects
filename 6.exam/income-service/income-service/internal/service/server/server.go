package server

import (
	"log"
	"net"
	"os"
	"transaction-service/internal/service/handler"
	"transaction-service/kafka"
	"transaction-service/proto/income"
	"transaction-service/storage"
	"transaction-service/storage/postgres"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func ConnGrpc() {
	lis, err := net.Listen("tcp", os.Getenv("transactions_server"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	db, err := postgres.ConnectPostgres(os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := storage.New(db)
	kafka, err := kafka.ConnectKafka(os.Getenv("kafka_url"))
	if err != nil {
		log.Fatal(err)
	}
	defer kafka.Close()

	transactionService := handler.NewTransactionService(queries, kafka)

	s := grpc.NewServer()

	income.RegisterTransactionServiceServer(s, transactionService)

	log.Println("Server is listening on port: ", os.Getenv("transactions_server"))
	if err = s.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
