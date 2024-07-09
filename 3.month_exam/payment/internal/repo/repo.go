package repo

import (
	"database/sql"
	"log"
	pb "payment/protos"
)

type PaymentRepo struct {
	DB *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{DB: db}
}

func (p *PaymentRepo) AddClientCardToDatabase(req *pb.CardRequest) error {
	query := `
		insert into client_account(client_id, card_number, balance)
		values($1, $2, $3)
	`
	_, err := p.DB.Exec(query, req.Id, req.CardNumber, req.Balance)
	if err != nil {
		log.Println("failed to insert client bank information:", err)
		return err
	}
	return nil
}

func (p *PaymentRepo) DeleteClientFromDatabase(req *pb.RequestId)(error){
	query :=`
		delete from client_account where client_id = $1
	`
	_, err :=p.DB.Exec(query, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		log.Fatal("failed to delete  client_account from database")
	}
	return nil 
}

func (p *PaymentRepo) DeleteDriverFromDatabase(req *pb.RequestId)(error){
	query :=`
		delete from driver_account where driver_id = $1
	`
	_, err :=p.DB.Exec(query, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		log.Fatal("failed to delete  driver_account from database")
	}
	return nil 
}

func (p *PaymentRepo) AddDriverCardToDatabase(req *pb.CardRequest) error {
	query := `
		insert into driver_account(driver_id, card_number, balance)
		values($1, $2, $3)
	`
	_, err := p.DB.Exec(query, req.Id, req.CardNumber, req.Balance)
	if err != nil {
		log.Println("failed to insert driver bank inforamtion:", err)
		return err
	}
	return nil
}
func (p *PaymentRepo) MakeProductPurchase(req *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	query := `
			update main_account 
			set balance = balance - $1
		`
	_, err := p.DB.Exec(query, req.Amount)
	if err != nil {
		log.Println("unable to get money from main account")
		return nil, err
	}
	return &pb.PurchaseResponse{Status: "success"}, nil
}
func (p *PaymentRepo) MakeOrderTransaction(req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	tx, _ := p.DB.Begin()
	defer tx.Rollback()
	discount_amount := (req.Discount * int32(req.TotalAmount)) / 100
	total_with_discount := req.TotalAmount - float32(discount_amount)
	driver_amount := (10 * req.TotalAmount) / 100
	remaining_amount := total_with_discount - driver_amount

	query := `update client_account
		set balance = balance - $1
		where client_id = $2
	`
	_, err :=tx.Exec(query, total_with_discount, req.ClientId)
	if err != nil {
		log.Println("Failed withdrawing money from client account")
		tx.Rollback()
	}
	query =`
		update driver_account
		set balance = balance + $1
		where driver_id = $2
	`

	_, err = tx.Exec(query, driver_amount, req.DriverId)
	if err != nil {
		log.Println("Failed adding money to driver account")
		tx.Rollback()
	}

	query = `
		update main_account
		set balance = balance + $1	
	`
	_,err = tx.Exec(query, remaining_amount)
	if err != nil {
		log.Println("Failed adding money to main account")
		return nil, err
	}

	tx.Commit()
	return &pb.PaymentResponse{
		Status: "success",
		DriverAmount: driver_amount,
		Total: req.TotalAmount,
		Discount: float32(req.Discount),
		DiscountAmount: float32(discount_amount),
		TotalWithDiscount: total_with_discount,
		}, nil
}
