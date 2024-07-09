package connection

import (
	"credentials/internal/repos"
	"credentials/internal/storage"
	"credentials/protos/adminpb"
	"credentials/protos/clientpb"
	"credentials/protos/driverpb"
	"credentials/services/admin"
	"credentials/services/client"
	"credentials/services/driver"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func ConnGrpc() {
	lis, err := net.Listen(os.Getenv("network"), os.Getenv("credential_url"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()
	db, err := storage.OpenSql(os.Getenv("driver_name"), os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()

	clrepo := repos.NewClientRepo(db)
	adrepo := repos.NewAdminRepo(db)
	drrepo := repos.NewDriverRepo(db)

	Clientserver := client.NewClientServer(clrepo)
	Adminserver := admin.NewAdminServer(adrepo)
	Driverserver := driver.NewDriverServer(drrepo)

	grpcServer := grpc.NewServer()
	
	clientpb.RegisterClientServiceServer(grpcServer, Clientserver)
	adminpb.RegisterAdminServiceServer(grpcServer, Adminserver)
	driverpb.RegisterDriverServiceServer(grpcServer, Driverserver)


	log.Println("Credentials: server is listening on port ", os.Getenv("credential_url"))
	log.Fatal(grpcServer.Serve(lis))
}
