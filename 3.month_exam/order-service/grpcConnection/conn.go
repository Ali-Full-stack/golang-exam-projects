package grpcconnection

import (
	"log"
	"net"
	"order-service/internal/repo"
	"order-service/internal/storage"
	"order-service/protos/order"
	"order-service/protos/product"
	"order-service/service"
	"os"

	"google.golang.org/grpc"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func ConnGrpc(){
	lis, err := net.Listen(os.Getenv("network"), os.Getenv("order_url"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()
	db, err := storage.OpenSql(os.Getenv("driver_name"), os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()

	pRepo :=repo.NewProductRepo(db)
	oRepo :=repo.NewOrderRepo(db)
	
	productServer :=service.NewProductServer(pRepo)
	orderServer :=service.NewOrderService(oRepo)
	
	grpcServer :=grpc.NewServer()

	order.RegisterOrderServiceServer(grpcServer, orderServer)
	product.RegisterProductServiceServer(grpcServer, productServer)

	log.Println("orders: server is listening on port ", os.Getenv("order_url"))
	log.Fatal(grpcServer.Serve(lis))
}
