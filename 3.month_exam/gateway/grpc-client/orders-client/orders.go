package ordersclient

import (
	"gateway/protos/ordersProto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialOrder(orders_url string)ordersProto.OrderServiceClient {

	conn, err := grpc.NewClient(orders_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect order-service")
	}
	order := ordersProto.NewOrderServiceClient(conn)
	return order
}
