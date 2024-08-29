package grpcclient

import (
	"booking-service/protos/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialUserClient(user_url string) user.UserServiceClient {

	conn, err := grpc.NewClient(user_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect user-service")
	}
	user := user.NewUserServiceClient(conn)
	return user
}
