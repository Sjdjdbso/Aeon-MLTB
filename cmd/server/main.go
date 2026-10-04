package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
}

type InfoResponse struct {
	Service  string `json:"service"`
	Version  string `json:"version"`
	Platform string `json:"platform"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := HealthResponse{
		Status:    "ok",
		Service:   "aeon-mltb-go",
		Timestamp: time.Now().UTC(),
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := InfoResponse{
		Service:  "aeon-mltb-go",
		Version:  "1.0.0",
		Platform: "Koyeb/Linux",
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/info", infoHandler)
	return mux
}

func main() {
	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8080"
	}

	mux := setupRoutes()
	addr := ":" + port
	log.Printf("Starting Aeon-MLTB Go server on %s", addr)

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
