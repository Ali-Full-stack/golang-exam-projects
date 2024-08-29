package client

import (
	"gateway/protos/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialClient(user_url string)user.UserServiceClient {

	conn, err := grpc.NewClient(user_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect client-service")
	}
	client := user.NewUserServiceClient(conn)
	return client
}
