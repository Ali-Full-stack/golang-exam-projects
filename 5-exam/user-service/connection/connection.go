package connection

import (
	"log"
	"net"
	"os"
	"user-service/internal/postgres"
	"user-service/internal/redisdb"
	"user-service/kafka"
	"user-service/pkg"
	"user-service/protos"
	"user-service/service"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func GrpcConn() {
	lis, err := net.Listen(os.Getenv("network"), os.Getenv("user_server"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	postgres, err := postgres.ConnectPostgres(os.Getenv("db_driver"), os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal(err)
	}
	defer postgres.DB.Close()

	redisClient := redisdb.ConnectRedis(os.Getenv("redis_url"))

	kafkaClient, err := kafka.ConnectKafka(os.Getenv("kafka_url"))
	if err != nil {
		log.Fatal(err)
	}
	defer kafkaClient.Client.Close()

	userService := service.NewUserService(redisClient, postgres, kafkaClient)

	grpcServer := grpc.NewServer()
	protos.RegisterUserServiceServer(grpcServer, userService)

	go func() {
		pkg.Shutdown(grpcServer)
	}()

	log.Println("User-Service: server is listening on port ", os.Getenv("user_server"))
	log.Fatal(grpcServer.Serve(lis))
}
