package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
	// Set memory limit to 100MB for Koyeb Free Tier (512MB total RAM)
	debug.SetMemoryLimit(100 * 1024 * 1024)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"service": "aeon-mltb-go",
		"status":  "running",
		"uptime":  time.Since(startTime).String(),
	})
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthzHandler)
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/status", statusHandler)
	return mux
}

func main() {
	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8085"
	}

	mux := setupRoutes()
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting Aeon-MLTB Go service on %s", addr)

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Go service failed: %v", err)
	}
}
