package ordersclient

import (
	"gateway/protos/productProto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Dialproduct(orders_url string) productProto.ProductServiceClient {

	conn, err := grpc.NewClient(orders_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect product-service")
	}
	product := productProto.NewProductServiceClient(conn)
	return product
}
