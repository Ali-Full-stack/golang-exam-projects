package models

type OrderRequest struct {
	Product_name string `json:"name"`
	Quantity     int32  `json:"quantity"`
}

type OrderResponse struct {
	DriverInfo 	DriverContact 	`json:"driverinfo"`
	Products 	[]Order  `json:"products"`
	Payment    Payment `json:"payment"`
}
type DriverContact struct {
	Name    string `json:"name"`
	Email 	string `json:"email"`
	Phone   string `json:"phone"`
	Vehicle string `json:"vehicle"`
}

type Clientcontact struct{
	Phone string `json:"phone"`
	Email string  `json:"email"`
	City string 	`json:"city"`
	Region 	string `json:"region" `
	HomeAddress string `` 	
}

type Order struct{
	Name 	string `json:"name"`
	Quantity int32 `json:"quantity"`
	Price 	float32 `json:"price"`
	Total  float32 `json:"total"`
}
type Payment struct {
	Discount int32 `json:"discount"`
	DiscountAmount float32 `json:"discount_amount"` 
	TotalAmount float32 `json:"total_amount"`
}

