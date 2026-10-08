package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"status":"running","service":"aeon-mltb-go","migration":"gradual"}`)
}

func main() {
	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8085"
	}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/status", statusHandler)

	log.Printf("Starting lightweight Go server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
