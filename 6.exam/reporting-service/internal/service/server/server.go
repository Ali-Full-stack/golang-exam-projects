package server

import (
	"log"
	"net"
	"os"
	"report-service/internal/service/client"
	"report-service/internal/service/handler"
	"report-service/proto/report"
	"report-service/storage"
	"report-service/storage/postgres"

	"google.golang.org/grpc"
)

func Routes() {
	lis, err := net.Listen("tcp", os.Getenv("report_server"))
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	defer lis.Close()

	budgetClient, err := client.DialBudgetClient((os.Getenv("budget_url")))
	if err != nil {
		log.Println("Failed to dial to budget-service:", err)
	}

	transactionClient, err := client.DialTransactionClient((os.Getenv("transaction_url")))
	if err != nil {
		log.Println("Failed to dial to transaction-service:", err)
	}

	db, err := postgres.ConnectPostgres(os.Getenv("postgres_url"))
	if err != nil {
		log.Println(err)
	}
	queries := storage.New(db)

	reportService := handler.NewReportService(budgetClient, transactionClient, queries)

	s := grpc.NewServer()
	report.RegisterReportServiceServer(s, reportService)

	log.Println("Report-Service: server is listening on port:", os.Getenv("report_server"))
	log.Fatal(s.Serve(lis))

}
