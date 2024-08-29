package postgres

import (
	"fmt"
	"hotel-service/internal/model"
	hpb "hotel-service/protos"

	"github.com/google/uuid"
)

func (p *Postgres) AddHotelIntoPostgres(req *hpb.HotelInfo) (*hpb.HotelID, error) {
	id := uuid.New().String()

	query := `insert into hotels(id, name, rating, city, region, street)
	values($1, $2, $3, $4, $5, $6)`
	_, err := p.DB.Exec(query, id, req.Name, req.Rating, req.Address.City, req.Address.Region, req.Address.Street)
	if err != nil {
		return nil, fmt.Errorf("failed to insert hotel into database: %v", err)
	}
	return &hpb.HotelID{Id: id}, nil
}

func (p *Postgres) CreateHotelRoomInPostgres(req *hpb.RoomInfo)(*hpb.HotelResponse, error){
	for _, room := range req.Rooms {
		query := `insert into rooms(hotel_id, type, pricePerNight, totalRooms)
		values($1, $2, $3, $4)`
		_, err := p.DB.Exec(query, req.HoteId, room.Type, room.PricePerNight, room.TotalRooms)
		if err != nil {
			return nil, fmt.Errorf("failed to insert room information into database: %v", err)
		}
	}
	return &hpb.HotelResponse{Message: "Room Created Successfully"}, nil
}

func (p *Postgres) UpdateHotelInPostgres(id string, req *hpb.HotelInfo) (*hpb.HotelResponse, error) {
	query := `update hotels
	set name=$1, rating=$2, city=$3, region=$4,  street=$5
	where id=$6
	`
	result, err := p.DB.Exec(query, req.Name, req.Rating, req.Address.City, req.Address.Region, req.Address.Street, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update hotel information: %v", err)
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return nil, fmt.Errorf("hotel does not exist with ID: %v", id)
	}
	return &hpb.HotelResponse{Message: "Hotel updated succesfully"}, nil
}

func (p *Postgres) UpdateHotelRoomInPostgres(req *hpb.RoomInfo) (*hpb.HotelResponse, error) {
	for _, room := range req.Rooms {
		query := `update rooms
		set type=$2, pricePerNight=$3, totalRooms= $4
		where hotel_id=$5
		`
		result, err := p.DB.Exec(query, room.Type, room.PricePerNight, room.TotalRooms, req.HoteId)
		if err != nil {
			return nil, fmt.Errorf("failed to update hotel room information: %v", err)
		}
		if n, _ := result.RowsAffected(); n == 0 {
			return nil, fmt.Errorf("hotel does not exist with ID: %v", req.HoteId)
		}
	}
	return &hpb.HotelResponse{Message: "Hotel Room updated succesfully"}, nil
}

func (p *Postgres) DeleteHotelFromPostgres(req *hpb.HotelID) (*hpb.HotelResponse, error) {
	query := `delete from hotels
	where id = $1`
	result, err := p.DB.Exec(query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete hotel information: %v", err)
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return nil, fmt.Errorf("hotel does not exist with ID: %v", req.Id)
	}
	return &hpb.HotelResponse{Message: "Hotel deleted  succesfully"}, nil
}
func (p *Postgres) DeleteHotelRoomFromPostgres(req *hpb.RoomType)(*hpb.HotelResponse,error){
	query :=`delete from rooms
	where hotel_id = $1 and type = $2`
	result, err := p.DB.Exec(query, req.HotelId, req.Type)
	if err != nil {
		return nil, fmt.Errorf("failed to delete hotel room : %v", err)
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return nil, fmt.Errorf("hotel room does not exist with Hotel ID: %v", req.HotelId)
	}
	return &hpb.HotelResponse{Message: "Hotel Room deleted  succesfully"}, nil
}

func (p *Postgres) GetAllHotelFromPostgres(req *hpb.Empty) ([]*hpb.HotelInfo, error) {
	query := `select name, rating, city, region, street from hotels`
	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all hotels from database : %v", err)
	}
	defer rows.Close()

	var hotels []*hpb.HotelInfo
	for rows.Next() {
		var h model.Hotel
		err := rows.Scan(&h.Name, &h.Rating, &h.City, &h.Region, &h.Street)
		if err != nil {
			return nil, fmt.Errorf("failed to scan hotel information: %v", err)
		}
		hotels = append(hotels, &hpb.HotelInfo{
			Name: h.Name,
			Rating: h.Rating,
			Address: &hpb.Address{
				City: h.City,
				Region: h.Region,
				Street: h.Street,
			},
		})
	}
	return hotels, nil
}

func (p *Postgres) GetHotelByIdFromPostgres(req *hpb.HotelID) (*hpb.HotelWithRoom, error) {
	tx, _ := p.DB.Begin()
	defer tx.Rollback()

	query := `select name, rating, city, region, street from hotels
	where id = $1`

	var hotel  model.HotelInfo
	row := tx.QueryRow(query, req.Id)
	if err := row.Scan(&hotel.Name, &hotel.Rating, &hotel.Address.City, &hotel.Address.Region, &hotel.Address.Street); err != nil {
		return nil, fmt.Errorf("failed to get hotel info: %v", err)
	}
	hotelInfo :=&hpb.HotelWithRoom{
		Name: hotel.Name,
		Rating: hotel.Rating,
		Address: &hpb.Address{
			City: hotel.Address.City,
			Region: hotel.Address.Region,
			Street: hotel.Address.Street,
		},
	}
	query = `select type, pricePerNight, totalRooms from rooms where hotel_id=$1`
	rows, err := tx.Query(query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get rooms information: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var r model.Room
		if err := rows.Scan(&r.Type, &r.PricePerNight, &r.TotalRooms); err != nil {
			return nil, fmt.Errorf("failed to scan room information: %v", err)
		}
		room :=&hpb.Room{
			Type: r.Type,
			PricePerNight: r.PricePerNight,
			TotalRooms: int32(r.TotalRooms),
		}
		hotelInfo.Rooms = append(hotelInfo.Rooms, room)
	}
	return hotelInfo, nil
}

func (p *Postgres) CheckAvailableRoomsInPostgres(req *hpb.RoomCount)(*hpb.RoomResponse, error){
	query :=`select type, pricePerNight, totalRooms from rooms
	where hotel_id=$1 and type=$2 and totalRooms > $3`

	row :=p.DB.QueryRow(query, req.HotelId, req.RoomType, req.Total)
	
	var room model.Room
	if err :=row.Scan(&room.Type, &room.PricePerNight, &room.TotalRooms); err != nil {
		return nil, fmt.Errorf("failed to  check hotel rooms : %v",err)
	}
	return &hpb.RoomResponse{
		Type: room.Type,
		PricePerNight: room.PricePerNight,
		Count: int32(room.TotalRooms),
	}, nil
}

func (p *Postgres) UpdateRoomCountInPostgres(req *hpb.RoomCount)(*hpb.CountResponse, error){
	query :=`update rooms
	set totalRooms = totalRooms + $1
	where hotel_id = $2 and type = $3
	returning totalRooms`

	row :=p.DB.QueryRow(query, req.Total, req.HotelId, req.RoomType)
	var total int32
	if err :=row.Scan(&total); err != nil {
		return nil, fmt.Errorf("failed to update Room count :%v",err)
	}
	return &hpb.CountResponse{
		Type: req.RoomType,
		Count: total,
		Status: "Updated Successfully",
	}, nil
}
