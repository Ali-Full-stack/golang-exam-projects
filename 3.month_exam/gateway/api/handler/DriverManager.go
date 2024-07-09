package handler

import (
	"context"
	"encoding/json"
	"gateway/internal/file"
	"gateway/internal/models"
	"gateway/pkg/code"
	"gateway/pkg/email"
	dpb "gateway/protos/driverProto"
	ppb "gateway/protos/paymentProto"
	"log"
	"net/http"
	"sync"
	"time"
)

type DriverManager struct {
	Driver  dpb.DriverServiceClient
	Payment ppb.PaymentServiceClient
	Log     *log.Logger
}

func NewDriverManager(d dpb.DriverServiceClient, p ppb.PaymentServiceClient, lg *log.Logger) *DriverManager {
	return &DriverManager{Driver: d, Payment: p, Log: lg}
}

func (h *DriverManager) AddNewDriver(w http.ResponseWriter, r *http.Request) {
	h.Log.Println("INFO: Received http Request from driver on CreateNewDriver handler,")

	driver := r.Context().Value("driver")
	httpDriverInfo, ok := driver.(models.DriverInfo)
	if !ok {
		log.Println("failed type assertion:")
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	var grpcDriverInfo = &dpb.DriverInfo{
		Name:          httpDriverInfo.Name,
		Email:         httpDriverInfo.Email,
		Phone:         httpDriverInfo.Phone,
		WorkingRegion: httpDriverInfo.Working_region,
		Vehicle:       httpDriverInfo.Vehicle,
		DriverAddress: &dpb.DriverAddress{
			City:        httpDriverInfo.City,
			Region:      httpDriverInfo.Region,
			HomeAddress: httpDriverInfo.Home_Address,
		},
	}
	resp, err := h.Driver.CreateDriver(ctx, grpcDriverInfo)
	if err != nil {
		h.Log.Println("ERROR: unable to get grpc response from server on CreateDriver service:", err)
		http.Error(w, "Request denied on adding New driver !!", http.StatusInternalServerError)
		return
	}

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		var reqCard = &ppb.CardRequest{
			Id:         resp.Id,
			CardNumber: httpDriverInfo.Card_number,
			Balance:    float32(httpDriverInfo.Balance),
			Role:       "driver",
		}
		_, err := h.Payment.AddClientCard(ctx, reqCard)
		if err != nil {
			log.Println("Unable to add drivers card inforamtion :", err)
		}
	}()

	go func() {
		defer wg.Done()
		err = file.WriteNewUserToFile("internal/data/drivers.json", models.ClientLogin{Id: resp.Id, Email: httpDriverInfo.Email})
		if err != nil {
			log.Println("failed to write new driver to file:", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := email.SendEmail(httpDriverInfo.Email, email.RegisterDriver(resp.Id, httpDriverInfo.Name)); err != nil {
			log.Println("failed to send an email to driver on registration", err)
		}
	}()
	var qrcode []byte
	go func() {
		defer wg.Done()
		qrcode = code.GenerateQrcode(resp.Id)
	}()

	wg.Wait()
	w.Header().Set("Content-type", "image/png")
	w.Write(qrcode)
	h.Log.Println("INFO: Sent  http Response to  driver succesfully on CreateNewDriver handler,")
}

func (h *DriverManager) DeleteDriver(w http.ResponseWriter, r *http.Request) {
	h.Log.Println("INFO:Received   http Request  from  client on Deleting Driver")
	var httpDriverID models.DriverID
	err := json.NewDecoder(r.Body).Decode(&httpDriverID)
	if err != nil {
		http.Error(w, "Invalid Request !! ", http.StatusBadRequest)
		return
	}

	resp, err := h.Driver.DeleteDriver(r.Context(), &dpb.DriverID{Id: httpDriverID.Id})
	if err != nil {
		h.Log.Println("ERROR: Unable to get grpc response from server on DeleteDriver service")
		http.Error(w, "Request Denied !!", http.StatusInternalServerError)
		return
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var DeleteCard = &ppb.RequestId{
			Role: "driver",
			Id:   httpDriverID.Id,
		}
		_, err := h.Payment.DeleteClientCard(r.Context(), DeleteCard)
		if err != nil {
			log.Println("Unable to Delete driver card information :", err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := file.DeleteClientFromFile("internal/data/drivers.json", httpDriverID.Id); err != nil {
			log.Println("Failed to delete driver from file:", err)
		}
	}()
	wg.Wait()
	
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.Log.Println("ERROR:Unable to send http response back to client on DeleteDriver handler:", err)
		http.Error(w, "Request Denied on Deleting Driver !", http.StatusInternalServerError)
		return
	}
	h.Log.Println("INFO:Sent http  response back to client on Deleting Driver")
}
