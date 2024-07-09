package handler

import (
	"context"
	"encoding/json"
	"gateway/auth/jwt"
	"gateway/internal/file"
	"gateway/internal/models"
	"gateway/pkg/code"
	"gateway/pkg/email"
	cpb "gateway/protos/clientProto"
	ppb "gateway/protos/paymentProto"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type ClientManager struct {
	Client  cpb.ClientServiceClient
	Payment ppb.PaymentServiceClient
	Log     *log.Logger
}

func NewClientManager(c cpb.ClientServiceClient, p ppb.PaymentServiceClient, lg *log.Logger) *ClientManager {
	return &ClientManager{Client: c, Payment: p, Log: lg}
}

func (h *ClientManager) RegisterNewClient(w http.ResponseWriter, r *http.Request) {
	h.Log.Println("INFO: Received http Request from client on RegisterNewClient handler,")
	clientInfo := r.Context().Value("client")
	httpClientInfo, ok := clientInfo.(models.ClientInfo)
	if !ok {
		log.Println("failed type assertion:")
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	var grpcClientinfo = cpb.ClientInfo{
		Name:  httpClientInfo.Name,
		Email: httpClientInfo.Email,
		Phone: httpClientInfo.Phone,
		Address: &cpb.Address{
			City:        httpClientInfo.City,
			Region:      httpClientInfo.Region,
			HomeAddress: httpClientInfo.Home_address,
		},
	}
	resp, err := h.Client.CreateClient(ctx, &grpcClientinfo)
	if err != nil {
		h.Log.Println("ERROR: unable to get grpc response from server on CreateClient service:", err)
		http.Error(w, "Request denied on registering client !!", http.StatusInternalServerError)
		return
	}
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		var reqCard = &ppb.CardRequest{
			Id:         resp.Id,
			CardNumber: httpClientInfo.Card_number,
			Balance:    float32(httpClientInfo.Balance),
			Role:       "client",
		}
		_, err := h.Payment.AddClientCard(ctx, reqCard)
		if err != nil {
			log.Println("Unable to add clients card inforamtion :", err)
		}
		wg.Done()
	}()
	go func() {
		var cLogin = models.ClientLogin{Id: resp.Id, Email: httpClientInfo.Email}
		err = file.WriteNewUserToFile("internal/data/clients.json", cLogin)
		if err != nil {
			log.Println("failed to write new client to file:", err)
		}
		wg.Done()
	}()

	go func() {
		if err := email.SendEmail(httpClientInfo.Email, email.RegisterClient(resp.Id, httpClientInfo.Name)); err != nil {
			log.Println("failed to send an email to client on registration", err)
		}
		wg.Done()
	}()

	var qrcode []byte
	go func() {
		defer wg.Done()
		qrcode = code.GenerateQrcode(resp.Id)
	}()
	wg.Wait()
	w.Header().Set("Content-type", "image/png")
	w.Write(qrcode)
	h.Log.Println("INFO: Sent  http Response to  client succesfully on RegisterNewClient handler,")

}

func (h *ClientManager) ClientLogin(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email")
	id := r.Context().Value("id")
	token := jwt.GenerateToken("order", email.(string), id.(string), os.Getenv("secret_key"))

	w.Header().Set("Content_type", "application/json")
	w.Write([]byte(token))
}

func (h *ClientManager) DeleteClient(w http.ResponseWriter, r *http.Request) {
	h.Log.Println("INFO: Received http Request from client on DeleteClient handler,")
	var httpClientId models.ClientID
	err := json.NewDecoder(r.Body).Decode(&httpClientId)
	if err != nil {
		http.Error(w, "Invalid Request !! ", http.StatusBadRequest)
		return
	}
	resp, err := h.Client.DeleteClient(r.Context(), &cpb.ClientID{Id: httpClientId.Id})
	if err != nil {
		h.Log.Println("ERROR: Unable to get grpc response from server on DeleteClient service")
		http.Error(w, "Request denied on deleting client !!", http.StatusInternalServerError)
		return
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var DeleteCard = &ppb.RequestId{
			Role:       "client",
			Id: httpClientId.Id,
		}
		_, err := h.Payment.DeleteClientCard(r.Context(), DeleteCard)
		if err != nil {
			log.Println("Unable to Delete clients card information :", err)
		}
	}()
	go func (){
		defer wg.Done()
		if err := file.DeleteClientFromFile("internal/data/clients.json", httpClientId.Id); err != nil {
			log.Println(err)
		}
	}()
	wg.Wait()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.Log.Println("ERROR:Unable to send http response back on DeleteClient  handler")
		http.Error(w, "Request denied on deleting client!!!", http.StatusInternalServerError)
		return
	}
	h.Log.Println("INFO: Sent  http Response to  client succesfully on DeleteClient handler,")
}
