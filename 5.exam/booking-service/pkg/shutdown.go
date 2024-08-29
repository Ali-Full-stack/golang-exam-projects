package pkg

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func Shutdown(grpcServer *grpc.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop
	log.Println("received shutdown signal..")
	log.Println("signal :", sig.String())
	grpcServer.GracefulStop()
	log.Println("shutting down server")
}
