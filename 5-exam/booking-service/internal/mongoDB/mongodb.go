package mongodb

import (
	"booking-service/internal/model"
	bpb "booking-service/protos/booking"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongoRepo(mongo_url string) (*Mongo, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongo_url))
	if err != nil {
		return nil, fmt.Errorf("error: failed mongoDB connection: %v", err)
	}
	return &Mongo{Client: client}, nil
}

func (m *Mongo) AddBookingIntoMongoDb(req *bpb.BookingInfo)(*bpb.BookingResponse, error){
	bookingCollection :=m.Client.Database("bookings").Collection("booking")

	result, err :=bookingCollection.InsertOne(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed insert booking info into mongoDB:%v",err)
	}
	id :=result.InsertedID.(primitive.ObjectID).Hex()
	return &bpb.BookingResponse{
		BookingId: id,
		UserId: req.UserId,
		HotelId: req.HotelId,
		RoomType: req.RoomType,
		TotalDays: req.TotalDays,
		CheckInDate: req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount: req.TotalAmount,
		Status: "Confirmed",
	},nil
}

func (m *Mongo) GetBookingByIdFromMongoDB(req *bpb.BookingId)(*bpb.BookingResponse,error){
	bookingCollection :=m.Client.Database("bookings").Collection("booking")
	idObj , _:=primitive.ObjectIDFromHex(req.Id)

	result :=bookingCollection.FindOne(context.Background(), bson.M{"_id" : idObj})
	var b model.BookingResponse
	if err :=result.Decode(&b); err != nil {
			return nil, fmt.Errorf("failed to decode  booking details:%v",err)
	}
	return &bpb.BookingResponse{
		BookingId: b.BookingID,
		UserId: b.UserID,
		HotelId: b.HotelID,
		RoomType: b.RoomType,
		TotalDays: int32(b.TotalDays),
		CheckInDate: b.CheckInDate,
		CheckOutDate: b.CheckOutDate,
		TotalAmount: b.TotalAmount,
		Status: b.Status,
	}, nil
}
func (m *Mongo) DeleteBookingFromMongoDB(req *bpb.BookingId)(*bpb.Response, error){
	bookingCollection :=m.Client.Database("bookings").Collection("booking")
	idObj , _:=primitive.ObjectIDFromHex(req.Id)

	result, err :=bookingCollection.DeleteOne(context.Background(), bson.M{"_id": idObj})
	if err != nil {
		return nil,  fmt.Errorf("failed to delete booking from mongoDB: %v", err)
	}
	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("booking does not exist with ID: %v", req.Id)
	}
	return &bpb.Response{Message: "Booking deleted succesfully"}, nil
}