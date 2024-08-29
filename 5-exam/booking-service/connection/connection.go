package connection

import (
	grpcclient "booking-service/grpc-client"
	mongodb "booking-service/internal/mongoDB"
	"booking-service/kafka"
	"booking-service/pkg"
	"booking-service/protos/booking"
	"booking-service/service"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func GrpcConn() {
	lis, err := net.Listen(os.Getenv("network"), os.Getenv("booking_server"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	kafkaClient, err := kafka.ConnectKafka(os.Getenv("kafka_url"))
	if err != nil {
		log.Fatal(err)
	}
	defer kafkaClient.Client.Close()

	mongoClient, err := mongodb.NewMongoRepo(os.Getenv("mongo_url"))
	if err != nil {
		log.Fatalf("failed to connect mongoDB:%v", err)
	}

	hotelClient := grpcclient.DialHotelClient(os.Getenv("hotel_url"))
	userClient := grpcclient.DialUserClient(os.Getenv("user_url"))

	bookingService := service.NewBookingService(mongoClient, kafkaClient, userClient, hotelClient)

	grpcServer := grpc.NewServer()
	booking.RegisterBookingServiceServer(grpcServer, bookingService)

	go func() {
		pkg.Shutdown(grpcServer)
	}()

	log.Println("Booking-Service: server is listening on port ", os.Getenv("booking_server"))
	log.Fatal(grpcServer.Serve(lis))
}
