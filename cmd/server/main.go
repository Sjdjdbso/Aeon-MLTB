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
	Uptime    string    `json:"uptime"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
}

var startTime time.Time

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := StatusResponse{
		Status:    "ok",
		Uptime:    time.Since(startTime).String(),
		Timestamp: time.Now(),
		Service:   "aeon-mltb-go-service",
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", healthHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/ping", pingHandler)
	return mux
}

func main() {
	startTime = time.Now()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := setupRoutes()
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting Go health server on %s", addr)

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Go server error: %v", err)
	}
}
