package models

type DriverInfo struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Working_region string `json:"working_region"`
	Vehicle        string `json:"vehicle"`
	Status         string `json:"status"`
	Hired_at       string `json:"hired_at"`
	City           string `json:"city"`
	Region         string `json:"region"`
	Home_Address   string `json:"home_address"`
	Card_number string `json:"card_number"`
	Balance	float64 	`json:"balance"` 
}

type DriverID struct {
	Id string `json:"id"`
}

type DriverResponse struct {
	Status string `json:"status"`
}


