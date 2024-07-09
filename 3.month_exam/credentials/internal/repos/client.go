package repos

import (
	cpb"credentials/protos/clientpb"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type ClientRepo struct {
	DB *sql.DB
}

func NewClientRepo(db *sql.DB) *ClientRepo {
	return &ClientRepo{DB: db}
}

func (cr *ClientRepo) AddNewClientToDatabase(req *cpb.ClientInfo) (*cpb.ClientID,error) {
	tx, _ :=cr.DB.Begin()
	defer tx.Rollback()

	req.Id = uuid.New().String()
	req.CreatedAt =time.Now().Format(time.ANSIC)

	query :=`
		insert into clients(id, name, email, phone, created_at)
		values($1, $2, $3, $4, $5)
	`
	_, err := tx.Exec(query, req.Id, req.Name, req.Email, req.Phone, req.CreatedAt)
	if err != nil {
		log.Println("failed to add new client to database:",err)
		return nil, err
	}

	query =`
		insert into client_location (client_id, city, region, home_address)
		values($1, $2, $3, $4 )
	`
	_, err =tx.Exec(query, req.Id, req.Address.City, req.Address.Region, req.Address.HomeAddress)
	if err != nil {
		log.Println("failed to add new client address to database:",err)
		return nil, err
	}

	tx.Commit()

	return &cpb.ClientID{Id: req.Id}, nil

}

func (cr *ClientRepo) DeleteClientFromDatabase(req *cpb.ClientID)(*cpb.ClientResponse, error){
	tx, _ :=cr.DB.Begin()
	defer tx.Rollback()
	query :=`
		delete from client_location
		where client_id = $1
	`
	n, err :=tx.Exec(query, req.Id)
	if err != nil {
		log.Println("Unable to  delete client_location from database:",err)
		return nil, err
	}
	if rowsEffected,_ :=n.RowsAffected(); rowsEffected == 0 {
		return &cpb.ClientResponse{Status: " Client does not  exist !!"}, nil
	}
	query = `
		delete from clients
		where id = $1
	`
	_, err =tx.Exec(query, req.Id)
	if err != nil {
		log.Println("Unable to  delete client from database:",err)
		return nil, err
	}
	tx.Commit()
	return &cpb.ClientResponse{Status: "Client deletion succesfull ."},nil
}	

func (c *ClientRepo) 	GetClientLocationFromDatabase(req *cpb.ClientID)(*cpb.ClientLocation,error){
	query :=`
		select c.phone, c.email, l.city, l.region, l.home_address from clients c
		join client_location l on c.id = l.client_id 
		where c.id = $1
	`
	var phone , email, city, region, homeAdress string
	err :=c.DB.QueryRow(query, req.Id).Scan(&phone, &email, &city, &region, &homeAdress)
	if err != nil {
		log.Println("Failed to get client location")
		return nil, err
	}

	return &cpb.ClientLocation{Phone: phone, Email: email, Address:&cpb.Address{City: city, Region: region, HomeAddress: homeAdress} }, nil
}