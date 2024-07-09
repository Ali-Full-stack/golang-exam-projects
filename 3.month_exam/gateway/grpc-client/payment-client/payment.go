package paymentclient

import (
	"gateway/protos/paymentProto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialPayment(paymentService_url string) paymentProto.PaymentServiceClient {

	conn, err := grpc.NewClient(paymentService_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect payment-service")
	}
	payment := paymentProto.NewPaymentServiceClient(conn)
	return payment
}
