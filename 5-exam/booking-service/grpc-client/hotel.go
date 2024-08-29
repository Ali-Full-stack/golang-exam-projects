package grpcclient

import (
	"booking-service/protos/hotel"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialHotelClient(hotel_url string) hotel.HotelServiceClient {
	conn, err := grpc.NewClient(hotel_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect hotel-service")
	}
	hotel := hotel.NewHotelServiceClient(conn)
	return hotel
}
