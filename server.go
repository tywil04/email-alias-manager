package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"EmailAliasManager/backend"
	"EmailAliasManager/frontend"
	"EmailAliasManager/sharedlib/db"

	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Starting server...")

	// load .env
	godotenv.Load()

	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = ":4041"
	}

	server := &http.Server{
		Addr: address,
	}

	// connect to db
	err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database, reason: %s\n", err)
	}

	// register backend
	backend.Backend()

	// register ui
	frontend.Frontend()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	log.Printf("Server started on %s", server.Addr)

	<-quit
	log.Print("Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server stopped ungracefully")
	}

	log.Print("Server stopped gracefully")
}
