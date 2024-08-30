package routes

import (
	"gateway/api/handler"
	"gateway/api/middleware"
	_ "gateway/api/swagger/docs"
	"gateway/client"
	ratelimiting "gateway/internal/rateLimiting"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	swag "github.com/swaggo/http-swagger"
)

// New ...
// @title  Project: Hotel Booking API
// @description This swagger UI was created to manage hotel bookings
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:9000
// @contact.email  ali.team@gmail.com
func Routes() {
	mux := http.NewServeMux()

	rateLimit := ratelimiting.NewRateLimiter(4, time.Minute)

	userClient := client.DialClient(os.Getenv("user_url"))
	userHandler := handler.NewUserHandler(userClient)

	mux.HandleFunc("POST /api/users/register", rateLimit.Limit(userHandler.RegisterNewUser))
	mux.HandleFunc("POST /api/users/login", rateLimit.Limit(userHandler.UserLogin))
	mux.Handle("PUT /api/users/{id}", middleware.ConfirmToken(rateLimit.Limit(userHandler.UpdateUser)))
	mux.Handle("DELETE /api/users/{id}", middleware.ConfirmToken(rateLimit.Limit(userHandler.DeleteUser)))
	mux.Handle("GET /api/users", middleware.ConfirmToken(rateLimit.Limit(userHandler.GetAllUsers)))
	mux.Handle("GET /api/users/{id}", middleware.ConfirmToken(rateLimit.Limit(userHandler.GetUserById)))

	hotelClient := client.DialHotelClient(os.Getenv("hotel_url"))
	hotelHandler := handler.NewHotelHandler(hotelClient)
	mux.Handle("POST /api/hotels", middleware.ConfirmToken(rateLimit.Limit(hotelHandler.CreateHotel)))
	mux.Handle("POST /api/hotels/{id}/rooms", middleware.ConfirmToken(rateLimit.Limit(hotelHandler.CreateHotelRoom)))
	mux.Handle("PUT /api/hotels/{id}", middleware.ConfirmToken(rateLimit.Limit(hotelHandler.UpdateHotel)))
	mux.Handle("PUT /api/hotels/{id}/rooms", middleware.ConfirmToken(rateLimit.Limit(hotelHandler.UpdateHotelRoom)))
	mux.Handle("DELETE /api/hotels/{id}", middleware.ConfirmToken(rateLimit.Limit(hotelHandler.DeleteHotel)))
	mux.Handle("DELETE /api/hotels/{id}/rooms", middleware.ConfirmToken(rateLimit.Limit(hotelHandler.DeleteHotelRoom)))
	mux.Handle("GET /api/hotels", middleware.ConfirmToken(rateLimit.Limit(hotelHandler.GetAllHotels)))
	mux.Handle("GET /api/hotels/{id}", middleware.ConfirmToken(rateLimit.Limit(hotelHandler.GetHotelById)))
	mux.Handle("POST /api/hotels/{id}/rooms/availability", middleware.ConfirmToken(rateLimit.Limit(hotelHandler.CheckAvailableRooms)))
	mux.Handle("PATCH /api/hotels/{id}/rooms", middleware.ConfirmToken(rateLimit.Limit(hotelHandler.UpdateRoomCount)))

	bookingClient := client.DialBookingClient(os.Getenv("booking_url"))
	bookingHandler := handler.NewBookingHandler(bookingClient)
	mux.Handle("POST /api/bookings", middleware.ConfirmToken(rateLimit.Limit(bookingHandler.CreateBooking)))
	mux.Handle("DELETE /api/bookings/{id}", middleware.ConfirmToken(rateLimit.Limit(bookingHandler.DeleteBooking)))
	mux.Handle("GET /api/bookings/{id}", middleware.ConfirmToken(rateLimit.Limit(bookingHandler.GetBookingById)))

	mux.Handle("/swagger/", swag.WrapHandler)

	// tlsConfig := &tls.Config{
	// 	CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	// }

	// srv := &http.Server{
	// 	Addr:      os.Getenv("api_url"),
	// 	Handler:   mux,
	// 	TLSConfig: tlsConfig,
	// }
	// go shutdown.GracefulShutdown(srv)

	// fmt.Printf("Server started on port %s\n", os.Getenv("api_url"))
	// err := srv.ListenAndServeTLS("./internal/tls/items.pem", "./internal/tls/items-key.pem")
	// log.Println(err.Error())
	// os.Exit(1)

	log.Println("Gateway : server is listening on port :", os.Getenv("api_url"))
	log.Fatal(http.ListenAndServe(os.Getenv("api_url"), mux))
}
