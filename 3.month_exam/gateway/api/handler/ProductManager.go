package handler

import (
	"encoding/json"
	"gateway/internal/models"
	ppb "gateway/protos/paymentProto"
	pb "gateway/protos/productProto"
	"io"
	"log"
	"net/http"
)

type ProductManager struct {
	Product pb.ProductServiceClient
	Payment ppb.PaymentServiceClient
	Log     *log.Logger
}

func NewProductManager(p pb.ProductServiceClient, pc ppb.PaymentServiceClient, lg *log.Logger) *ProductManager {
	return &ProductManager{Product: p, Payment: pc, Log: lg}
}

func (p *ProductManager) AddNewProducts(w http.ResponseWriter, r *http.Request) {
	p.Log.Println("INFO: Received http request from client on AddNewProducts handler .")
	var listProducts []models.ProductInfo
	if err := json.NewDecoder(r.Body).Decode(&listProducts); err != nil {
		http.Error(w, "invalid Request !", http.StatusBadRequest)
		return
	}

	stream, err := p.Product.CreateProducts(r.Context())
	if err != nil {
		p.Log.Println("ERROR:Unable to client stream new products to order-service !", err)
		http.Error(w, "Request denied !!!", http.StatusInternalServerError)
		return
	}
	for _, v := range listProducts {
		response := &pb.ProductInfo{
			Name:     v.Name,
			Category: v.Category,
			Quantity: v.Quantity,
			Price:    float32(v.Price),
		}
		if err := stream.Send(response); err != nil {
			log.Println("failed to send a response to server !", err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("failed to receive response from server on CreateProducts")
		http.Error(w, "unable to make purchase !!", http.StatusInternalServerError)
		return
	}

	payResp, err := p.Payment.MakePurchase(r.Context(), &ppb.PurchaseRequest{Amount: resp.TotalAmount})
	if err != nil {
		http.Error(w, "Request Denied !", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(payResp); err != nil {
		http.Error(w, "Product Purchase Denied !", http.StatusInternalServerError)
		return
	}
	p.Log.Println("INFO: Product purchase is succesfully Done !")
}

func (p *ProductManager) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	var request models.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid Request !", http.StatusBadRequest)
		return
	}

	var listProducts []models.ProductInfo
	req := pb.CategoryRequest{CategoryName: request.Category}
	stream, err := p.Product.GetAllProducts(r.Context(), &req)
	if err != nil {
		log.Println("failed  grpc streaming on GetAllProducts !")
		http.Error(w, "Something went wrong !", http.StatusInternalServerError)
		return
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("unable to get stream response from server:", err)
			http.Error(w, "Category does not exists !", http.StatusInternalServerError)
			return
		}

		listProducts = append(listProducts, models.ProductInfo{
			Name:      resp.Name,
			Category:  resp.Category,
			Quantity:  resp.Quantity,
			Price:     float64(resp.Price),
			CreatedAt: resp.CreatedAt,
			ExpiredAt: resp.ExpiredAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(listProducts); err != nil {
		log.Println("failed to encode listproducts:")
		http.Error(w, "Request Denied !", http.StatusInternalServerError)
		return
	}

}
