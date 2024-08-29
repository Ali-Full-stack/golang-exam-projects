package service

import (
	"context"
	"fmt"
	"hotel-service/internal/postgres"
	"hotel-service/kafka"
	hpb "hotel-service/protos"
)

type HotelService struct {
	hpb.UnimplementedHotelServiceServer
	Postgres *postgres.Postgres
	Kafka *kafka.Kafka
}

func NewHotelService(p *postgres.Postgres, k *kafka.Kafka) *HotelService {
	return &HotelService{Postgres: p, Kafka: k}
}

func (h *HotelService) CreateHotel(ctx context.Context, req *hpb.HotelInfo) (*hpb.HotelID, error) {

	response, err := h.Postgres.AddHotelIntoPostgres(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (h *HotelService) CreateHotelRoom(ctx context.Context, req *hpb.RoomInfo) (*hpb.HotelResponse, error) {

	response, err := h.Postgres.CreateHotelRoomInPostgres(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (h *HotelService) UpdateHotel(ctx context.Context, req *hpb.HotelInfo) (*hpb.HotelResponse, error) {
	id := ctx.Value("id").(string)
	response, err := h.Postgres.UpdateHotelInPostgres(id, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (h *HotelService) UpdateHotelRoom(ctx context.Context, req *hpb.RoomInfo) (*hpb.HotelResponse, error) {
	response, err := h.Postgres.UpdateHotelRoomInPostgres(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (h *HotelService) DeleteHotel(ctx context.Context, req *hpb.HotelID) (*hpb.HotelResponse, error) {
	response, err := h.Postgres.DeleteHotelFromPostgres(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (h *HotelService) DeleteHotelRoom(ctx context.Context, req *hpb.RoomType) (*hpb.HotelResponse, error) {
	response, err := h.Postgres.DeleteHotelRoomFromPostgres(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (h *HotelService) GetAllHotel(req *hpb.Empty, stream hpb.HotelService_GetAllHotelsServer) error {
	hotels, err := h.Postgres.GetAllHotelFromPostgres(req)
	if err != nil {
		return err
	}
	for _, h := range hotels {
		err := stream.Send(h)
		if err != nil {
			return fmt.Errorf("unable to send streaming message: %v", err)
		}
	}
	return nil
}

func (h *HotelService) GetHotelById(ctx context.Context, req *hpb.HotelID) (*hpb.HotelWithRoom, error) {
	response, err :=h.Postgres.GetHotelByIdFromPostgres(req)
	if err != nil {
		return nil, err
	}
	return response , nil
}

func (h *HotelService) CheckAvailableRooms(ctx context.Context, req *hpb.RoomCount)(*hpb.RoomResponse,error){
	response, err :=h.Postgres.CheckAvailableRoomsInPostgres(req)
	if err != nil {
		return nil, err
	}
	return response , nil
}

func (h *HotelService) UpdateRoomCount(ctx context.Context, req *hpb.RoomCount)(*hpb.CountResponse,error){
	response, err :=h.Postgres.UpdateRoomCountInPostgres(req)
	if err != nil {
		return nil, err
	}
	if req.Total < 0 {
		h.Kafka.ProduceRoomAvailability(&hpb.RoomCount{
			HotelId: req.HotelId,
			RoomType: req.RoomType,
			Total: response.Count,
		 })
	}
	return response , nil
}