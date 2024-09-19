package server

import (
	"log"
	"net"
	"os"
	"user-service/internal/service/handler"
	"user-service/proto"
	"user-service/storage"
	"user-service/storage/postgres"

	_ "github.com/joho/godotenv/autoload"
	_"github.com/lib/pq"
	"google.golang.org/grpc"
)

func ConnGrpc() {
	lis, err := net.Listen("tcp", os.Getenv("user_server"))
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

	userService := handler.NewUserService(queries)
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, userService)

	log.Println("Server is listening on port: ", os.Getenv("user_server"))
	if err = s.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
