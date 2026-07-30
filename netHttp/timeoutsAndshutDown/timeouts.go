package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Mux *http.ServeMux = http.NewServeMux()

func main() {
	Mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		for i := 0; i < 100000000000000; i++ {
		}
	})

	var server http.Server = http.Server{
		Addr:         ":5000",
		Handler:      Mux,
		ReadTimeout:  15 * time.Second, // Max time to read the whole request
		WriteTimeout: 15 * time.Second, // Max time to write the response
		IdleTimeout:  60 * time.Second, // Max time to wait for next request on keep-alive
	}
	go func() {
		log.Println("Server starting on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
