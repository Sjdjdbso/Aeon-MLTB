package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"aeon-mltb/pkg/utils"
)

type RewriteRequest struct {
	URL string `json:"url"`
}

type RewriteResponse struct {
	OriginalURL  string `json:"original_url"`
	RewrittenURL string `json:"rewritten_url"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.GetHealthStatus())
}

func rewriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RewriteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	rewritten := utils.RewritePixeldrainURL(req.URL)
	resp := RewriteResponse{
		OriginalURL:  req.URL,
		RewrittenURL: rewritten,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Set low GOMAXPROCS for Koyeb Free Tier memory constraints if not already set
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(1)
	}

	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/rewrite-url", rewriteHandler)

	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println(utils.FormatLogMessage(fmt.Sprintf("Starting Go server on port %s...", port)))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Go server error: %v", err)
	}
}
