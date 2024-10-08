package connection

import (
	"hotel-service/internal/postgres"
	"hotel-service/kafka"
	"hotel-service/pkg"
	"hotel-service/protos"
	"hotel-service/service"
	"log"
	"net"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func GrpcConn() {
	lis, err := net.Listen(os.Getenv("network"), os.Getenv("hotel_server"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	postgres, err := postgres.ConnectPostgres("postgres", os.Getenv("PG_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer postgres.DB.Close()
	time.Sleep(10 *time.Second)
	kafkaClient, err := kafka.ConnectKafka(os.Getenv("kafka_url"))
	if err != nil {
		log.Fatal(err)
	}
	defer kafkaClient.Client.Close()
	
	hotelService := service.NewHotelService(postgres, kafkaClient)
	grpcServer := grpc.NewServer()
	protos.RegisterHotelServiceServer(grpcServer, hotelService)
	
	go pkg.Shutdown(grpcServer)
		
	log.Println("Hotel-Service: server is listening on port ", os.Getenv("hotel_server"))
	log.Fatal(grpcServer.Serve(lis))
}
