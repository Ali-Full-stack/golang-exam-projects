package admin

import (
	"context"
	"credentials/internal/repos"
	apb "credentials/protos/adminpb"
	"fmt"
)

type AdminServer struct {
	apb.UnimplementedAdminServiceServer
	AdminRepo  *repos.AdminRepo
}

func NewAdminServer(adRepo *repos.AdminRepo)*AdminServer{
	return &AdminServer{AdminRepo: adRepo}
}

func (a *AdminServer) AddNewAdmin(ctx context.Context, req *apb.AdminInfo)(*apb.AdminResponse,error){

		resp,err :=a.AdminRepo.AddNewAdminToDatabase(req)
		if err != nil {
			return nil, fmt.Errorf("unable to add new admin into database: %v",err)
		}
		return resp, nil
}