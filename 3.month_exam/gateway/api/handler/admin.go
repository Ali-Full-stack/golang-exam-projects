package handler

import (
	"context"
	"encoding/json"
	"gateway/internal/file"
	"gateway/internal/models"
	"gateway/pkg/email"
	apb "gateway/protos/adminProto"
	"log"
	"net/http"
	"time"
)

type AdminHandler struct {
	Client apb.AdminServiceClient
	Log    *log.Logger
}

func NewAdminHandler(ad apb.AdminServiceClient, lg *log.Logger) *AdminHandler {
	return &AdminHandler{Client: ad, Log: lg}
}

func (a *AdminHandler) AddNewAdmin(w http.ResponseWriter, r *http.Request) {
	a.Log.Println("INFO: Received http Request from client on AddNewAdmin handler,")

	var httpAdminInfo models.AdminInfo
	if err := json.NewDecoder(r.Body).Decode(&httpAdminInfo); err != nil {
		http.Error(w, "Invalid Request !! ", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	grpcAdminInfo := &apb.AdminInfo{
		Name:     httpAdminInfo.Name,
		Email:    httpAdminInfo.Email,
		Password: httpAdminInfo.Password,
	}

	resp, err := a.Client.AddNewAdmin(ctx, grpcAdminInfo)
	if err != nil {
		a.Log.Println("ERROR: Unable to connect grpcServer on AddNewAdmin service:", err)
		http.Error(w, "Request denied on Adding New Admin !!", http.StatusInternalServerError)
		return
	}

	a.Log.Println("INFO: new admin joined : name:", httpAdminInfo.Name, "role", "normal")

	if err := file.WriteNewAdminToFile("internal/data/admins.json", models.AdminLogin{Id: resp.Id, HashPassword: resp.HashPassword}); err != nil {
		log.Println("unable to write new admin to file", err)
	}

	if err := email.SendEmail(httpAdminInfo.Email, email.SendNewAdmin(resp.Id, httpAdminInfo.Name)); err != nil {
		log.Println("unable to send an email to new admin:", err)
	}

	var adResponse = models.AdminResponse{Status: resp.Status}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(adResponse); err != nil {
		a.Log.Println("ERROR:Unable to send http response back on AddNewAdmin handler")
		http.Error(w, "Request denied !!!", http.StatusInternalServerError)
		return
	}
	a.Log.Println("INFO: Sent  http Response to  admin succesfully on AddNewAdmin handler,")

}
