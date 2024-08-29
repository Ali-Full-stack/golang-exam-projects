package service

import (
	mongodb "booking-service/internal/mongoDB"
	"booking-service/kafka"
	bpb "booking-service/protos/booking"
	"booking-service/protos/hotel"
	"booking-service/protos/user"
	"context"
	"fmt"
)

type BookingService struct {
	bpb.UnimplementedBookingServiceServer
	Mongo *mongodb.Mongo
	Kafka *kafka.Kafka
	Hotel hotel.HotelServiceClient
	User  user.UserServiceClient
}

func NewBookingService(m *mongodb.Mongo, k *kafka.Kafka, u user.UserServiceClient, h hotel.HotelServiceClient) *BookingService {
	return &BookingService{Mongo: m, Kafka: k, Hotel: h, User: u}
}

func (b *BookingService) CreateBooking(ctx context.Context, req *bpb.BookingInfo) (*bpb.BookingResponse, error) {
	user, err := b.User.GetUserById(ctx, &user.UserID{Id: req.UserId})
	if err != nil {
		return nil, fmt.Errorf("failed to get user information: %v", err)
	}
	room, err := b.Hotel.CheckAvailableRooms(ctx, &hotel.RoomCount{HotelId: req.HotelId, RoomType: req.RoomType, Total: req.TotalDays})
	if err != nil {
		return nil, fmt.Errorf("failed to check available hotel's  room: %v", err)
	}
	req.TotalAmount = room.PricePerNight * float32(req.TotalDays)
	resp, err := b.Mongo.AddBookingIntoMongoDb(req)
	if err != nil {
		return nil, err
	}
	b.Kafka.ProduceBookingConfirmation(user, resp)
	return resp, nil
}

func (b *BookingService) GetBookingById(ctx context.Context, req *bpb.BookingId) (*bpb.BookingResponse, error) {
	resp, err := b.Mongo.GetBookingByIdFromMongoDB(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BookingService) DeleteBooking(ctx context.Context, req *bpb.BookingId) (*bpb.Response, error) {
	resp, err := b.Mongo.DeleteBookingFromMongoDB(req)
	if err != nil {
		return nil, err
	}
	booking, err := b.Mongo.GetBookingByIdFromMongoDB(req)
	if err != nil {
		return nil, err
	}
	_, err = b.Hotel.UpdateRoomCount(ctx, &hotel.RoomCount{
		HotelId:  booking.HotelId,
		RoomType: booking.RoomType,
		Total:    -(booking.TotalDays),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update room count in DeleteBooking")
	}
	return resp, nil
}
