package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type StatusResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
	Service   string    `json:"service"`
}

var startTime time.Time

func getPort() string {
	if port := os.Getenv("GO_PORT"); port != "" {
		return port
	}
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	resp := StatusResponse{
		Status:    "running",
		Timestamp: time.Now().UTC(),
		Uptime:    time.Since(startTime).String(),
		Service:   "aeon-go-server",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	startTime = time.Now()
	port := getPort()

	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/status", statusHandler)

	addr := fmt.Sprintf(":%s", port)
	log.Printf("[AEON GO] Server starting on port %s...", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("[AEON GO] Server error: %v", err)
	}
}
