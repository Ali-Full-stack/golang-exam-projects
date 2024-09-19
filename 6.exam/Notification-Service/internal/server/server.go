package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"notification/internal/handler"
	"notification/kafka"
	"notification/protos"
	"notification/storage/mongodb"
	"notification/websocket"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func ConnGrpc() {

	mongoRepo, err := mongodb.NewMongoRepo(os.Getenv("mongo_url"))
	if err != nil {
		log.Fatalf("failed to connect mongoDB: %v", err)
	}
	defer mongoRepo.Client.Disconnect(context.Background())

	notifyService := handler.NewNotficationService(mongoRepo)

	kafkaClient, err := kafka.ConnectKafka(os.Getenv("kafka_url"), "budget-exceeding")
	if err != nil {
		log.Fatalf("failed to connect kafka:%v", err)
	}
	defer kafkaClient.Close()

	go kafkaClient.NotifyConsumer(mongoRepo)

	go func() {
		lis, err := net.Listen("tcp", os.Getenv("notify_url"))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		protos.RegisterNotificationServiceServer(grpcServer, notifyService)

		log.Println("Notification-service: started gRPC server on port:", os.Getenv("notify_url"))
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	http.HandleFunc("/ws", websocket.HandleWebSocket)
	log.Println("WebSocket server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
