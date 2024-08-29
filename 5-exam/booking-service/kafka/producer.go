package kafka

import (
	"booking-service/protos/booking"
	"booking-service/protos/user"
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/encoding/protojson"
)

func (k *Kafka) ProduceBookingConfirmation(user *user.UserWithID, req *booking.BookingResponse) error {
	bookingInfo := &booking.BookingEmail{
		BookingId:    req.BookingId,
		UserId:       req.UserId,
		HotelId:      req.HotelId,
		RoomType:     req.RoomType,
		TotalDays:    req.TotalDays,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount:  req.TotalAmount,
		Status:       req.Status,
		Username:     user.Username,
		Email:        user.Email,
	}
	data, err := protojson.Marshal(bookingInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal booking info in KAFKA: %v", err)
	}
	record := kgo.Record{
		Topic: "booking-confirmation",
		Value: data,
	}
	err = k.Client.ProduceSync(context.Background(), &record).FirstErr()
	if err != nil {
		return fmt.Errorf("failed to produce message on Confirmation: %v", err)
	}
	return nil
}
