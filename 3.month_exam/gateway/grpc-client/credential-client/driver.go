package credentialclient

import (
	"gateway/protos/driverProto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialDriver(credentialService_url string) driverProto.DriverServiceClient{

	conn, err := grpc.NewClient(credentialService_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc driver to connect driver-service")
	}
	driver := driverProto.NewDriverServiceClient(conn)
	return driver
}
