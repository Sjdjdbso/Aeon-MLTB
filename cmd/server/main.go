package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var startTime = time.Now()

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Uptime  string `json:"uptime"`
}

type StatusResponse struct {
	Status    string    `json:"status"`
	UptimeSec float64   `json:"uptime_seconds"`
	Timestamp time.Time `json:"timestamp"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := HealthResponse{
		Status:  "ok",
		Service: "aeon-mltb-go",
		Uptime:  time.Since(startTime).String(),
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := StatusResponse{
		Status:    "running",
		UptimeSec: time.Since(startTime).Seconds(),
		Timestamp: time.Now(),
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/status", statusHandler)
	return mux
}

func main() {
	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8085"
	}

	mux := setupRoutes()
	log.Printf("Starting Go server on port %s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
