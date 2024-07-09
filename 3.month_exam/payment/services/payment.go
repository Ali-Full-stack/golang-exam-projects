package services

import (
	"context"
	"log"
	"payment/internal/repo"
	pb"payment/protos"
)

type PaymentServer struct {
	pb.UnimplementedPaymentServiceServer
	PaymentRepo *repo.PaymentRepo
}

func NewPaymentServer(pr *repo.PaymentRepo) *PaymentServer {
	return &PaymentServer{PaymentRepo: pr}
}

func (p *PaymentServer) AddClientCard(ctx context.Context, req *pb.CardRequest) (*pb.Empty, error) {

	switch req.Role {
	case "client":
		if err := p.PaymentRepo.AddClientCardToDatabase(req); err != nil {
			return nil, err
		}
	case "driver":
		if err := p.PaymentRepo.AddDriverCardToDatabase(req); err != nil {
			return nil, err
		}
	}
	return &pb.Empty{}, nil
}

func (p *PaymentServer) DeleteClientCard(ctx context.Context, req *pb.RequestId) (*pb.CardResponse, error) {

	switch req.Role {
	case "client":
		if err := p.PaymentRepo.DeleteClientFromDatabase(req); err != nil {
			return nil, err
		}
	case "driver":
		if err := p.PaymentRepo.DeleteDriverFromDatabase(req); err != nil {
			return nil, err
		}
	}
	return &pb.CardResponse{Status: "deleted successfully"}, nil
}



func (p *PaymentServer) MakePurchase(ctx context.Context, req *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	resp, err := p.PaymentRepo.MakeProductPurchase(req)
	if err != nil {
		log.Println("failed to update main account balance:", err)
		return nil, err
	}
	return resp, nil
}

func (p *PaymentServer) MakeOrderPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	resp, err :=p.PaymentRepo.MakeOrderTransaction(req)
	if err != nil {
		log.Println("Transaction Failed !")
		return nil, err
	}
	return resp, nil
}
