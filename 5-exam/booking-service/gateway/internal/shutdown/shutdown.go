package shutdown

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutdown(srv *http.Server) {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	<-shutdownCh
	log.Println("Shutdown signal received, initiating graceful shutdown...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Println("Server shutdown encountered an error: " + err.Error())
	} else {
		log.Println("Server gracefully stopped")
	}

	select {
	case <-shutdownCtx.Done():
		if shutdownCtx.Err() == context.DeadlineExceeded {
			log.Println("Shutdown deadline exceeded, forcing server to stop")
		}
	default:
		log.Println("Shutdown completed within the timeout period")
	}
}
