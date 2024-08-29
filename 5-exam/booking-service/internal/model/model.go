package model

type BookingInfo struct {
	UserID      string `json:"user_id" bson:"user_id"`
	HotelID     string `json:"hotel_id" bson:"hotel_id"`
	RoomType    string `json:"roomType" bson:"roomType"`
	CheckInDate string `json:"checkInDate" bson:"checkInDate"`
	TotalDays  int         `json:"totalDays" bson:"totalDays"`
	CheckOutDate string `json:"checkOutDate" bson:"checkOutDate"`
	TotalAmount string `json:"totalAmount" bson:"totalAmount"`
}

type BookingResponse struct {
	BookingID     string `json:"booking_id" bson:"booking_id"`
	UserID        string `json:"user_id" bson:"user_id"`
	HotelID       string `json:"hotel_id" bson:"hotel_id"`
	RoomType      string `json:"roomType" bson:"roomType"`
	CheckInDate   string `json:"checkInDate" bson:"checkInDate"`
	CheckOutDate  string `json:"checkOutDate" bson:"checkOutDate"`
	TotalDays  int         `json:"totalDays" bson:"totalDays"`
	TotalAmount   float32 `json:"totalAmount" bson:"totalAmount"`
	Status        string `json:"status" bson:"status"`
}

type BookingId struct {
	ID string `json:"id" bson:"id"`
}

type Response struct {
	Message string `json:"message" bson:"message"`
}