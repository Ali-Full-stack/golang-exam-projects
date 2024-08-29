package handler

import (
	"context"
	"encoding/json"
	"gateway/internal/model"
	hpb "gateway/protos/hotel"
	"io"
	"log"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
)

type HotelHandler struct {
	Client hpb.HotelServiceClient
}

func NewHotelHandler(hcl hpb.HotelServiceClient) *HotelHandler {
	return &HotelHandler{Client: hcl}
}
// @Router  				/api/hotels  [post]
// @Summary 			Creates New Hotel
// @Description 		This method used  to create  new  hotel 
// @Security 				BearerAuth
// @Tags					 HOTELS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					body    body    hotel.HotelInfo    true  "Hotel  Details"
// @Success					201 	{object}   hotel.HotelID	"Hotel  ID"
// @Failure					 400 {object} error "Unable to read Hotel information !"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (h *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	var hotel *hpb.HotelInfo
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read request body:", err)
		http.Error(w, "Unable to read Hotel information !", http.StatusBadRequest)
		return
	}
	if err := protojson.Unmarshal(bytes, hotel); err != nil {
		log.Println("Unable to Unmarshal request body:", err)
		http.Error(w, "Invalid  Hotel information !", http.StatusBadRequest)
		return
	}
	resp, err := h.Client.CreateHotel(r.Context(), hotel)
	if err != nil {
		log.Println("failed to get response on Create Hotel : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/hotels/{id}/rooms  [post]
// @Summary 			Creates  Hotel's Room
// @Description 		This method used  to create  hotel's room 
// @Security 				BearerAuth
// @Tags					 ROOMS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					id     path    	string    true    "Hotel ID"
// @Param 					body    body    hotel.RoomInfo    true  "Room  Details"
// @Success					201 	{object}   hotel.HotelResponse	"Room Created Succesfully"
// @Failure					 400 {object} error "Unable to read room  information !"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (h *HotelHandler) CreateHotelRoom(w http.ResponseWriter, r *http.Request) {
	var roomInfo *hpb.RoomInfo
	roomInfo.HoteId = r.PathValue("id")

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read request body:", err)
		http.Error(w, "Unable to read room information !", http.StatusBadRequest)
		return
	}
	if err := protojson.Unmarshal(bytes, roomInfo); err != nil {
		log.Println("Unable to Unmarshal request body:", err)
		http.Error(w, "Invalid  room  information !", http.StatusBadRequest)
		return
	}
	resp, err := h.Client.CreateHotelRoom(r.Context(), roomInfo)
	if err != nil {
		log.Println("failed to get response on Create Hotel's Room : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
// @Router  				/api/hotels/{id}  [put]
// @Summary 			Updates  Hotel  Information
// @Description 		This method used  to update  hotel's Information 
// @Security 				BearerAuth
// @Tags					 HOTELS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					id     path    	string    true    "Hotel ID"
// @Param 					body    body    hotel.HotelInfo    true  "Hotel  Details"
// @Success					201 	{object}   hotel.HotelResponse	"Hotel Updated Succesfully"
// @Failure					 400 {object} error "Unable to read Hotel  information !"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (h *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	var hotel *hpb.HotelInfo
	id := r.PathValue("id")
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read request body:", err)
		http.Error(w, "Unable to read hotel information !", http.StatusBadRequest)
		return
	}
	if err := protojson.Unmarshal(bytes, hotel); err != nil {
		log.Println("Unable to Unmarshal request body:", err)
		http.Error(w, "Invalid  hotel  information !", http.StatusBadRequest)
		return
	}
	ctx := context.WithValue(r.Context(), "id", id)

	resp, err := h.Client.UpdateHotel(ctx, hotel)
	if err != nil {
		log.Println("failed to get response on Update Hotel : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/hotels/{id}/rooms  [put]
// @Summary 			Updates  Hotel's Room
// @Description 		This method used  to update  hotel's room 
// @Security 				BearerAuth
// @Tags					 ROOMS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					id     path    	string    true    "Hotel ID"
// @Param 					body    body    hotel.RoomInfo    true  "Room  Details"
// @Success					201 	{object}   hotel.HotelResponse	"Room Updated Succesfully"
// @Failure					 400 {object} error "Unable to read hotel's room  information !"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (h *HotelHandler) UpdateHotelRoom(w http.ResponseWriter, r *http.Request) {
	var roomInfo *hpb.RoomInfo
	roomInfo.HoteId = r.PathValue("id")
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read request body:", err)
		http.Error(w, "Unable to read hotel's room  information !", http.StatusBadRequest)
		return
	}
	if err := protojson.Unmarshal(bytes, roomInfo); err != nil {
		log.Println("Unable to Unmarshal request body:", err)
		http.Error(w, "Invalid  hotel's room  information !", http.StatusBadRequest)
		return
	}
	resp, err := h.Client.UpdateHotelRoom(r.Context(), roomInfo)
	if err != nil {
		log.Println("failed to get response on Update Hotel's room : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/hotels/{id}  [delete]
// @Summary 			Deletes  Hotel  Information
// @Description 		This method used  to delete  hotel  information 
// @Security 				BearerAuth
// @Tags					 HOTELS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					id     path    	string    true    "Hotel ID"
// @Success					201 	{object}   hotel.HotelResponse	"Hotel Deleted Succesfully"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (h *HotelHandler) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	resp, err := h.Client.DeleteHotel(r.Context(), &hpb.HotelID{Id: id})
	if err != nil {
		log.Println("failed to get response on Delete Hotel : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/hotels/{id}/rooms  [delete]
// @Summary 			Deletes  Hotel's Room
// @Description 		This method used  to Delete  hotel's room 
// @Security 				BearerAuth
// @Tags					 ROOMS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					id     path    	string    true    "Hotel ID"
// @Param 					body    body    model.Room    true  "Room  Details"
// @Success					201 	{object}   hotel.HotelResponse	"Room Deleted Succesfully"
// @Failure					 400 {object} error "Invalid Room type !"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (h *HotelHandler) DeleteHotelRoom(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var room model.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		log.Println("Unable to decode request body on Delete Hotel's room : ", err)
		http.Error(w, "Invalid Room type !", http.StatusBadRequest)
		return
	}
	resp, err := h.Client.DeleteHotelRoom(r.Context(), &hpb.RoomType{HotelId: id, Type: room.Type})
	if err != nil {
		log.Println("failed to get response on Delete Hotel's room : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/hotels  [get]
// @Summary 			Gets  All Hotel's  Information
// @Description 		This method used  to get  all hotel's Information 
// @Security 				BearerAuth
// @Tags					 HOTELS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Success					200 	{object}   []hotel.HotelInfo	"Hotel Updated Succesfully"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (h *HotelHandler) GetAllHotels(w http.ResponseWriter, r *http.Request) {
	stream, err := h.Client.GetAllHotels(r.Context(), &hpb.HotelEmpty{})
	if err != nil {
		log.Println("Unable to stream on Get All hotels:", err)
		http.Error(w, "Unable to get Response !", http.StatusInternalServerError)
		return
	}
	var hotels []*hpb.HotelInfo
	for {
		hotel, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("failed to get streaming messages on Get All hotels: ", err)
			http.Error(w, "Unable to get Response !", http.StatusInternalServerError)
			return
		}
		hotels =append(hotels, hotel)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotels)
}
// @Router  				/api/hotels/{id} [get]
// @Summary 			Gets   Hotel's  Information By ID
// @Description 		This method used  to get   hotel's Information by id 
// @Security 				BearerAuth
// @Tags					 HOTELS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					id    path    string    true    "Hotel ID"
// @Success					200 	{object}   hotel.HotelWithRoom	"Hotel Details"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (h *HotelHandler) GetHotelById(w http.ResponseWriter, r *http.Request) {
	id :=r.PathValue("id")
	resp, err :=h.Client.GetHotelById(r.Context(), &hpb.HotelID{Id: id})
	if err != nil {
		log.Println("failed to get response on Get Hotel Information : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/hotels/{id}/rooms/availability  [post]
// @Summary 			Checks  Hotel's Available Rooms
// @Description 		This method used  to check   hotel's available  rooms 
// @Security 				BearerAuth
// @Tags					 ROOMS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					id     path    	string    true    "Hotel ID"
// @Param 					body    body    model.Room    true  "Room  Details"
// @Success					201 	{object}   hotel.RoomResponse	"Room Details"
// @Failure					 400 {object} error "Invalid Room type !"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (h *HotelHandler) CheckAvailableRooms(w http.ResponseWriter, r *http.Request) {
	var room model.Room
	id :=r.PathValue("id")
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		log.Println("Unable to decode request body on Check Available Hotel's room : ", err)
		http.Error(w, "Invalid Room type !", http.StatusBadRequest)
		return
	}

	resp, err :=h.Client.CheckAvailableRooms(r.Context(), &hpb.RoomCount{HotelId: id, RoomType: room.Type, Total: room.TotalRooms})
	if err != nil {
		log.Println("failed to get response on Check Available Rooms : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Router  				/api/hotels/{id}/rooms  [patch]
// @Summary 			Updates  Hotel's  Rooms Count
// @Description 		This method used  to updates   hotel's   rooms count 
// @Security 				BearerAuth
// @Tags					 ROOMS
// @accept					json
// @Produce				  json
// @Param 					role    header    string    true    "Role"
// @Param 					id     path    	string    true    "Hotel ID"
// @Param 					body    body    model.Room    true  "Room  Details"
// @Success					201 	{object}   hotel.CountResponse	"Room Details"
// @Failure					 400 {object} error "Invalid Room type !"
// @Failure					 500 {object} error  "Unable to get response"
// @Failure					 403 {object} error "Unauthorized access"
func (h *HotelHandler) UpdateRoomCount(w http.ResponseWriter, r *http.Request) {
	var room model.Room
	id :=r.PathValue("id")
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		log.Println("Unable to decode request body on Update Hotel's room : ", err)
		http.Error(w, "Invalid Room type !", http.StatusBadRequest)
		return
	}
	resp, err :=h.Client.UpdateRoomCount(r.Context(), &hpb.RoomCount{HotelId: id, RoomType: room.Type, Total: room.TotalRooms})
	if err != nil {
		log.Println("failed to get response on Update  Room's  count : ", err)
		http.Error(w, "Unable to get response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
