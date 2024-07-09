package handler

import (
	"encoding/json"
	"fmt"
	"gateway/internal/models"
	"gateway/pkg/email"
	cpb "gateway/protos/clientProto"
	dpb "gateway/protos/driverProto"
	opb "gateway/protos/ordersProto"
	ppb "gateway/protos/paymentProto"
	"io"
	"log"
	"net/http"
	"sync"
)

type OrderManager struct {
	Client  cpb.ClientServiceClient
	Driver  dpb.DriverServiceClient
	Payment ppb.PaymentServiceClient
	Order   opb.OrderServiceClient
	Log     *log.Logger
}

func NewOrderManager(c cpb.ClientServiceClient, d dpb.DriverServiceClient, p ppb.PaymentServiceClient, o opb.OrderServiceClient, lg *log.Logger) *OrderManager {
	return &OrderManager{Client: c, Driver: d, Payment: p, Order: o, Log: lg}
}

func (o *OrderManager) CreateOrders(w http.ResponseWriter, r *http.Request) {
	o.Log.Println("INFO: received httprequest on CreateOrders handler ")
	clientId := r.Context().Value("id")

	var listOrders []models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&listOrders); err != nil {
		http.Error(w, "Invalid request!", http.StatusBadRequest)
		return
	}
	fmt.Println(listOrders)
	var totalAmount float32
	var listProducts []models.Order
	var orderResponse models.OrderResponse

	stream, err := o.Order.CreateAllOrders(r.Context())
	if err != nil {
		o.Log.Println("ERROR: failed gRPC streaming on CreateAllOrders:", err)
		http.Error(w, "Request denied!", http.StatusInternalServerError)
		return
	}

	for _, v := range listOrders {
		order := &opb.OrderRequest{
			ProductName: v.Product_name,
			Quantity:    v.Quantity,
		}
		fmt.Println(order)
		err = stream.Send(order)
		if err != nil {
			log.Println("Unable to send a stream request to server:", err)
			return
		}
	}

	err = stream.CloseSend()
	if err != nil {
		log.Println("Failed to close stream send:", err)
		return
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Cannot get response from gRPC server:", err)
			return
		}

		listProducts = append(listProducts, models.Order{
			Name:     resp.ProductName,
			Quantity: resp.Quantity,
			Price:    resp.Price,
			Total:    resp.TotalAmount,
		})

		totalAmount += resp.TotalAmount
	}

	orderResponse.Products = listProducts

	resp, err := o.Client.GetClientLocation(r.Context(), &cpb.ClientID{Id: clientId.(string)})
	if err != nil {
		log.Println("Unable to get client locations !")
		http.Error(w, "request denied !", http.StatusInternalServerError)
		return
	}
	clientContact := models.Clientcontact{Phone: resp.Phone, Email: resp.Email, City: resp.Address.City, Region: resp.Address.Region, HomeAddress: resp.Address.HomeAddress}

	driverInfo, err := o.Driver.GetAvailableDriver(r.Context(), &dpb.GetLocationRequest{Location: clientContact.Region})
	if err != nil {
		log.Println("Unable to get available drivers !")
		http.Error(w, "request denied !", http.StatusInternalServerError)
		return
	}
	driverContact := models.DriverContact{Name: driverInfo.Name, Email: driverInfo.Email, Phone: driverInfo.Phone, Vehicle: driverInfo.Vehicle}
	orderResponse.DriverInfo = driverContact
	var wg sync.WaitGroup
	wg.Add(3)
	var drivercharge float32
	go func() {
		defer wg.Done()
		paymentResp, err := o.Payment.MakeOrderPayment(r.Context(), &ppb.PaymentRequest{ClientId: resp.Id, DriverId: driverInfo.Id, TotalAmount: totalAmount, Discount: 8})
		if err != nil {
			log.Println("Unable to get available drivers !")
			http.Error(w, "request denied !", http.StatusInternalServerError)
			return
		}
		drivercharge = paymentResp.DriverAmount
		paymentInfo := models.Payment{Discount: 8, DiscountAmount: paymentResp.DiscountAmount, TotalAmount: paymentResp.TotalWithDiscount}
		orderResponse.Payment = paymentInfo
	}()
	go func() {
		defer wg.Done()
		if err := email.SendEmail(clientContact.Email, email.SendOrderConfirmationToClient()); err != nil {
			log.Println("failed to send confirmation response to client")
		}
	}()
	go func() {
		defer wg.Done()
		if err := email.SendEmail(driverContact.Email, email.SendDriverConfirmationOrder(clientContact, driverContact.Name, driverInfo.Id, drivercharge)); err != nil {
			log.Println("failed to send email to driver on order confirmation!", err)
		}
	}()
	wg.Wait()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orderResponse); err != nil {
		http.Error(w, "Order request denied ! something went wrong ", http.StatusInternalServerError)
		return
	}

}
