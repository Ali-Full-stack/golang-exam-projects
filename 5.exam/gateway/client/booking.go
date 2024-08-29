package client

import (
	"gateway/protos/booking"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialBookingClient(booking_url string) booking.BookingServiceClient {

	conn, err := grpc.NewClient(booking_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect booking-service")
	}
	booking := booking.NewBookingServiceClient(conn)
	return booking
}
