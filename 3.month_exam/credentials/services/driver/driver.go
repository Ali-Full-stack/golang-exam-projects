package driver

import (
	"context"
	"credentials/internal/repos"
	dpb "credentials/protos/driverpb"
	"log"
)

type DriverServer struct {
	dpb.UnimplementedDriverServiceServer
	DriverRepo *repos.DriverRepo
}

func NewDriverServer(drRepo *repos.DriverRepo)*DriverServer{
	return &DriverServer{DriverRepo: drRepo}
}

func (d *DriverServer) CreateDriver(ctx context.Context, req *dpb.DriverInfo)(*dpb.DriverID, error){
	resp, err :=d.DriverRepo.AddNewDriverToDatabase(req)
	if err != nil {
		log.Println("Failed to add new driver to database",err)
		return nil, err
	}

	return resp, nil
}

func (d *DriverServer) DeleteDriver(ctx context.Context, req *dpb.DriverID)(*dpb.DriverResponse, error){
	resp, err :=d.DriverRepo.DeleteDriverFromDatabase(req)
	if err != nil {
		log.Println("Failed to delete   driver from database",err)
		return nil, err
	}
	return resp, nil
}

func (d *DriverServer) GetAvailableDriver(ctx context.Context, req *dpb.GetLocationRequest)(*dpb.GetLocationResponse, error){
	resp, err :=d.DriverRepo.GetActiveDriverFromDatabase(req)
	if err != nil {
		log.Println("Unable to get active drivers from database",err)
		return nil, err
	}
	return resp, nil
}