package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func getPort() string {
	if port := os.Getenv("GO_PORT"); port != "" {
		return port
	}
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}

func getPythonBackendURL() string {
	if backend := os.Getenv("PYTHON_BACKEND_URL"); backend != "" {
		return backend
	}
	return "http://127.0.0.1:8000"
}

func findPageTemplate() string {
	candidates := []string{
		filepath.Join("web", "templates", "page.html"),
		filepath.Join("..", "web", "templates", "page.html"),
		filepath.Join("..", "..", "web", "templates", "page.html"),
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	return ""
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	backendURLStr := getPythonBackendURL()
	backendURL, err := url.Parse(backendURLStr)
	if err != nil {
		log.Fatalf("Invalid PYTHON_BACKEND_URL: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(backendURL)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok","service":"aeon-mltb-go"}`))
	})

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("/app/files/torrent", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	mux.HandleFunc("/app/files", func(w http.ResponseWriter, r *http.Request) {
		if pagePath := findPageTemplate(); pagePath != "" {
			http.ServeFile(w, r, pagePath)
			return
		}
		proxy.ServeHTTP(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			proxy.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "<h1>See mirror-leech-telegram-bot <a href='https://www.github.com/anasty17/mirror-leech-telegram-bot'>@GitHub</a> By <a href='https://github.com/anasty17'>Anas</a></h1>")
	})

	return mux
}

func main() {
	port := getPort()
	mux := setupRoutes()

	addr := ":" + strings.TrimPrefix(port, ":")
	log.Printf("Starting aeon-mltb Go server on %s (Python Backend: %s)", addr, getPythonBackendURL())
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Go server failed to start: %v", err)
	}
}
