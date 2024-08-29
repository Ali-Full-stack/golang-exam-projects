package model

type HotelInfo struct {
	Name    string  `json:"name"`
	Rating  float64 `json:"rating"`
	Address Address `json:"address"`
}

type Address struct {
	City   string `json:"city"`
	Region string `json:"region"`
	Street string `json:"street"`
}

type Room struct {
	Type         string  `json:"type"`
	PricePerNight float64 `json:"price_per_night"`
	TotalRooms   int32   `json:"total_rooms"`
}