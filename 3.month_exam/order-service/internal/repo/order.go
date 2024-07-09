package repo

import (
	"database/sql"
	"log"
	opb "order-service/protos/order"
)

type OrderRepo struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{DB: db}
}

func (o *OrderRepo) MakeClientOrder(req *opb.OrderRequest)(*opb.OrderResponse, error){
	query :=`
		select price from products where name = $1
	`
	var resp opb.OrderResponse 
	err :=o.DB.QueryRow(query, req.ProductName).Scan(&resp.Price)
	if err != nil {
		log.Println("failed to finding products from database",err)
		return nil, err
	}
	resp.TotalAmount =resp.Price * float32(req.Quantity)
	resp.ProductName = req.ProductName
	resp.Quantity = req.Quantity
	return &resp, nil
}
