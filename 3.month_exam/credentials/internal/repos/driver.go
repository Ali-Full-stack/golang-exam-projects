package repos

import (
	dpb "credentials/protos/driverpb"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type DriverRepo struct {
	DB *sql.DB
}

func NewDriverRepo(db *sql.DB) *DriverRepo {
	return &DriverRepo{DB: db}
}

func (dr *DriverRepo) AddNewDriverToDatabase(req *dpb.DriverInfo) (*dpb.DriverID, error) {
	tx, _ := dr.DB.Begin()
	defer tx.Rollback()

	req.Id = uuid.New().String()
	req.HiredAt = time.Now().Format(time.ANSIC)
	req.Status = "active"
	query := `
		insert into drivers (id, name, email, phone, working_region, vehicle, status, hired_at)
		values($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := tx.Exec(query, req.Id, req.Name, req.Email, req.Phone, req.WorkingRegion, req.Vehicle, req.Status, req.HiredAt)
	if err != nil {
		log.Println("failed to add new driver to database:", err)
		return nil, err
	}
	query = `
		insert into driver_location(driver_id, city, region, home_address)
		values($1, $2, $3, $4)
	`
	_, err = tx.Exec(query, req.Id, req.DriverAddress.City, req.DriverAddress.Region, req.DriverAddress.HomeAddress)
	if err != nil {
		log.Println("failed to add driver's location to database:", err)
		return nil, err
	}
	tx.Commit()
	return &dpb.DriverID{Id: req.Id}, nil
}

func (dr *DriverRepo) DeleteDriverFromDatabase(req *dpb.DriverID) (*dpb.DriverResponse, error) {
	tx, _ := dr.DB.Begin()
	defer tx.Rollback()
	query := `
		delete from driver_location
		where driver_id = $1
	`
	n, err := tx.Exec(query, req.Id)
	if err != nil {
		log.Fatal("failed to delete driver's location from database :", err)
		return nil, err
	}
	if rowsEffect, _ := n.RowsAffected(); rowsEffect == 0 {
		return &dpb.DriverResponse{Status: "Driver does not exists"}, nil
	}
	query = `
		delete from drivers
		where id = $1
	`
	_, err = tx.Exec(query, req.Id)
	if err != nil {
		log.Fatal("failed to delete driver from database :", err)
		return nil, err
	}
	tx.Commit()
	return &dpb.DriverResponse{Status: "Driver is deleted succesfully"}, nil
}

func (dr *DriverRepo) GetActiveDriverFromDatabase(req *dpb.GetLocationRequest) (*dpb.GetLocationResponse, error) {
	var resp dpb.GetLocationResponse
	query := `
		select id, name, email, phone, vehicle from drivers
		where working_region =$1 and status =$2 
	`
	if err := dr.DB.QueryRow(query, req.Location, "active").Scan(&resp.Id, &resp.Name, &resp.Email, &resp.Phone, &resp.Vehicle); err != nil {
		log.Println("failed to get available driver from database")
		return nil, err
	}
	return &resp, nil
}
