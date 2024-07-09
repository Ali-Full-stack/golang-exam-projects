package grpcclient

import (
	"log"
	"net"
	"os"
	"payment/internal/repo"
	"payment/internal/storage"
	"payment/protos"
	"payment/services"

	"google.golang.org/grpc"
	_"github.com/joho/godotenv/autoload"
	_"github.com/lib/pq"
)

func ConnGrpc() {
	lis, err := net.Listen(os.Getenv("network"), os.Getenv("payment_url"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()
	db, err := storage.OpenSql(os.Getenv("driver_name"), os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()

	pRepo :=repo.NewPaymentRepo(db)
	pServer :=services.NewPaymentServer(pRepo)

	grpcServer :=grpc.NewServer()
	protos.RegisterPaymentServiceServer(grpcServer, pServer)

	log.Println("Payment: server is listening on port ", os.Getenv("payment_url"))
	log.Fatal(grpcServer.Serve(lis))
}
