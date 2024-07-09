package credentialclient

import (
	"gateway/protos/clientProto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialClient(credentialService_url string) (clientProto.ClientServiceClient) {

	conn, err := grpc.NewClient(credentialService_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect client-service")
	}
	client := clientProto.NewClientServiceClient(conn)
	return client
}
