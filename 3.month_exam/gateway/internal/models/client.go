package models

type ClientInfo struct {
	Name         string  `json:"name" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Phone        string  `json:"phone" binding:"required"`
	City         string  `json:"city" binding:"required"`
	Region       string  `json:"region" binding:"required"`
	Home_address string  `json:"home_address" binding:"required"`
	Card_number  string  `json:"card_number" binding:"required"`
	Balance      float64 `json:"balance" binding:"required"`
}

type ClientID struct {
	Id string `json:"id"`
}

type ClientResponse struct {
	Status string `json:"status"`
}

type ClientLogin struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
