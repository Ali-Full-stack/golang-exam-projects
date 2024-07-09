package client

import (
	"context"
	"credentials/internal/repos"
	cpb"credentials/protos/clientpb"
	"fmt"
)

type ClientServer struct {
	cpb.UnimplementedClientServiceServer
	ClientRepo  *repos.ClientRepo	
}

func NewClientServer(clRepo *repos.ClientRepo)*ClientServer{
	return &ClientServer{ClientRepo: clRepo}
}

func (c *ClientServer) CreateClient(ctx context.Context, req *cpb.ClientInfo)(*cpb.ClientID, error){

	resp, err := c.ClientRepo.AddNewClientToDatabase(req)
	if err != nil {
		return nil, fmt.Errorf("failed to add new client:%v",err)
	}
	return resp, nil
}

func (c *ClientServer) DeleteClient(ctx context.Context, req *cpb.ClientID)(*cpb.ClientResponse, error){
	resp, err :=c.ClientRepo.DeleteClientFromDatabase(req)
	if err != nil {
		return nil, fmt.Errorf("unable to delete client from database %w ",err)
	}

	return resp, nil
}

func(c *ClientServer) GetClientLocation(ctx context.Context, req *cpb.ClientID)(*cpb.ClientLocation, error){

	resp, err := c.ClientRepo.GetClientLocationFromDatabase(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
