package main

import (
	"log"
	"os"

	"aeon-mltb/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := server.NewServer(port)
	log.Printf("Starting Go web server on port %s...", port)
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
