package main

import (
	"log"

	"s-store/internal/app"
)

func main() {
	server, err := app.Setup()
	if err != nil {
		log.Fatalf("Error setting up application: %v", err)
	}
	addr := ":8002"
	log.Printf("Server is running on %s", addr)
	if err := server.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
