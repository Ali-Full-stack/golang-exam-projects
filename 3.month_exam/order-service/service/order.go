package service

import (
	"io"
	"log"
	"order-service/internal/repo"
	opb "order-service/protos/order"
)

type OrderService struct {
	opb.UnimplementedOrderServiceServer
	OrderRepo 	*repo.OrderRepo
}

func NewOrderService(o *repo.OrderRepo)*OrderService{
	return &OrderService{OrderRepo: o}
}

func (o *OrderService) CreateAllOrders(stream opb.OrderService_CreateAllOrdersServer)error{
	for {
		req , err :=stream.Recv()
		if err ==io.EOF{
			log.Println("Streaming request finished .")
			break
		}
		if err != nil {
			log.Println("failed to get response !")
			return err
		} 
		resp, err :=o.OrderRepo.MakeClientOrder(req)
		if err !=nil {
			log.Println("Unable to get products:",err)
			return err
		}
		if err :=stream.Send(resp); err != nil {
			log.Println("Failed to send a stream response",err)
		}
	}
	
	return nil
}