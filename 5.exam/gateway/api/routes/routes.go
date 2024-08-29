package routes

import (
	"gateway/api/handler"
	_ "gateway/api/swagger/docs"
	"gateway/client"
	ratelimiting "gateway/internal/rateLimiting"
	"log"
	"net/http"
	"os"
	"time"
	_"gateway/api/swagger/docs"
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
	mux.HandleFunc("PUT /api/users/{id}", rateLimit.Limit(userHandler.UpdateUser))
	mux.HandleFunc("DELETE /api/users/{id}", rateLimit.Limit(userHandler.DeleteUser))
	mux.HandleFunc("GET /api/users", rateLimit.Limit(userHandler.GetAllUsers))
	mux.HandleFunc("GET /api/users/{id}", rateLimit.Limit(userHandler.GetUserById))

	hotelClient := client.DialHotelClient(os.Getenv("hotel_url"))
	hotelHandler := handler.NewHotelHandler(hotelClient)
	mux.HandleFunc("POST /api/hotels", rateLimit.Limit(hotelHandler.CreateHotel))
	mux.HandleFunc("POST /api/hotels/{id}/rooms", rateLimit.Limit(hotelHandler.CreateHotelRoom))
	mux.HandleFunc("PUT /api/hotels/{id}", rateLimit.Limit(hotelHandler.UpdateHotel))
	mux.HandleFunc("PUT /api/hotels/{id}/rooms", rateLimit.Limit(hotelHandler.UpdateHotelRoom))
	mux.HandleFunc("DELETE /api/hotels/{id}", rateLimit.Limit(hotelHandler.DeleteHotel))
	mux.HandleFunc("DELETE /api/hotels/{id}/rooms", rateLimit.Limit(hotelHandler.DeleteHotelRoom))
	mux.HandleFunc("GET /api/hotels", rateLimit.Limit(hotelHandler.GetAllHotels))
	mux.HandleFunc("GET /api/hotels/{id}", rateLimit.Limit(hotelHandler.GetHotelById))
	mux.HandleFunc("POST /api/hotels/{id}/rooms/availability", rateLimit.Limit(hotelHandler.CheckAvailableRooms))
	mux.HandleFunc("PATCH /api/hotels/{id}/rooms", rateLimit.Limit(hotelHandler.UpdateRoomCount))

	bookingClient := client.DialBookingClient(os.Getenv("booking_url"))
	bookingHandler := handler.NewBookingHandler(bookingClient)
	mux.HandleFunc("POST /api/bookings", rateLimit.Limit(bookingHandler.CreateBooking))
	mux.HandleFunc("DELETE /api/bookings/{id}", rateLimit.Limit(bookingHandler.DeleteBooking))
	mux.HandleFunc("GET /api/bookings/{id}", rateLimit.Limit(bookingHandler.GetBookingById))

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
