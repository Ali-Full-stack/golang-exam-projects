package model

type HotelInfo struct {
	Name    string  `json:"name"`
	Rating  float32 `json:"rating"`
	Address Address `json:"Address"`
	Rooms   []Room  `json:"rooms"`
}
type Hotel struct {
	Name   string  `json:"name"`
	Rating float32 `json:"rating"`
	City   string `json:"City"`
	Region string `json:"Region"`
	Street string `json:"street"`
}

type Address struct {
	City   string `json:"City"`
	Region string `json:"Region"`
	Street string `json:"street"`
}

type Room struct {
	Type          string  `json:"type"`
	PricePerNight float32 `json:"PricePerNight"`
	TotalRooms int `json:"totalRooms"`
}
