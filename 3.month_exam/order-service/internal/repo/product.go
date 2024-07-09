package repo

import (
	"database/sql"
	"fmt"
	"log"
	pb "order-service/protos/product"
	"time"

	"github.com/google/uuid"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (pr *ProductRepo) AddNewProductToDatabase(req *pb.ProductInfo) (*pb.ProductResponse, error) {
	req.Id = uuid.New().String()
	req.CreatedAt = time.Now().Format(time.ANSIC)
	switch req.Category {
	case "Electronics":
		req.ExpiredAt = time.Now().AddDate(1, 0, 0).Format(time.ANSIC)
	case "Apparel":
		req.ExpiredAt = time.Now().AddDate(0, 0, 180).Format(time.ANSIC)
	case "Home & Kitchen":
		req.ExpiredAt = time.Now().AddDate(1, 0, 0).Format(time.ANSIC)
	case "Sports & Outdoors":
		req.ExpiredAt = time.Now().AddDate(0, 0, 180).Format(time.ANSIC)
	case "Books & Media":
		req.ExpiredAt = time.Now().AddDate(0, 0, 180).Format(time.ANSIC)
	case "Beauty & Personal Care":
		req.ExpiredAt = time.Now().AddDate(1, 0, 0).Format(time.ANSIC)
	case "Toys & Games":
		req.ExpiredAt = time.Now().AddDate(1, 0, 0).Format(time.ANSIC)
	case "Furniture":
		req.ExpiredAt = time.Now().AddDate(1, 0, 0).Format(time.ANSIC)
	case "Grocery":
		req.ExpiredAt = time.Now().AddDate(0, 1, 0).Format(time.ANSIC)
	case "Automotive":
		req.ExpiredAt = time.Now().AddDate(1, 0, 0).Format(time.ANSIC)
	default:
		req.ExpiredAt = time.Now().AddDate(0, 1, 0).Format(time.ANSIC)
	}
	query := `
		insert into products (id, name, category, quantity, price, created_at, expired_at)
		values($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := pr.DB.Exec(query, req.Id, req.Name, req.Category, req.Quantity, req.Price, req.CreatedAt, req.ExpiredAt)
	if err != nil {
		log.Println("failed to add new products !", err)
		return nil, err
	}

	return &pb.ProductResponse{Status: "success"}, nil
}

func (pr *ProductRepo) GetProductsByCategory(req *pb.CategoryRequest) ([]*pb.ProductInfo, error) {
	var listProducts []*pb.ProductInfo
	query := `
		select name, category, quantity, price, created_at, expired_at from products where category = $1
	`
	rows, err := pr.DB.Query(query, req.CategoryName)
	if err != nil {
		log.Println("failed to get products by category:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p pb.ProductInfo
		if err := rows.Scan(&p.Name, &p.Category, &p.Quantity, &p.Price, &p.CreatedAt, &p.ExpiredAt); err != nil {
			log.Println("failed to scan products by category:", err)
			return nil, err
		}
		listProducts = append(listProducts, &p)
	}
	if len(listProducts) == 0 {
			return nil, fmt.Errorf("category does not exist ")
	}
	return listProducts, nil
}
