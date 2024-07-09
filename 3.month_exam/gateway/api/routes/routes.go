package routes

import (
	"gateway/api/handler"
	"gateway/api/middleware"
	credentialclient "gateway/grpc-client/credential-client"
	ordersclient "gateway/grpc-client/orders-client"
	paymentclient "gateway/grpc-client/payment-client"
	"gateway/internal/logs"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	lg := logs.GetLogger("internal/logs/history.log")

	crClient := credentialclient.DialClient(os.Getenv("credential_url"))
	adClient := credentialclient.DialAdmin(os.Getenv("credential_url"))
	drClient := credentialclient.DialDriver(os.Getenv("credential_url"))
	pClient := paymentclient.DialPayment(os.Getenv("payment_url"))
	proClient :=ordersclient.Dialproduct(os.Getenv("orders_url"))
	orClient :=ordersclient.DialOrder(os.Getenv("orders_url"))


	clientManager := handler.NewClientManager(crClient,  pClient, lg)
	mux.Handle("POST /client/register",middleware.ValidateRequestBody(http.HandlerFunc( clientManager.RegisterNewClient)))
	mux.Handle("DELETE /client/delete", middleware.CheckAdminPassword(http.HandlerFunc(clientManager.DeleteClient)))
	mux.Handle("GET /client/login", middleware.LoginAccess(http.HandlerFunc(clientManager.ClientLogin)))


	driverManager :=handler.NewDriverManager(drClient, pClient, lg)
	mux.Handle("POST /driver/create", middleware.CheckAdminPassword(middleware.ValidateRequestDriver(http.HandlerFunc(driverManager.AddNewDriver))))
	mux.Handle("DELETE /driver/delete", middleware.CheckAdminPassword(http.HandlerFunc(driverManager.DeleteDriver)))


	productManager :=handler.NewProductManager(proClient,pClient, lg)
	mux.Handle("POST /product/create", middleware.CheckAdminPassword(http.HandlerFunc(productManager.AddNewProducts)))
	mux.HandleFunc("GET /product/category", productManager.GetProductsByCategory)

	orderManager :=handler.NewOrderManager(crClient, drClient,pClient,orClient,lg)
	mux.Handle("POST /order/create",middleware.CheckOrderToken(http.HandlerFunc(orderManager.CreateOrders)))

	
	adHandler := handler.NewAdminHandler(adClient, lg)
	mux.Handle("POST /admin/create", middleware.IsSuperAdmin(http.HandlerFunc(adHandler.AddNewAdmin)))

	return mux

}
