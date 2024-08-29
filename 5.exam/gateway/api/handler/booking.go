package handler

import (
	"encoding/json"
	bpb "gateway/protos/booking"
	"io"
	"log"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
)

type BookingHandler struct {
	Client bpb.BookingServiceClient
}

func NewBookingHandler(bcl bpb.BookingServiceClient) *BookingHandler {
	return &BookingHandler{Client: bcl}
}
// @Router  				/api/bookings  [post]
// @Summary 			Creates New Booking
// @Description 		This method used  to create new  booking 
// @Security 				BearerAuth
// @Tags					 BOOKINGS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					body    body    booking.BookingInfo    true  "Booking Details"
// @Success					201 	{object}   booking.BookingResponse		"Booking  Confirmation"
// @Failure					 400 {object} error "Unable to read booking  details !"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (b *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var bookingInfo *bpb.BookingInfo
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read request body:", err)
		http.Error(w, "Unable to read booking  details !", http.StatusBadRequest)
		return
	}
	if err := protojson.Unmarshal(bytes, bookingInfo); err != nil {
		log.Println("Unable to Unmarshal request body:", err)
		http.Error(w, "Invalid  booking  details !", http.StatusBadRequest)
		return
	}
	resp, err := b.Client.CreateBooking(r.Context(), bookingInfo)
	if err != nil {
		log.Println("failed to get response on Create booking : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/bookings/{id}  [delete]
// @Summary 			Deletes  Booking Details
// @Description 		This method used  to delete   booking detail 
// @Security 				BearerAuth
// @Tags					 BOOKINGS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					id    	path    string    true    "Booking ID"
// @Success					201 	{object}   booking.Response		"Deleted Succesfully"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (b *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	id :=r.PathValue("id")
	resp, err :=b.Client.DeleteBooking(r.Context(), &bpb.BookingId{Id: id})
	if err != nil {
		log.Println("failed to get response on Delete booking : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/bookings/{id}  [get]
// @Summary 			gets  Booking Details By ID
// @Description 		This method used  to get   booking detail 
// @Security 				BearerAuth
// @Tags					 BOOKINGS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					id    	path    string    true    "Booking ID"
// @Success					201 	{object}   booking.Response		"Deleted Succesfully"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (b *BookingHandler) GetBookingById(w http.ResponseWriter, r *http.Request) {
	id :=r.PathValue("id")
	resp, err :=b.Client.GetBookingById(r.Context(), &bpb.BookingId{Id: id})
	if err != nil {
		log.Println("failed to get response on Get booking : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
