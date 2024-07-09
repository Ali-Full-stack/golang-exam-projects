package repos

import (
	"credentials/auth/hash"
	apb "credentials/protos/adminpb"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type AdminRepo struct {
	DB *sql.DB
}

func NewAdminRepo(db *sql.DB) *AdminRepo {
	return &AdminRepo{DB: db}
}

func(ar *AdminRepo) AddNewAdminToDatabase(req *apb.AdminInfo) (*apb.AdminResponse, error) {
	req.Id = uuid.New().String()
	hash :=hash.GenerateHashPassword(req.Password)
	role :="normal"
	query :=`
		insert into admins (id, name, email, password, role)
		values($1, $2, $3, $4, $5)
	`
	_, err :=ar.DB.Exec(query, req.Id, req.Name, req.Email, hash, role)
	if err != nil {
		log.Println("failed to add new admin to database",err)
		return nil, err
	}
	return &apb.AdminResponse{Id: req.Id, HashPassword: hash, Status: "active"}, nil
}	

