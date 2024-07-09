package credentialclient

import (
	"gateway/protos/adminProto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialAdmin(credentialService_url string) (adminProto.AdminServiceClient) {

	conn, err := grpc.NewClient(credentialService_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect admin-service")
	}
	admin := adminProto.NewAdminServiceClient(conn)
	return admin
}
