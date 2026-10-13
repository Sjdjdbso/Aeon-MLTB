package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var startTime time.Time

func getListenPort() string {
	if port := os.Getenv("GO_PORT"); port != "" {
		return port
	}
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	status := map[string]interface{}{
		"status":      "running",
		"uptime_sec":  int(time.Since(startTime).Seconds()),
		"alloc_bytes": m.Alloc,
		"sys_bytes":   m.Sys,
		"goroutines":  runtime.NumGoroutine(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(status)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`<h1>Aeon-MLTB Go Microservice</h1><p>Status: Active</p>`))
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"files":   []string{},
		"engine":  "golang",
		"message": "Aeon-MLTB Go File Service Ready",
	})
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api/status", statusHandler)
	mux.HandleFunc("/app/files", filesHandler)
	return mux
}

func main() {
	startTime = time.Now()
	port := getListenPort()
	addr := ":" + port

	mux := setupRoutes()

	log.Printf("Starting Aeon-MLTB Go Server on port %s...", port)
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
