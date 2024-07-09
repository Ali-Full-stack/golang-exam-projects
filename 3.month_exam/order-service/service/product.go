package service

import (
	"io"
	"log"
	"order-service/internal/repo"
	pb "order-service/protos/product"
	"time"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	Prepo *repo.ProductRepo
}

func NewProductServer(pr *repo.ProductRepo) *ProductServer {
	return &ProductServer{Prepo: pr}
}

func (p *ProductServer) CreateProducts(stream pb.ProductService_CreateProductsServer) error {
	var total float32
	for {
		product, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.ProductResponse{Status: "success", TotalAmount: total})
			}
			return err
		}
		if _, err := p.Prepo.AddNewProductToDatabase(product); err != nil {
			log.Println("failed to add new product to database:", err)
			return err
		}
		total = product.Price * float32(product.Quantity)
	}
}

func (p *ProductServer) GetAllProducts(req *pb.CategoryRequest, stream pb.ProductService_GetAllProductsServer) error {

	listProducts, err := p.Prepo.GetProductsByCategory(req)
	if err != nil {
		log.Println("failed getting products from database", err)
		return err
	}
	for _, product := range listProducts {
		if err := stream.Send(product); err != nil {
			log.Println("Unable to send a stream response :", err)
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}
